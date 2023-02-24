package initializers

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"os"
)

type EnvironmentConf struct {
	Port       string `env:"PORT" envDefault:"8080"`
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbName     string `env:"DB_NAME" envDefault:"postgres"`
	DbUsername string `env:"DB_USERNAME" envDefault:"postgres"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DbSync     bool   `env:"DB_SYNC" envDefault:"true"`
}

var Env *EnvironmentConf

func loadEnvVariables() error {
	var envCfg = &EnvironmentConf{}

	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err := env.Parse(envCfg); err != nil {
		return err
	}

	Env = envCfg
	return nil
}
