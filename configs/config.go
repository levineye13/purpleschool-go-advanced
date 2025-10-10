package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Dsn string
}

type Config struct {
	Db DbConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}, nil
}
