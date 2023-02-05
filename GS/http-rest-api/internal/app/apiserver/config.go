package apiserver

// Config объявление структуры
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
	//Store    *store.ConfigStore `toml:"Store"`
}

// NewConfig иннициализация структуры
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8000",
		LogLevel:    "error",
		DatabaseURL: "host=localhost port=5432 user=postgres password=domcrat dbname=restapi_dev sslmode=disable",
		//Store:    store.NewConfig(),
	}
}
