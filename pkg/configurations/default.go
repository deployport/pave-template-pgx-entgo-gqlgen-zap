package configurations

import (
	_ "embed" // embed default config file

	"go.deployport.com/pavement/config"
)

// default variable bytes loaded using embed
//
//go:embed default.yaml
var defaultBytes []byte

// New is the root config
func New() *config.Backed[Config] {
	return config.NewBacked[Config](defaultBytes)
}
