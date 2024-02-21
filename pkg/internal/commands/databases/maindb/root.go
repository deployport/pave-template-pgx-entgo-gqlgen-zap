package maindb

import (
	"context"

	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/internal/databases/maindb/migrations"
	"github.com/spf13/cobra"
	dbcmd "go.deployport.com/pavement/commands/database"
	"go.deployport.com/pavement/logging"
	sqlconfig "go.deployport.com/pavement/sql/config"
)

// BuildRootCommand creates the root command
func BuildRootCommand(
	ctx context.Context,
	dbconfig *sqlconfig.Connection,
	logger *logging.Logger,
) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "maindb",
	}
	rootCmd.AddCommand(dbcmd.Build(ctx, dbcmd.RootParams{
		NewCatalog: migrations.NewCatalog,
		Connection: dbconfig,
		Logger:     logger,
	})...)
	return rootCmd
}
