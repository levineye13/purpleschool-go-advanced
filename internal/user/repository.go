package user

import "purpleschool-go/advanced/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	tx := repo.Database.DB.Create(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User

	tx := repo.Database.DB.First(&user, "email = ?", email)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
