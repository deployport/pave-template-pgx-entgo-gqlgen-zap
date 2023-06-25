package configurations

import (
	loggingconfig "github.com/deployport/pavement/logging/config"
	dbconfigs "github.com/gitpushy/pave/cmd/pave/templates/pkg/internal/databases/config"
	webconfig "github.com/gitpushy/pave/cmd/pave/templates/pkg/web/config"
)

// Config is the configuration for the program
type Config struct {
	Databases dbconfigs.Databases  `yaml:"databases"`
	Logging   loggingconfig.Config `yaml:"logging"`
	Web       webconfig.Web        `yaml:"web"`
}
