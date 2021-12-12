package config

import (
	"time"

	"github.com/jlgallego99/OSTfind/internal/server"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Config struct {
}

func (cfg *Config) ETCD() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	return cli, err
}

func (cfg *Config) HTTP() (*server.Config, error) {
	return &server.Config{}, nil
}
