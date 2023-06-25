package web

import (
	"context"

	"github.com/deployport/pavement/logging"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/configurations"
	"github.com/spf13/cobra"
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
