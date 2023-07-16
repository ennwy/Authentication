package main

import (
	"github.com/ennwy/auth/internal/server"
	"github.com/ennwy/auth/internal/logger"
)

type Config struct {
	Logger logger.Config `yaml:"logger"`
	HTTP   server.Config `yaml:"http"`
}

func NewConfig() *Config {
	config := &Config{}
	config.Logger.Set()
	config.HTTP.Set()

	return config
}
