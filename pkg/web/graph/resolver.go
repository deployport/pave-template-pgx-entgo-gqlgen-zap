package graph

import (
	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/web/graph/implementation"
)

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	implementation.Resolver
}
