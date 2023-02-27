package initializers

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"os"
)

type environmentConf struct {
	Port        string `env:"PORT" envDefault:"8080"`
	DbHost      string `env:"DB_HOST" envDefault:"localhost"`
	DbPort      string `env:"DB_PORT" envDefault:"5432"`
	DbName      string `env:"DB_NAME" envDefault:"postgres"`
	DbUsername  string `env:"DB_USERNAME" envDefault:"postgres"`
	DbPassword  string `env:"DB_PASSWORD" envDefault:"postgres"`
	DbSync      bool   `env:"DB_SYNC" envDefault:"true"`
	ReleaseMode bool   `env:"RELEASE_MODE" envDefault:"false"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"debug"`
	LogFile     string `env:"LOG_FILE" envDefault:"gin.log"`
}

var envConf *environmentConf

func loadEnvVariables() error {
	var envCfg = &environmentConf{}

	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if err := env.Parse(envCfg); err != nil {
		return err
	}

	envConf = envCfg
	return nil
}
