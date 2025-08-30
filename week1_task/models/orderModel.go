package models

import "errors"

type Order struct {
	CommonsModel
	UserID      uint    `json:"UserID"`
	Status      string  `json:"status"`
	Total_cents float32 `json:"total_cents"`
}

func (Order) TableName() string {
	return "Order"
}
func (o *Order) Validate() error {
	if o.UserID <= 0 {
		return errors.New("invalid user id")
	}
	if o.Total_cents <= 0 {
		return errors.New("invalid Cents")
	}
	return nil
}
