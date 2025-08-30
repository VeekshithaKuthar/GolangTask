package models

type TradesModel struct {
	ID           uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Symbol       string  `json:"symbol"`
	Action       string  `json:"action"`
	Quantity     uint    `json:"quantity"`
	Price        float64 `json:"price"`
	LastModified int64   `json:"last_modified" gorm:"index"`
}
