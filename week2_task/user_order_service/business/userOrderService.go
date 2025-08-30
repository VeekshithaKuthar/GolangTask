package business

import (
	"time"
	"userorders/models"
	"userorders/repositories"

	"github.com/google/uuid"
)

type IOrderService interface {
	CreateOrder(order *models.UserOrders) (string, error)
}
type OrderService struct {
	repo repositories.IOrderRepository
}

func NewOrderService(repo repositories.IOrderRepository) IOrderService {
	return &OrderService{repo: repo}
}
func (service *OrderService) CreateOrder(order *models.UserOrders) (string, error) {
	order.OrderID = "ORD-" + uuid.NewString()
	order.Status = "PLACED"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	if err := service.repo.CreateOrder(order); err != nil {
		return "", err
	}
	return order.OrderID, nil
}
