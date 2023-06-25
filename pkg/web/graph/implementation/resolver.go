package implementation

import (
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/databases/maindb"
)

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	MainDB maindb.Client
}
