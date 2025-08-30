package models

import "errors"

type Order struct {
	CommonModel
	UserID   uint    `json:"user_id"`
	Amount   float32 `json:"amount"`
	ItemName string  `json:"item_name" gorm:"column:item_name"`
}

func (o *Order) Validate() error {
	if o.UserID <= 0 {
		return errors.New("invalid user id")
	}
	if o.ItemName == "" {
		return errors.New("invalid item")
	}
	return nil
}
