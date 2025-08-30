// package business

// import (
// 	"fmt"
// 	"orders/models"
// 	"orders/redis"
// 	"orders/repositories"
// 	"time"
// )

// type PlaceOrderService struct {
// 	repo repositories.IPlaceOrderRepository
// }

// type IPlaceOrderService interface {
// 	CreateOrder(payment *models.PlaceOrderRequest) (*models.PlaceOrderResponse, error)
// }

// func NewPlaceOrderService(repo repositories.IPlaceOrderRepository) IPlaceOrderService {
// 	return &PlaceOrderService{repo: repo}
// }

// func (service *PlaceOrderService) CreateOrder(placeorder *models.PlaceOrderRequest) (*models.PlaceOrderResponse, error) {

// 	exists, err := service.repo.SchemeExists(placeorder.SchemeCode)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to check scheme: %w", err)
// 	}
// 	if !exists {
// 		return nil, fmt.Errorf("invalid scheme_code: %s", placeorder.SchemeCode)
// 	}

// 	nav, err := redis.GetNAV(placeorder.SchemeCode)
// 	if err != nil || nav <= 0 {
// 		return nil, fmt.Errorf("failed to fetch NAV from Redis for %s: %w", placeorder.SchemeCode, err)
// 	}
// 	units := placeorder.Amount / nav
// 	if placeorder.Side == "SELL" {
// 		currentUnits, err := service.repo.GetHoldingUnits(placeorder.UserID, placeorder.SchemeCode)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get holdings: %w", err)
// 		}
// 		if currentUnits < units {
// 			return nil, fmt.Errorf("insufficient units to sell: have %.2f, trying to sell %.2f", currentUnits, units)
// 		}
// 	}

// 	order := models.Order{
// 		UserID:      placeorder.UserID,
// 		SchemeCode:  placeorder.SchemeCode,
// 		Side:        placeorder.Side,
// 		Amount:      placeorder.Amount,
// 		Status:      "pending", // default status when placing
// 		Units:       units,
// 		Nav_Used:    nav,
// 		ContractURL: fmt.Sprintf("http://contracts.example.com/%s_%d", placeorder.UserID, time.Now().Unix()),
// 		PlacedAt:    time.Now(),
// 		ConfirmedAt: nil, // not confirmed yet
// 	}
// 	resp, err := service.repo.CreateOrder(&order)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if order.Side == "BUY" {
// 		existingUnits, err := service.repo.GetHoldingUnits(order.UserID, order.SchemeCode)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get current holdings: %w", err)
// 		}
// 		if existingUnits > 0 {
// 			err = service.repo.UpdateHolding(order.UserID, order.SchemeCode, existingUnits+order.Units)
// 		} else {
// 			err = service.repo.CreateHolding(order.UserID, order.SchemeCode, order.Units)
// 		}
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to update holdings: %w", err)
// 		}
// 	} else if order.Side == "SELL" {
// 		// Subtract sold units from holdings
// 		existingUnits, err := service.repo.GetHoldingUnits(order.UserID, order.SchemeCode)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get current holdings: %w", err)
// 		}
// 		newUnits := existingUnits - order.Units
// 		err = service.repo.UpdateHolding(order.UserID, order.SchemeCode, newUnits)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to update holdings after sell: %w", err)
// 		}
// 	}

//		return resp, nil
//	}
package business

import (
	"fmt"
	"log"
	"orders/kafka"
	"orders/models"
	"orders/redis"
	"orders/repositories"
	"time"
)

type IPlaceOrderService interface {
	PlaceOrder(req *models.PlaceOrderRequest) (*models.PlaceOrderResponse, error)
}

type PlaceOrderService struct {
	repo     repositories.IPlaceOrderRepository
	producer *kafka.OrderProducer
}

func NewPlaceOrderService(repo repositories.IPlaceOrderRepository, producer *kafka.OrderProducer) IPlaceOrderService {
	return &PlaceOrderService{repo: repo, producer: producer}
}

func (service *PlaceOrderService) PlaceOrder(placeorder *models.PlaceOrderRequest) (*models.PlaceOrderResponse, error) {

	exists, err := service.repo.SchemeExists(placeorder.SchemeCode)
	if err != nil {
		return nil, fmt.Errorf("failed to check scheme: %w", err)
	}
	if !exists {
		return nil, fmt.Errorf("invalid scheme_code: %s", placeorder.SchemeCode)
	}

	nav, err := redis.GetNAV(placeorder.SchemeCode)
	if err != nil || nav <= 0 {
		return nil, fmt.Errorf("NAV not available for scheme: %s", placeorder.SchemeCode)
	}
	units := placeorder.Amount / nav
	if placeorder.Side == "SELL" {
		currentUnits, err := service.repo.GetHoldingUnits(placeorder.UserID, placeorder.SchemeCode)
		if err != nil {
			return nil, fmt.Errorf("failed to get holdings: %w", err)
		}
		if currentUnits < units {
			return nil, fmt.Errorf("insufficient units to sell: have %f, trying to sell %f", currentUnits, units)
		}
	}

	order := models.Order{
		UserID:      placeorder.UserID,
		SchemeCode:  placeorder.SchemeCode,
		Side:        placeorder.Side,
		Amount:      placeorder.Amount,
		Status:      "placed",
		Units:       units,
		Nav_Used:    nav,
		ContractURL: fmt.Sprintf("http://contracts.example.com/%s_%d", placeorder.UserID, time.Now().Unix()),
		PlacedAt:    time.Now(),
		ConfirmedAt: nil,
	}
	resp, err := service.repo.CreateOrder(&order)
	if err != nil {
		return nil, err
	}

	err = service.producer.SendOrder(&order)
	if err != nil {
		log.Printf("Failed to send Kafka message: %v\n", err)
	}
	if order.Side == "BUY" {
		existingUnits, err := service.repo.GetHoldingUnits(order.UserID, order.SchemeCode)
		if err != nil {
			return nil, fmt.Errorf("failed to get current holdings: %w", err)
		}
		if existingUnits > 0 {
			err = service.repo.UpdateHolding(order.UserID, order.SchemeCode, existingUnits+order.Units)
		} else {
			err = service.repo.CreateHolding(order.UserID, order.SchemeCode, order.Units)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to update holdings: %w", err)
		}
	} else if order.Side == "SELL" {
		existingUnits, err := service.repo.GetHoldingUnits(order.UserID, order.SchemeCode)
		if err != nil {
			return nil, fmt.Errorf("failed to get current holdings: %w", err)
		}
		newUnits := existingUnits - order.Units
		err = service.repo.UpdateHolding(order.UserID, order.SchemeCode, newUnits)
		if err != nil {
			return nil, fmt.Errorf("failed to update holdings after sell: %w", err)
		}
	}

	return resp, nil
}
