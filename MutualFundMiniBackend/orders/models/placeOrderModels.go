package models

import (
	"errors"
	"time"
)

type PlaceOrderRequest struct {
	UserID     string  `json:"user_id"`
	SchemeCode string  `json:"scheme_code"`
	Side       string  `json:"side"`
	Amount     float64 `json:"amount"`
}

type PlaceOrderResponse struct {
	UserID  string `json:"user_id"`
	OrderID uint   `json:"order_id"`
	Status  string `json:"status"`
}

type Order struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      string     `json:"user_id"`
	SchemeCode  string     `json:"scheme_code"`
	Side        string     `json:"side"`
	Amount      float64    `json:"amount"`
	Units       float64    `json:"units"`
	Status      string     `json:"status"`
	Nav_Used    float64    `json:"nav_used"`
	PlacedAt    time.Time  `json:"placed_at"`
	ConfirmedAt *time.Time `json:"confirmed_at"`
	ContractURL string     `gorm:"column:contract_url"`
}

type Holding struct {
	UserID     string  `gorm:"primaryKey;column:user_id" json:"user_id"`
	SchemeCode string  `gorm:"primaryKey;column:scheme_code" json:"scheme_code"`
	Units      float64 `gorm:"column:units" json:"units"`
}

func (req *PlaceOrderRequest) Validate() error {
	if req.UserID == "" {
		return errors.New("user_id is required")
	}
	if req.SchemeCode == "" {
		return errors.New("scheme_code is required")
	}
	if req.Side != "BUY" && req.Side != "SELL" {
		return errors.New("side must be BUY or SELL")
	}
	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	return nil
}
