package main

import (
	"os"
	"purpleschool-go/advanced/internal/link"
	"purpleschool-go/advanced/internal/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&link.Link{}, &user.User{})

	if err != nil {
		panic(err)
	}
}
