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

func LoadConfig(mode string) (*Config, error) {
	var envPath string

	switch mode {
	case "test":
		envPath = "../.env.test"
	default:
		envPath = ".env.dev"
	}

	err := godotenv.Load(envPath)

	if err != nil {
		return nil, err
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}, nil
}
