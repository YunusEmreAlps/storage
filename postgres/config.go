package postgres

import "time"

// Config defines the config for storage.
type Config struct {
	GCInterval time.Duration
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	GCInterval: 10 * time.Second,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if int(cfg.GCInterval) == 0 {
		cfg.GCInterval = ConfigDefault.GCInterval
	}
	return cfg
}