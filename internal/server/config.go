package server

import "os"

type Config struct {
	Host string
	Port string
}

func (c *Config) Set() {
	c.Host = os.Getenv("HTTP_HOST")
	c.Port = os.Getenv("HTTP_PORT")
}
