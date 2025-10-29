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

func (repo *LinkRepository) GetAll(limit, offset int) []Link {
	var links []Link

	repo.Database.DB.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)

	return links
}

func (repo *LinkRepository) Count() int64 {
	var count int64

	repo.Database.DB.
		Table("links").
		Where("deleted_at is null").
		Count(&count)

	return count
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
