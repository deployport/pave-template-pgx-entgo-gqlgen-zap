package graph

import (
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/web/graph/implementation"
)

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	implementation.Resolver
}
