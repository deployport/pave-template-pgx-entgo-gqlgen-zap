package implementation

import (
	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/internal/databases/maindb"
)

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	MainDB maindb.Client
}
