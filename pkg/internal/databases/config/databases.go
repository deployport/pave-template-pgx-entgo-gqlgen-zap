package config

import sqlconfig "github.com/deployport/pavement/sql/config"

// Databases is the configuration for the databases
type Databases struct {
	MainDB sqlconfig.Connection `yaml:"maindb"`
}
