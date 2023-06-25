package maindb

import (
	"context"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	pavementsql "github.com/deployport/pavement/sql"
	sqlconfig "github.com/deployport/pavement/sql/config"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/databases/maindb/ent"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/databases/maindb/migrations"
	_ "github.com/jackc/pgx/v5/stdlib" // register pgx driver
	"go.uber.org/zap"
)

// Open opens a connection to the database.
// This instance can be shared safely across requests
// as connection pooling is handled by the driver.
func Open(ctx context.Context, logger *zap.Logger, config sqlconfig.Connection) (Client, error) {
	logger = logger.Named("maindb")
	client, err := pavementsql.NewClient[*ent.Tx, *ent.Client](
		ctx,
		logger,
		config,
		func(drv *entsql.Driver) *ent.Client {
			return ent.NewClient(ent.Driver(drv))
		},
		migrations.NewCatalog,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create database client, %w", err)
	}
	return client, nil
}

// Client is a wrapper around the ent.Client
type Client *pavementsql.Client[*ent.Tx, *ent.Client]
