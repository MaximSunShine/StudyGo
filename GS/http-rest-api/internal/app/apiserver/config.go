package apiserver

import "github.com/GS/http-rest-api/internal/app/store"

// Config объявление структуры
type Config struct {
	BindAddr string             `toml:"bind_addr"`
	LogLevel string             `toml:"log_level"`
	Store    *store.ConfigStore `toml:"Store"`
}

// NewConfig иннициализация структуры
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "error",
		Store:    store.NewConfig(),
	}
}
