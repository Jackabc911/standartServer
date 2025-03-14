package apiserver

import "github.com/Jackabc911/standartServer/store"

//General config for rest api
type Config struct {
	//Port for start api
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

//Should return default config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
