package link

import "purpleschool-go/advanced/pkg/db"

type LinkRepository struct {
	db *db.Db
}

func NewLinkRepository(db *db.Db) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (repo *LinkRepository) Create(link *Link) {
	repo.db.Create(link)
}
