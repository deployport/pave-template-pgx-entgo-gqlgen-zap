package implementation

import (
	"context"

	"github.com/gitpushy/pave/cmd/pave/templates/pkg/web/graph/generated/model"
)

// CreateProduct is the resolver for the createProduct field.
func (resolver *Resolver) CreateProduct(ctx context.Context, input model.NewProductInput) (*model.Product, error) {
	r, err := resolver.MainDB.Client.Product.Create().SetID(input.ID).SetName(input.Name).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Product{
		ID:   r.ID,
		Name: r.Name,
	}, nil
}
