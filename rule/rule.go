package rule

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/sahidrahman404/gigachad-api/ent/privacy"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/viewer"
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
		uCtx, ok := ctx.Value(types.UserContextKey).(*types.User)
		if !ok {
			return privacy.Denyf("missing user information in viewer")
		}
		if uCtx.Ent == nil {
			return privacy.Denyf("missing user information in viewer")
		}
		uf, ok := f.(UsersFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		uf.WhereUserID(entql.StringEQ(string(uCtx.Ent.ID)))
		return privacy.Skip
	})
}

func FilterExerciseRule() privacy.QueryMutationRule {
	type UsersFilter interface {
		WhereUserID(p entql.StringP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		uCtx, ok := ctx.Value(types.UserContextKey).(*types.User)
		if !ok {
			return privacy.Denyf("missing user information in viewer")
		}

		uf, ok := f.(UsersFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		if uCtx.Ent != nil {
			uf.WhereUserID(entql.StringOr(entql.StringEQ(string(uCtx.Ent.ID)), entql.StringNil()))
			return privacy.Skip
		}

		uf.WhereUserID(entql.StringNil())
		return privacy.Skip
	})
}
