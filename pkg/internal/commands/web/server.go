package web

import (
	"context"

	"github.com/deployport/pavement/logging"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/configurations"
	"github.com/gitpushy/pave/cmd/pave/templates/pkg/web"
	"github.com/spf13/cobra"
)

func serverCommand(
	ctx context.Context,
	cfg *configurations.Config,
	logger *logging.Logger,
) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run the mecha web server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return web.RunServer(
				ctx,
				cfg,
				logger,
			)
		},
	}
	return cmd
}
