package rbac

import (
	"context"

	"github.com/sensu/sensu-go/backend/authorization"
	"github.com/sensu/sensu-go/backend/store"
	"github.com/sensu/sensu-go/types"
	"github.com/sirupsen/logrus"
)

// Authorizer implements an authorizer interface using Role-Based Acccess
// Control (RBAC)
type Authorizer struct {
	Store store.Store
}

// Authorize determines if a request is authorized based on its attributes
func (a *Authorizer) Authorize(attrs *authorization.Attributes) (bool, error) {
	ctx := context.Background()

	if attrs != nil {
		logger = logger.WithFields(logrus.Fields{
			"zz_debug": map[string]string{
				"apiGroup":     attrs.APIGroup,
				"apiVersion":   attrs.APIVersion,
				"namespace":    attrs.Namespace,
				"resource":     attrs.Resource,
				"resourceName": attrs.ResourceName,
				"username":     attrs.User.Username,
				"verb":         attrs.Verb,
			},
		})
	}

	// Get cluster roles binding
	clusterRoleBindings, err := a.Store.ListClusterRoleBindings(ctx)
	if err != nil {
		switch err := err.(type) {
		case *store.ErrNotFound:
			// No ClusterRoleBindings founds, let's continue with the RoleBindings
			logger.WithError(err).Debug("no ClusterRoleBindings found")
			break
		default:
			logger.WithError(err).Warning("could not retrieve the ClusterRoleBindings")
			return false, err
		}
	}

	// Inspect each cluster role binding
	for _, clusterRoleBinding := range clusterRoleBindings {
		bindingName := clusterRoleBinding.Name
		roleName := clusterRoleBinding.RoleRef.Name

		// Verify if this cluster role binding matches our user
		if !matchesUser(attrs.User, clusterRoleBinding.Subjects) {
			logger.Debugf("the user is not a subject of the ClusterRoleBinding %s", bindingName)
			continue
		}

		// Get the cluster role that matched our user
		clusterRole, err := a.Store.GetClusterRole(ctx, roleName)
		if err != nil {
			logger.WithError(err).Warningf(
				"could not retrieve the ClusterRole %s", roleName,
			)
			return false, err
		} else if clusterRole == nil {
			logger.Warningf(
				"ClusterRole %s is empty", roleName,
			)
			continue
		}

		// Loop through the cluster role rules
		for _, rule := range clusterRole.Rules {
			// Verify if this rule applies to our request
			allowed, reason := ruleAllows(attrs, rule)
			if allowed {
				logger.Debugf("request authorized by the ClusterRoleBinding %s", bindingName)
				return true, nil
			}
			logger.Tracef("%s by rule %+v", reason, rule)
		}
		logger.Debugf("could not authorize the request with the ClusterRoleBinding %s",
			bindingName,
		)
		logger.Debugf("could not authorize the request with any ClusterRoleBindings")
	}

	// None of the cluster roles authorized our request. Let's try with roles
	// First, make sure we have a namespace
	if len(attrs.Namespace) > 0 {
		// Get roles binding
		roleBindings, err := a.Store.ListRoleBindings(ctx)
		if err != nil {
			switch err := err.(type) {
			case *store.ErrNotFound:
				// No ClusterRoleBindings founds, let's continue with the RoleBindings
				logger.WithError(err).Debug("no RoleBindings found")
				break
			default:
				logger.WithError(err).Warning("could not retrieve the ClusterRoleBindings")
				return false, err
			}
		}

		// Inspect each role binding
		for _, roleBinding := range roleBindings {
			bindingName := roleBinding.Name
			roleName := roleBinding.RoleRef.Name

			// Verify if this role binding matches our user
			if !matchesUser(attrs.User, roleBinding.Subjects) {
				logger.Debugf("the user is not a subject of the RoleBinding %s", bindingName)
				continue
			}

			// Get the role that matched our user
			role, err := a.Store.GetRole(ctx, roleName)
			if err != nil {
				logger.WithError(err).Warningf(
					"could not retrieve the Role %s", roleName,
				)
				return false, err
			} else if role == nil {
				logger.Warningf(
					"Role %s is empty", roleName,
				)
				continue
			}

			// Loop through the role rules
			for _, rule := range role.Rules {
				// Verify if this rule applies to our request
				allowed, reason := ruleAllows(attrs, rule)
				if allowed {
					logger.Debugf("request authorized by the RoleBinding %s", bindingName)
					return true, nil
				}
				logger.Tracef("%s by rule %+v", reason, rule)
			}
			logger.Debugf("could not authorize the request with the RoleBinding %s",
				bindingName,
			)
		}
		logger.Debugf("could not authorize the request with any RoleBindings")
	}

	logger.Debugf("unauthorized request")
	return false, nil
}

// matchesUser returns whether any of the subjects matches the specified user
func matchesUser(user types.User, subjects []types.Subject) bool {
	for _, subject := range subjects {
		switch subject.Kind {
		case types.UserKind:
			if user.Username == subject.Name {
				return true
			}

		case types.GroupKind:
			for _, group := range user.Groups {
				if group == subject.Name {
					return true
				}
			}
		}
	}

	return false
}

// ruleAllows returns whether the specified rule allows the request based on its
// attributes and if not, the reason why
func ruleAllows(attrs *authorization.Attributes, rule types.Rule) (bool, string) {
	if matches := rule.VerbMatches(attrs.Verb); !matches {
		return false, "forbidden verb"
	}

	if matches := rule.ResourceMatches(attrs.Resource); !matches {
		return false, "forbidden resource"
	}

	if matches := rule.ResourceNameMatches(attrs.ResourceName); !matches {
		return false, "forbidden resource name"
	}

	return true, ""
}