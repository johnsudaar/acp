package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	URL string
}

func getViperConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("acp.yml")
	v.AddConfigPath("$HOME/.config/acp")
	v.AddConfigPath("/etc/acp")

	v.SetDefault("database.url", "mongodb://localhost:2701")
	v.BindEnv("database.url")

	v.SetDefault("server.port", "8081")
	v.BindEnv("server.port")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return v
}

func Init() error {
	v := getViperConfig()
	err := v.ReadInConfig()
	// If the error is not a file not found error: ignore
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	err = v.Unmarshal(&config)
	if err != nil {
		return err
	}

	os.Setenv("MONGO_URL", config.Database.URL)
	return nil
}

func Get() Config {
	return config
}
