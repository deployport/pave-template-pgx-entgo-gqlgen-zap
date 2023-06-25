package config

// Web is the configuration for the web server
type Web struct {
	AllowedOrigins []string `yaml:"allowedOrigins"`
	Port           int      `yaml:"port"`
}
