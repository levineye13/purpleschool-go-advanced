package link

import (
	"purpleschool-go/advanced/pkg/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	tx := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	tx := repo.Database.DB.Delete(&Link{}, id)

	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
