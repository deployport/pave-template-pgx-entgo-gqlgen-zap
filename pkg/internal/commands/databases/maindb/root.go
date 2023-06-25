package maindb

import (
	"context"

	dbcmd "github.com/deployport/pavement/commands/database"
	"github.com/deployport/pavement/logging"
	sqlconfig "github.com/deployport/pavement/sql/config"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/databases/maindb/migrations"
	"github.com/spf13/cobra"
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
