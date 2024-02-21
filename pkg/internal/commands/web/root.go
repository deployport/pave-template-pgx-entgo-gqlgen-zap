package web

import (
	"context"

	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/configurations"
	"github.com/spf13/cobra"
	"go.deployport.com/pavement/logging"
)

// BuildRootCommand creates the root command
func BuildRootCommand(
	ctx context.Context,
	cfg *configurations.Config,
	logger *logging.Logger,
) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "web",
		Short: "Run the mecha web server",
	}
	rootCmd.AddCommand(serverCommand(
		ctx,
		cfg,
		logger,
	))
	return rootCmd
}
