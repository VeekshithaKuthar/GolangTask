package repositories

import (
	"orders/models"

	"gorm.io/gorm"
)

type IOrderBookRepository interface {
	FetchOrders(userID string) ([]models.Order, error)
}

type OrderBookRepository struct {
	db *gorm.DB
}

func NewOrderBookRepository(db *gorm.DB) IOrderBookRepository {
	return &OrderBookRepository{db: db}
}

func (r *OrderBookRepository) FetchOrders(userID string) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
