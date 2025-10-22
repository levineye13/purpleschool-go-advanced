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

func (link *Link) GenerateHash() {
	link.Hash = generateHash(6)
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}

	link.Hash = generateHash(6)

	return link
}
