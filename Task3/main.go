package main

import (
	"fmt"
	"os"
	"time"
	"trades/models"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	DSN   string
	PORT  string
	debug bool
)

func main() {

	DSN = os.Getenv("DSN")
	if DSN == "" {
		DSN = `host=localhost user=app password=app123 dbname=usersdb port=5433 sslmode=disable`
		log.Info().Msg(DSN)
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	db, err := GetConnection(DSN)

	if err != nil {
		fmt.Print(err)
	}
	Init(db)

	trades := []models.TradesModel{
		{Symbol: "INFY", Action: "BUY", Quantity: 100, Price: 1400, LastModified: time.Now().Unix()},
		{Symbol: "TCS", Action: "SELL", Quantity: 50, Price: 3200, LastModified: time.Now().Unix()},
	}
	udb := NewUserDB(db)
	if _, err := udb.AddTrade(&trades); err != nil {
		fmt.Println("Error in adding trade:", err)
	}

	udb.GetNetPosition()

}

func Init(db *gorm.DB) {
	db.AutoMigrate(&models.TradesModel{})
}
