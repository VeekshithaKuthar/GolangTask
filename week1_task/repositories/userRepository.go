package repositories

import (
	"ordersApi/models"

	"gorm.io/gorm"
)

type UserDb struct {
	DB *gorm.DB
}

type IuserDB interface {
	Create(user *models.User) (*models.User, error)
	GetBy(id uint) (*models.User, error)
}

func NewUserDB(db *gorm.DB) IuserDB {
	return &UserDb{DB: db}
}

func (udb *UserDb) Create(user *models.User) (*models.User, error) {
	tx := udb.DB.Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (udb *UserDb) GetBy(id uint) (*models.User, error) {
	user := new(models.User)
	tx := udb.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}
