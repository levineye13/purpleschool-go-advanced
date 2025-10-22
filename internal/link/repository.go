package link

import (
	"purpleschool-go/advanced/pkg/db"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: db,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	tx := repo.Database.DB.Create(link)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return link, nil
}

func (repo *LinkRepository) GetAll() ([]Link, error) {
	var links []Link

	tx := repo.Database.DB.Find(&links)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return links, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link

	tx := repo.Database.DB.First(&link, "hash = ?", hash)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &link, nil
}
