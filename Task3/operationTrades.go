package main

import (
	"fmt"
	"log"
	"trades/models"

	"gorm.io/gorm"
)

type Trades interface {
	Create(user *models.TradesModel) (*models.TradesModel, error)
}

type Position struct {
	Shares        int
	NetInvestment float64
}
type UserDb struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDb {
	return &UserDb{db}
}

func (udb *UserDb) AddTrade(trades *[]models.TradesModel) (*[]models.TradesModel, error) {
	tx := udb.DB.Create(trades)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return trades, nil
}

func (udb *UserDb) GetNetPosition() error {
	var trades []models.TradesModel
	if err := udb.DB.Find(&trades).Error; err != nil {
		return err
	}

	positions := make(map[string]Position)

	for _, t := range trades {
		pos := positions[t.Symbol]
		switch t.Action {
		case "BUY":
			pos.Shares += int(t.Quantity)
			pos.NetInvestment += float64(t.Quantity) * t.Price
		case "SELL":
			pos.Shares -= int(t.Quantity)
			pos.NetInvestment -= float64(t.Quantity) * t.Price
		default:
			log.Printf("Unknown trade action:")
		}

		positions[t.Symbol] = pos
	}

	for symbol, pos := range positions {
		fmt.Printf("%s: %d shares, Net Investment: ₹%.2f\n", symbol, pos.Shares, pos.NetInvestment)
	}
	return nil
}
