package databases

import (
	"context"

	"github.com/deployport/pavement/logging"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/commands/databases/maindb"
	dbconfigs "github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/databases/config"
	"github.com/spf13/cobra"
)

// BuildRootCommand creates the root command
func BuildRootCommand(
	ctx context.Context,
	dbconfig *dbconfigs.Databases,
	logger *logging.Logger,
) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "databases",
	}
	rootCmd.AddCommand(maindb.BuildRootCommand(
		ctx,
		&dbconfig.MainDB,
		logger,
	))
	return rootCmd
}
