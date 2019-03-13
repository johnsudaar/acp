package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var config Config

type Config struct {
	Port int `envconfig:"PORT" default:"8081"`
}

func Init() error {
	err := envconfig.Process("", &config)
	if err != nil {
		return errors.Wrap(err, "fail to parse environment")
	}
	return nil
}

func Get() Config {
	return config
}
