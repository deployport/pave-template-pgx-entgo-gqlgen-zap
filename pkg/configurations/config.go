package configurations

import (
	dbconfigs "github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/internal/databases/config"
	webconfig "github.com/deployport/pave-template-pgx-entgo-gqlgen-zap/pkg/web/config"
	loggingconfig "github.com/deployport/pavement/logging/config"
)

// Config is the configuration for the program
type Config struct {
	Databases dbconfigs.Databases  `yaml:"databases"`
	Logging   loggingconfig.Config `yaml:"logging"`
	Web       webconfig.Web        `yaml:"web"`
}
