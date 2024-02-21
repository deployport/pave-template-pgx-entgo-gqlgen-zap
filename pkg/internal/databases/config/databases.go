package config

import sqlconfig "go.deployport.com/pavement/sql/config"

// Databases is the configuration for the databases
type Databases struct {
	MainDB sqlconfig.Connection `yaml:"maindb"`
}
