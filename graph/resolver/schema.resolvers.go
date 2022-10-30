package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/KazukiHayase/server-template/graph/generated"
	"github.com/KazukiHayase/server-template/model"
)

// Sample is the resolver for the sample field.
func (r *queryResolver) Sample(ctx context.Context) (model.Sample, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
