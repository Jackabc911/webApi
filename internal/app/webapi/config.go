package webapi

import "github.com/Jackabc911/webApi/storage"

// Base config for Web API
type Config struct {
	// server port
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Storage  *storage.Config
}

// Default config for Web API
func NewConfig() *Config {
	return &Config{
		// server port
		BindAddr: ":8080",
		LogLevel: "debug",
		Storage:  storage.NewConfig(),
	}
}
