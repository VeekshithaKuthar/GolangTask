package models

import (
	"encoding/json"
	"strings"
	"time"
	"userorders/constants"
)

type UserOrders struct {
	OrderID      string    `json:"order_id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name"`
	Amount       float64   `json:"amount"`
	Address      string    `json:"address"`
	Item         string    `json:"item"`
	Size         string    `json:"size"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (request *UserOrders) Validate() error {
	if strings.TrimSpace(request.CustomerName) == "" {
		return constants.ErrInvalidCustomerName
	}
	if strings.TrimSpace(request.Address) == "" {
		return constants.ErrInvalidAddress
	}
	if strings.TrimSpace(request.Item) == "" {
		return constants.ErrInvalidItem
	}

	// validSizes := map[string]bool{
	// 	"small":  true,
	// 	"medium": true,
	// 	"large":  true,
	// }
	// if !validSizes[strings.ToLower(request.Size)] {
	// 	return constants.ErrInvalidSize
	// }

	if request.Size != "small" && request.Size != "medium" && request.Size != "large" {
		return constants.ErrInvalidSize
	}

	return nil
}

func (u *UserOrders) ToBytes() []byte {
	bytes, _ := json.Marshal(u)
	return bytes
}
