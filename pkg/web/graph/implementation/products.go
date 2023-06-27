package implementation

import (
	"context"

	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/web/graph/generated/model"
)

// Products is the resolver for the products field
func (resolver *Resolver) Products(ctx context.Context) ([]*model.Product, error) {
	q, err := resolver.MainDB.Client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var products []*model.Product
	for _, product := range q {
		products = append(products, &model.Product{
			ID:   product.ID,
			Name: product.Name,
		})
	}
	return products, nil
}
