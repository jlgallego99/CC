package config

import "github.com/jlgallego99/OSTfind/internal/server"

type Config struct {
}

func (cfg *Config) HTTP() (*server.Config, error) {
	return &server.Config{}, nil
}
