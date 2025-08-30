package main

import (
	"authentication/constants"
	"authentication/database"
	"authentication/models"
	"authentication/router"
	"flag"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	DSN   string
	PORT  string
	debug bool
	SEEDS string
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
		// DSN = "host=pg user=app password=app123 dbname=users port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		DSN = `host=localhost user=app password=app123 dbname=orderdb port=5433 sslmode=disable`
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	SEEDS = os.Getenv("KAFKA_BROKERS")
	if SEEDS == "" {
		SEEDS = "kafka1:9092,kafka2:9092,kafka3:9092"
	}

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msgf("unable to connect to the database %s", constants.Service)
	}
	log.Info().Str("service", constants.Service).Msg("database connection is established")

	if err := db.AutoMigrate(&models.LoginRequestModel{}); err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msg("failed to run automigrate")
	}
	log.Info().Msg("database migration complete")

	// msgOrdersCreated := kafka.NewMessaging("userorders.created", strings.Split(SEEDS, ","))
	// go msgOrdersCreated.ProduceRecords()

	// paymentsProcessedConsumer := kafka.NewConsumer("payments.processed", strings.Split(SEEDS, ","))

	app := fiber.New()

	// router.SetupRoutes(app, db, msgOrdersCreated.ChMessaging, paymentsProcessedConsumer)
	router.SetupRoutes(app, db)

	log.Info().Msgf("starting server on port %s", PORT)
	log.Fatal().Err(app.Listen(":" + PORT)).Msg("server exited")
}
