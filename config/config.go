package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var config Config

type Config struct {
	Port       int    `envconfig:"PORT" default:"8081"`
	PublicPath string `envconfig:"PUBLIC_PATH" default:"public/"`
	Version    string `envconfig:"VERSION" default:""`
}

type ConfigOpts struct {
	Version string
}

func Init(opts ConfigOpts) error {
	err := envconfig.Process("", &config)
	if err != nil {
		return errors.Wrap(err, "fail to parse environment")
	}
	config.Version = opts.Version
	return nil
}

func Get() Config {
	return config
}
