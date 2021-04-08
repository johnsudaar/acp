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
	Port       int
	AssetsPath string `mapstructure:"assets_path"`
}

type DatabaseConfig struct {
	URL string
}

func GetViperConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigFile("acp.yml")
	v.AddConfigPath("$HOME/.config/acp")
	v.AddConfigPath("/etc/acp")
	v.AddConfigPath(".")

	v.SetDefault("database.url", "mongodb://localhost:27017")
	v.BindEnv("database.url")

	v.SetDefault("server.port", "8081")
	v.BindEnv("server.port")
	v.SetDefault("server.assets_path", "front/")
	v.BindEnv("server.assets_path")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return v
}

func Init() error {
	v := GetViperConfig()
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
