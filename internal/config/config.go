package config

// Config holds application settings.
type Config struct {
	Port int
}

// NewConfig provides the default Port.
func NewConfig() *Config {
	return &Config{Port: 8080}
}
