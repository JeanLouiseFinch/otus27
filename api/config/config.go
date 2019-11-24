package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
)

type ConfigLog struct {
	TypeLog string `config:"typelog"`
}

type ConfigRMQ struct {
	HostRMQ     string `config:"hostrmq"`
	UserRMQ     string `config:"userrmq"`
	PortRMQ     string `config:"portrmq"`
	PasswordRMQ string `config:"passwordrmq"`
	QueueRMQ    string `config:"queuermq"`
	TimeoutRMQ  int    `config:"timeoutrmq"`
	DurationRMQ int    `config:"durationrmq"`
}

type ConfigDB struct {
	HostDB     string `config:"host"`
	PortDB     string `config:"port"`
	UserDB     string `config:"user"`
	PasswordDB string `config:"password"`
	NameDB     string `config:"dbname"`
}

type Config struct {
	Log *ConfigLog
	RMQ *ConfigRMQ
	DB  *ConfigDB
}

func GetConfig(filename string) (*Config, error) {
	c := Config{}
	db, err := GetConfigDB(filename)
	if err != nil {
		return nil, err
	}
	log, err := GetConfigLog(filename)
	if err != nil {
		return nil, err
	}
	rmq, err := GetConfigRMQ(filename)
	if err != nil {
		return nil, err
	}
	c.DB, c.Log, c.RMQ = db, log, rmq
	return &c, err
}
func GetConfigRMQ(filename string) (*ConfigRMQ, error) {
	var (
		c      ConfigRMQ
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
	c = ConfigRMQ{}
	err = loader.Load(context.Background(), &c)
	if err != nil {
		panic(err)
	}
	return &c, err
}
func GetConfigDB(filename string) (*ConfigDB, error) {
	var (
		c      ConfigDB
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
	c = ConfigDB{}
	err = loader.Load(context.Background(), &c)
	if err != nil {
		panic(err)
	}
	return &c, err
}
func GetConfigLog(filename string) (*ConfigLog, error) {
	var (
		c      ConfigLog
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
	c = ConfigLog{}
	err = loader.Load(context.Background(), &c)
	if err != nil {
		panic(err)
	}
	return &c, err
}
