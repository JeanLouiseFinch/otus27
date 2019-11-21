package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
)

type Config struct {
	TypeLog string
}

func GetConfig() (*Config, error) {
	var (
		c      *Config
		err    error
		loader *confita.Loader
	)
	loader = confita.NewLoader(
		file.NewBackend("confita.yaml"),
	)
	c = &Config{}
	err = loader.Load(context.Background(), c)
	if err != nil {
		panic(err)
	}
	return c, err
}
