package useractions

import (
	"context"
	"errors"
	"testing"

	"github.com/sensu/sensu-go/testing/mockstore"
	"github.com/sensu/sensu-go/testing/testutil"
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewEventActions(t *testing.T) {
	assert := assert.New(t)

	store := &mockstore.MockStore{}
	eventActions := NewEventActions(store)

	assert.NotNil(eventActions)
	assert.Equal(store, eventActions.Store)
	assert.NotNil(eventActions.Policy)
	assert.Nil(eventActions.Context)

	ctx := context.TODO()
	eventActions = eventActions.WithContext(ctx)

	assert.Equal(ctx, eventActions.Context)
}

func TestQuery(t *testing.T) {
	defaultCtx := testutil.NewContext(testutil.ContextWithRules(
		types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermRead),
	))

	testCases := []struct {
		name        string
		ctx         context.Context
		events      []*types.Event
		params      QueryParams
		expectedLen int
		storeErr    error
		expectedErr error
	}{
		{
			name:        "No Params No Events",
			ctx:         defaultCtx,
			events:      []*types.Event{},
			params:      QueryParams{},
			expectedLen: 0,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name: "No Params With Events",
			ctx:  defaultCtx,
			events: []*types.Event{
				types.FixtureEvent("entity1", "check1"),
				types.FixtureEvent("entity2", "check2"),
			},
			params:      QueryParams{},
			expectedLen: 2,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name: "No Params With Only Create Access",
			ctx: testutil.NewContext(testutil.ContextWithRules(
				types.FixtureRuleWithPerms(types.RuleTypeEvent, types.RulePermCreate),
			)),
			events: []*types.Event{
				types.FixtureEvent("entity1", "check1"),
				types.FixtureEvent("entity2", "check2"),
			},
			params:      QueryParams{},
			expectedLen: 0,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name: "Entity Param",
			ctx:  defaultCtx,
			events: []*types.Event{
				types.FixtureEvent("entity1", "check1"),
			},
			params: QueryParams{
				"entity": "entity1",
			},
			expectedLen: 1,
			storeErr:    nil,
			expectedErr: nil,
		},
		{
			name:   "Store Failure",
			ctx:    defaultCtx,
			events: nil,
			params: QueryParams{
				"entity": "entity1",
			},
			expectedLen: 0,
			storeErr:    errors.New(""),
			expectedErr: NewError(InternalErr, errors.New("")),
		},
	}

	for _, tc := range testCases {
		store := &mockstore.MockStore{}
		eventActions := NewEventActions(store)

		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			// Mock store methods
			store.On("GetEvents", tc.ctx).Return(tc.events, tc.storeErr)
			store.On("GetEventsByEntity", tc.ctx, mock.Anything).Return(tc.events, tc.storeErr)

			// Exec Query
			fetcher := eventActions.WithContext(tc.ctx)
			results, err := fetcher.Query(tc.params)

			// Assert
			assert.EqualValues(tc.expectedErr, err)
			assert.Len(results, tc.expectedLen)
		})
	}
}

func TestFind(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, true)
}
