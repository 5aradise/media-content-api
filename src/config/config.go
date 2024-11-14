package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		DB     DB
		Server Server
	}

	DB struct {
		URL string `envconfig:"DATABASE_URL" required:"true"`
	}

	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
	}
)

var Cfg Config

func Load() error {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		err := godotenv.Load(configPath)
		if err != nil {
			return fmt.Errorf("cannot load from config file: %w", err)
		}
	}

	err := envconfig.Process("", &Cfg)
	if err != nil {
		return fmt.Errorf("cannot set config: %s", err.Error())
	}
	return nil
}
