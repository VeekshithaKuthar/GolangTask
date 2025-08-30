package repository

import (
	"userorders/models"

	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(order *models.UserOrders) (*models.UserOrders, error)
	UpdateOrderStatus(orderID string, status string) (*models.UserOrders, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) IOrderRepository {
	return &OrderRepository{db: db}
}

func (repo *OrderRepository) CreateOrder(order *models.UserOrders) (*models.UserOrders, error) {
	txt := repo.db.Create(order)
	if txt.Error != nil {
		return nil, txt.Error
	}
	return order, nil
}

func (repo *OrderRepository) UpdateOrderStatus(orderID string, status string) (*models.UserOrders, error) {
	var order models.UserOrders

	// Find the order by ID
	if err := repo.db.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return nil, err
	}

	// Update status
	order.Status = status
	if err := repo.db.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
