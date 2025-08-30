// package models

// import (
// 	"errors"
// 	"time"
// )

// type Payments struct {
// 	PatmentId     uint      `json:"payment_id" gorm:"autoIncrement;primaryKey"`
// 	PaymentAmount float64   `json:"payment_amount" `
// 	PaymentStatus string    `json:"payment_status"`
// 	Created_at    time.Time `json:"created_at" gorm:"autoCreateTime"`
// 	Updated_at    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
// 	Order_id      uint      `json:"order_id"`
// }

// type UserOrders struct {
// 	OrderID      string  `json:"order_id"`
// 	UserID       int     `json:"user_id"`
// 	CustomerName string  `json:"customer_name"`
// 	Amount       float64 `json:"amount"`
// 	Status       string  `json:"status"`
// }

// func (Payments) TableName() string {
// 	return "Payments"
// }

//	func (p *Payments) Validate() error {
//		if p.Order_id == 0 {
//			return errors.New("invalid order id")
//		}
//		if p.PaymentAmount <= 0 {
//			return errors.New("invalid payment amount")
//		}
//		return nil
//	}
package models

import (
	"errors"
	"time"
)

type Payments struct {
	PaymentId     uint      `json:"payment_id" gorm:"column:payment_id;primaryKey;autoIncrement"`
	OrderID       string    `json:"order_id" gorm:"column:order_id;not null"`
	PaymentAmount float64   `json:"payment_amount" gorm:"column:payment_amount"`
	PaymentStatus string    `json:"payment_status" gorm:"column:payment_status"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Payments) TableName() string {
	return "payments"
}

type UserOrders struct {
	OrderID string `json:"order_id"`
	// UserID       int     `json:"user_id"`
	CustomerName string  `json:"customer_name"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
}

func (p *Payments) Validate() error {
	if p.OrderID == "" {
		return errors.New("invalid order id")
	}
	if p.PaymentAmount <= 0 {
		return errors.New("invalid payment amount")
	}
	return nil
}
