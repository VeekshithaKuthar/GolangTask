package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"paymenst/constants"
	"paymenst/database"
	"paymenst/messaging"
	"paymenst/models"
	"paymenst/repositories"
	"strings"
	"syscall"
	"time"

	"paymenst/router"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	DSN   string
	PORT  string
	debug bool
)

func main() {
	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
	flag.Parse()
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	DSN = os.Getenv("DSN")
	if DSN == "" {
		DSN = `host=pg user=app password=app123 dbname=paymentdb port=5432 sslmode=disable`
		//DSN = `host=localhost user=app password=app123 dbname=paymentdb port=5433 sslmode=disable`
		log.Info().Msg(DSN)
	}

	SEEDS := os.Getenv("KAFKA_BROKERS")
	if SEEDS == "" {
		SEEDS = "kafka1:9092,kafka2:9092,kafka3:9092"
		//SEEDS = "localhost:19092,localhost:29092,localhost:39092"
	}

	Port := os.Getenv("PORT")
	if Port == "" {
		PORT = "8083"
	}
	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal().Msg("unable to connect to the database..." + err.Error())
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msg("unable to connect to the database ")
		return
	}

	log.Info().Str("service", constants.Service).Msg("database connection is established")

	// if err := db.AutoMigrate(&models.Payments{}); err != nil {
	// 	log.Fatal().
	// 		Err(err).
	// 		Str("service", constants.Service).
	// 		Msg("failed to run automigrate")
	// }
	log.Info().Msg("database migration complete")
	paymentDb := repositories.NewPaymentDB(db)
	consumer, err := messaging.NewConsumer(strings.Split(SEEDS, ","), "order.created", "food-consumer", paymentDb)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start consumer")
	}
	go consumer.Run()

	app := fiber.New()
	router.SetupRoutes(app, db)

	// Run Fiber in goroutine
	// go func() {
	// 	if err := app.Listen(":" + PORT); err != nil {
	// 		log.Fatal().Err(err).Msg("fiber server stopped")
	// 	}
	// }()

	// Wait for termination signal (Ctrl+C / kill)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Info().Msg("Shutting down gracefully...")

	// Shutdown Fiber
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Error().Err(err).Msg("Error shutting down Fiber")
	}

	log.Info().Msgf("starting server on port %s", PORT)
	log.Fatal().Err(app.Listen(":" + PORT)).Msg("server exited")

}

func Init(db *gorm.DB) error {
	db.AutoMigrate(&models.Payments{})
	return nil
}
