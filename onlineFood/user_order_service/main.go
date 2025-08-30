package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"userorders/constants"
	"userorders/database"
	"userorders/messaging"
	"userorders/models"
	"userorders/repository"
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
		DSN = `host=pg user=app password=app123 dbname=usersdb port=5432 sslmode=disable`
		//DSN = "host=localhost user=app password=app123 dbname=usersdb port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	SEEDS := os.Getenv("KAFKA_BROKERS")
	if SEEDS == "" {
		SEEDS = "kafka1:9092,kafka2:9092,kafka3:9092" // default inside docker-compose network
		//SEEDS = "localhost:19092,localhost:29092,localhost:39092"
	}

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msgf("unable to connect to the database %s", constants.Service)
	}

	if err := db.AutoMigrate(&models.UserOrders{}); err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msg("failed to run automigrate")
	}
	log.Info().Msg("database migration complete")
	msgUsersCreated := messaging.NewMessaging("order.created", strings.Split(SEEDS, ","))
	go msgUsersCreated.Producerecord()
	log.Info().Str("service", constants.Service).Msg("database connection is established")

	orderDb := repository.NewOrderRepository(db)
	consumer, err := messaging.NewPaymentConsumer(strings.Split(SEEDS, ","), "payment_status_updated", "order-consumer", orderDb)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start consumer")
	}
	go consumer.Run()
	app := fiber.New()

	router.SetupRoutes(app, db, msgUsersCreated)

	// Run Fiber in goroutine
	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			log.Fatal().Err(err).Msg("fiber server stopped")
		}
	}()

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

	// log.Info().Msgf("starting server on port %s", PORT)
	// log.Fatal().Err(app.Listen(":" + PORT)).Msg("server exited")
}
