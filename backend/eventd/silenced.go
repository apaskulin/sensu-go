package eventd

import (
	"context"
	"fmt"
	"time"

	"github.com/sensu/sensu-go/backend/store"
	"github.com/sensu/sensu-go/types"
	stringsutil "github.com/sensu/sensu-go/util/strings"
)

// addToSilencedBy takes a silenced entry ID and adds it to a silence of IDs if
// it's not already present in order to avoid duplicated elements
func addToSilencedBy(id string, ids []string) []string {
	if !stringsutil.InArray(id, ids) {
		ids = append(ids, id)
	}
	return ids
}

// getSilenced retrieves all silenced entries for a given event, using the
// entity subscription, the check subscription and the check name while
// supporting wildcard silenced entries (e.g. subscription:*)
func getSilenced(ctx context.Context, event *types.Event, s store.Store) error {
	entries := []*types.Silenced{}
	if !event.HasCheck() {
		return nil
	}

	// Retrieve silenced entries using the entity subscription
	entitySubscription := types.GetEntitySubscription(event.Entity.Name)
	results, err := s.GetSilencedEntriesBySubscription(ctx, entitySubscription)
	if err != nil {
		return err
	}
	entries = append(entries, results...)

	// Retrieve silenced entries using the check subscriptions
	for _, value := range event.Check.Subscriptions {
		results, err = s.GetSilencedEntriesBySubscription(ctx, value)
		if err != nil {
			return err
		}
		entries = append(entries, results...)
	}

	// Retrieve silenced entries using the check name
	results, err = s.GetSilencedEntriesByCheckName(ctx, event.Check.Name)
	if err != nil {
		return err
	}
	entries = append(entries, results...)

	// Determine which entries silence this event
	silencedIDs := silencedBy(event, entries)

	// Add to the event all silenced entries ID that actually silence it
	event.Check.Silenced = silencedIDs

	return nil
}

// silencedBy determines which of the given silenced entries silenced a given
// event and return a list of silenced entry IDs
func silencedBy(event *types.Event, silencedEntries []*types.Silenced) []string {
	silencedBy := []string{}
	if !event.HasCheck() {
		return silencedBy
	}

	// Loop through every silenced entries in order to determine if it applies to
	// the given event
	for _, entry := range silencedEntries {

		// Is this event silenced for all subscriptions? (e.g. *:check_cpu)
		if entry.Name == fmt.Sprintf("*:%s", event.Check.Name) && entry.StartSilence(time.Now().Unix()) {
			silencedBy = addToSilencedBy(entry.Name, silencedBy)
			continue
		}

		// Is this event silenced by the entity subscription? (e.g. entity:id:*)
		if entry.Name == fmt.Sprintf("%s:*", types.GetEntitySubscription(event.Entity.Name)) && entry.StartSilence(time.Now().Unix()) {
			silencedBy = addToSilencedBy(entry.Name, silencedBy)
			continue
		}

		// Is this event silenced for this particular entity? (e.g.
		// entity:id:check_cpu)
		if entry.Name == fmt.Sprintf("%s:%s", types.GetEntitySubscription(event.Entity.Name), event.Check.Name) && entry.StartSilence(time.Now().Unix()) {
			silencedBy = addToSilencedBy(entry.Name, silencedBy)
			continue
		}

		for _, subscription := range event.Check.Subscriptions {
			// Make sure the entity is subscribed to this specific subscription
			if !stringsutil.InArray(subscription, event.Entity.Subscriptions) {
				continue
			}

			// Is this event silenced by one of the check subscription? (e.g.
			// load-balancer:*)
			if entry.Name == fmt.Sprintf("%s:*", subscription) && entry.StartSilence(time.Now().Unix()) {
				silencedBy = addToSilencedBy(entry.Name, silencedBy)
				continue
			}

			// Is this event silenced by one of the check subscription for this
			// particular check? (e.g. load-balancer:check_cpu)
			if entry.Name == fmt.Sprintf("%s:%s", subscription, event.Check.Name) && entry.StartSilence(time.Now().Unix()) {
				silencedBy = addToSilencedBy(entry.Name, silencedBy)
				continue
			}
		}
	}

	return silencedBy
}

func handleExpireOnResolveEntries(ctx context.Context, event *types.Event, store store.Store) error {
	// Make sure we have a check and that the event is a resolution
	if !event.HasCheck() || !event.IsResolution() {
		return nil
	}

	nonExpireOnResolveEntries := []string{}

	for _, silencedID := range event.Check.Silenced {
		silencedEntry, err := store.GetSilencedEntryByName(ctx, silencedID)
		if err != nil {
			return err
		}

		if silencedEntry.ExpireOnResolve {
			err := store.DeleteSilencedEntryByName(ctx, silencedID)
			if err != nil {
				return err
			}
		} else {
			nonExpireOnResolveEntries = append(nonExpireOnResolveEntries, silencedID)
		}
	}

	event.Check.Silenced = nonExpireOnResolveEntries

	return nil
}
