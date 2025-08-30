package business

import (
	"orders/models"
	"orders/repositories"
)

type IOrderBookService interface {
	GetOrders(userID string) ([]models.OrderBookResponse, error)
}

type OrderBookService struct {
	repo repositories.IOrderBookRepository
}

func NewOrderBookService(repo repositories.IOrderBookRepository) IOrderBookService {
	return &OrderBookService{repo: repo}
}

func (service *OrderBookService) GetOrders(userID string) ([]models.OrderBookResponse, error) {
	orders, err := service.repo.FetchOrders(userID)
	if err != nil {
		return nil, err
	}
	var resp []models.OrderBookResponse
	for _, ord := range orders {
		resp = append(resp, models.OrderBookResponse{
			ID:          ord.ID,
			SchemeCode:  ord.SchemeCode,
			Side:        ord.Side,
			Amount:      ord.Amount,
			Units:       ord.Units,
			Status:      ord.Status,
			ContractURL: ord.ContractURL,
		})
	}
	return resp, nil
}
