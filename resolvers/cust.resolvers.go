package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"
)

func (r *mutationResolver) AddCustomer(ctx context.Context, input ent.MstCustomerInput) (*ent.MstCustomer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListCustomer(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MstCustomerOrder) (*ent.MstCustomerConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns myeduate.MutationResolver implementation.
func (r *Resolver) Mutation() myeduate.MutationResolver { return &mutationResolver{r} }

// Query returns myeduate.QueryResolver implementation.
func (r *Resolver) Query() myeduate.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
