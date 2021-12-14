package config

import (
	"os"
	"time"

	"github.com/jlgallego99/OSTfind/internal/server"
	"github.com/joho/godotenv"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Config struct {
}

func (cfg *Config) ETCD() (*clientv3.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD_ENDPOINT")},
		DialTimeout: 5 * time.Second,
	})

	return cli, err
}

func (cfg *Config) HTTP() (*server.Config, error) {
	return &server.Config{}, nil
}
