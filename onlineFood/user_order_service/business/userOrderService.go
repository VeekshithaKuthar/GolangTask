package business

import (
	"time"
	"userorders/models"
	"userorders/repository"

	"github.com/google/uuid"
)

type IOrderService interface {
	CreateOrder(order *models.UserOrders) (*models.UserOrders, error)
}
type OrderService struct {
	repo repository.IOrderRepository
}

func NewOrderService(repo repository.IOrderRepository) IOrderService {
	return &OrderService{repo: repo}
}
func (service *OrderService) CreateOrder(order *models.UserOrders) (*models.UserOrders, error) {
	order.OrderID = "ORD-" + uuid.NewString()
	order.Status = "PLACED"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order, err := service.repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
