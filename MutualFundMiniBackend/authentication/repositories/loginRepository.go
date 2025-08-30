package repositories

import (
	"authentication/models"

	"gorm.io/gorm"
)

type ILoginRepository interface {
	CreateUser(user *models.LoginRequestModel) error
}
type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) ILoginRepository {
	return &LoginRepository{db: db}
}

func (repo *LoginRepository) CreateUser(user *models.LoginRequestModel) error {
	return repo.db.Create(user).Error
}
