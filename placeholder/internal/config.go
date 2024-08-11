package internal

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppName    string `env:"NAME"`
	AppEnv     string `env:"ENV"`
	LogLevel   string `env:"LOG_LEVEL"`
	ServerPort string `env:"SERVER_PORT"`
	Database   DBConfig
}

var globalConfig Config

func init() {
	ctx := context.Background()
	var config Config

	if err := envconfig.Process(ctx, &config); err != nil {
		log.Fatal(err)
	}

	globalConfig = config
}

func (this Config) Validate() error {
	if this.ServerPort == "" {
		return nil
	}

	return nil
}

func GetConfig() Config {
	return globalConfig
}
