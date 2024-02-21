package web

import (
	"context"

	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/configurations"
	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/web"
	"github.com/spf13/cobra"
	"go.deployport.com/pavement/logging"
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
