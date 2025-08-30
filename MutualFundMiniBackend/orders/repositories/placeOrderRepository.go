package repositories

import (
	"errors"
	"orders/models"
	"time"

	"gorm.io/gorm"
)

type IPlaceOrderRepository interface {
	CreateOrder(order *models.Order) (*models.PlaceOrderResponse, error)
	SchemeExists(schemeCode string) (bool, error)
	GetHoldingUnits(userID string, schemeCode string) (float64, error)
	CreateHolding(userID string, schemeCode string, units float64) error
	UpdateHolding(userID string, schemeCode string, newUnits float64) error
}
type PlaceOrderRepository struct {
	db *gorm.DB
}

func NewPlaceOrderRepository(db *gorm.DB) IPlaceOrderRepository {
	return &PlaceOrderRepository{db: db}
}

func (repo *PlaceOrderRepository) CreateOrder(order *models.Order) (*models.PlaceOrderResponse, error) {
	now := time.Now()
	order.ConfirmedAt = &now
	result := repo.db.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &models.PlaceOrderResponse{
		UserID:  order.UserID,
		OrderID: order.ID,
		Status:  order.Status,
	}
	return resp, nil

}

func (repo *PlaceOrderRepository) SchemeExists(schemeCode string) (bool, error) {
	var count int64
	err := repo.db.Table("schemes").Where("scheme_code = ?", schemeCode).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (repo *PlaceOrderRepository) GetHoldingUnits(userID, schemeCode string) (float64, error) {
	var holding models.Holding
	result := repo.db.Where("user_id = ? AND scheme_code = ?", userID, schemeCode).First(&holding)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return 0, nil // not an error; just means no holding
	}
	if result.Error != nil {
		return 0, result.Error
	}
	return holding.Units, nil
}

func (repo *PlaceOrderRepository) CreateHolding(userID string, schemeCode string, units float64) error {
	holding := models.Holding{
		UserID:     userID,
		SchemeCode: schemeCode,
		Units:      units,
	}

	if err := repo.db.Create(&holding).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PlaceOrderRepository) UpdateHolding(userID string, schemeCode string, newUnits float64) error {
	tx := repo.db.Model(&models.Holding{}).
		Where("user_id = ? AND scheme_code = ?", userID, schemeCode).
		Update("units", newUnits)
	return tx.Error
}
