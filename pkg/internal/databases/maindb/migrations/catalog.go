package migrations

import (
	"embed"

	sqlconfig "github.com/deployport/pavement/sql/config"
	sqlmigrations "github.com/deployport/pavement/sql/migrations"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // migrate for postgres
	"go.uber.org/zap"
)

//go:embed sql/*.sql
var fs embed.FS

// NewCatalog creates a new Catalog instance
func NewCatalog(logger *zap.Logger, config sqlconfig.Connection) (*sqlmigrations.Catalog, error) {
	logger = logger.Named("migration-catalog")
	c, err := sqlmigrations.NewCatalog(logger, config, fs)
	if err != nil {
		return nil, err
	}
	return c, nil
}
