package rule

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/sahidrahman404/gigachad-api/ent/privacy"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/viewer"
)

var (
	InactiveAccountErr     = privacy.Denyf("your user account must be activated to access this resource")
	AuthenticationRequired = privacy.Denyf("you must be authenticated to access this resource")
)

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing")
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

func FilterRoutineRule() privacy.QueryMutationRule {
	type UsersFilter interface {
		WhereUserID(p entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		userCtx, ok := ctx.Value(types.UserContextKey).(*types.User)
		if !ok || userCtx.IsAnonymous() {
			return AuthenticationRequired
		}

		uf, ok := f.(UsersFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		uf.WhereUserID(entql.StringEQ(string(userCtx.Ent.ID)))
		return privacy.Skip
	})
}

func FilterExerciseRule() privacy.QueryMutationRule {
	type UsersFilter interface {
		WhereUserID(p entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		userCtx, ok := ctx.Value(types.UserContextKey).(*types.User)
		if !ok || userCtx.IsAnonymous() {
			return AuthenticationRequired
		}

		if userCtx.Ent.Activated == 0 {
			return InactiveAccountErr
		}

		uf, ok := f.(UsersFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		if userCtx.Ent != nil {
			uf.WhereUserID(entql.StringOr(entql.StringEQ(string(userCtx.Ent.ID)), entql.StringNil()))
			return privacy.Skip
		}

		uf.WhereUserID(entql.StringNil())
		return privacy.Skip
	})
}

func AllowEditExerciseIfOwner() privacy.MutationRule {
	type UsersFilter interface {
		WhereUserID(p entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		userCtx, ok := ctx.Value(types.UserContextKey).(*types.User)
		if !ok || userCtx.IsAnonymous() {
			return AuthenticationRequired
		}

		if userCtx.Ent.Activated == 0 {
			return InactiveAccountErr
		}

		uf, ok := f.(UsersFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		uf.WhereUserID(entql.StringEQ(string(userCtx.Ent.ID)))

		return privacy.Skip
	})
}
