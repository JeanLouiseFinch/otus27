package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
)

type Config struct {
	TypeLog  string `config:"typelog"`
	Host     string `config:"host"`
	Port     string `config:"port"`
	User     string `config:"user"`
	Password string `config:"password"`
	DBName   string `config:"dbname"`
}

func GetConfig(filename string) (*Config, error) {
	var (
		c      Config
		err    error
		loader *confita.Loader
	)
	if filename == "" {
		loader = confita.NewLoader(
			env.NewBackend(),
		)
	} else {
		loader = confita.NewLoader(
			file.NewBackend(filename),
			env.NewBackend(),
		)
	}
	c = Config{}
	err = loader.Load(context.Background(), &c)
	if err != nil {
		panic(err)
	}
	return &c, err
}
