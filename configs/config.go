package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

type Config struct {
	Db   DbConfig
	Auth AuthConfig
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
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}, nil
}
