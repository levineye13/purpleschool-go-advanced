package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"link"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateHash(length int) string {
	hash := make([]rune, length)

	for index := range length {
		hash[index] = charset[rand.Intn(len(charset))]
	}

	return string(hash)
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: generateHash(6),
	}
}
