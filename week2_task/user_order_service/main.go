package main

import (
	"flag"
	"os"
	"userorders/constants"
	"userorders/database"
	"userorders/models"
	"userorders/router"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	DSN   string
	PORT  string
	debug bool
)

func main() {
	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
	flag.Parse()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	DSN = os.Getenv("DSN")
	if DSN == "" {
		DSN = "host=pg user=app password=app123 dbname=userordersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msgf("unable to connect to the database %s", constants.Service)
	}

	log.Info().Str("service", constants.Service).Msg("database connection is established")

	if err := db.AutoMigrate(&models.UserOrders{}); err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msg("failed to run automigrate")
	}
	log.Info().Msg("database migration complete")

	app := fiber.New()

	router.SetupRoutes(app, db)

	log.Info().Msgf("starting server on port %s", PORT)
	log.Fatal().Err(app.Listen(":" + PORT)).Msg("server exited")
}
