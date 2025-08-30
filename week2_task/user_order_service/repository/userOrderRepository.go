package repository

import (
	"userorders/models"

	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(order *models.UserOrders) error
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{db: db}
}

func (repo *OrderRepository) CreateOrder(order *models.UserOrders) error {
	return repo.db.Create(order).Error
}
