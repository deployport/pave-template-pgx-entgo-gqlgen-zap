package commands

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/configurations"
	dbcmd "github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/internal/commands/databases"
	webcmd "github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/internal/commands/web"
	"github.com/spf13/cobra"
	configcmd "go.deployport.com/pavement/commands/config"
	"go.deployport.com/pavement/logging"
)

var rootConfig = configurations.New()
var logger = logging.NewLogger()

var mainProcessConfigFilename = "mainprocess-config.yaml"

var configFilesFlag = []string{
	"$HOME/.mainprocess/config.yaml",
	"/etc/mainprocess/.config.yaml",
	"$PWD/" + mainProcessConfigFilename,
}

// Execute runs the root command
func Execute() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// run runs the root command
func run(ctx context.Context) error {
	var rootCmd = &cobra.Command{
		Use: "mainprocess",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := initializeConfig(cmd); err != nil {
				return err
			}
			return nil
		},
	}

	rootCmd.Flags().StringArrayVarP(&configFilesFlag, "config", "c", configFilesFlag, fmt.Sprintf("config file (default is %s)", strings.Join(configFilesFlag, ",")))

	rootCmd.AddCommand(dbcmd.BuildRootCommand(ctx, &rootConfig.C.Databases, logger))
	rootCmd.AddCommand(configcmd.Root(ctx, configcmd.RootParams[configurations.Config]{
		BackedConfig: rootConfig,
		InitFilename: mainProcessConfigFilename,
	}))

	rootCmd.AddCommand(webcmd.BuildRootCommand(ctx, rootConfig.C, logger))
	return rootCmd.Execute()
}

func initializeConfig(cmd *cobra.Command) error {
	for _, configFile := range configFilesFlag {
		// expand env filename
		filename := os.ExpandEnv(configFile)
		if !fileExists(filename) {
			continue
		}
		if err := rootConfig.OverrideFromFile(filename); err != nil {
			return err
		}
	}
	logger.Configure(rootConfig.C.Logging)

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
