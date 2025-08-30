// package main

// import (
// 	"flag"
// 	"orders/constants"
// 	"orders/database"
// 	"orders/models"
// 	"orders/redis"
// 	"orders/router"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/log"
// )

// var (
// 	DSN   string
// 	PORT  string
// 	debug bool
// )

// func main() {
// 	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
// 	flag.Parse()
// 	// Default level for this example is info, unless debug flag is present
// 	zerolog.SetGlobalLevel(zerolog.InfoLevel)
// 	if debug {
// 		zerolog.SetGlobalLevel(zerolog.DebugLevel)
// 	}
// 	DSN = os.Getenv("DSN")
// 	if DSN == "" {
// 		//DSN = `host=pg user=app password=app123 dbname=paymentdb port=5432 sslmode=disable`
// 		DSN = `host=localhost user=app password=app123 dbname=orderdb port=5433 sslmode=disable`
// 		log.Info().Msg(DSN)
// 	}

// 	Port := os.Getenv("PORT")
// 	if Port == "" {
// 		PORT = "8082"
// 	}

// 	db, err := database.GetConnection(DSN)
// 	if err != nil {
// 		log.Fatal().Msg("unable to connect to the database..." + err.Error())
// 		log.Fatal().
// 			Err(err).
// 			Str("service", constants.Service).
// 			Msg("unable to connect to the database ")
// 		return
// 	}

// 	log.Info().Str("service", constants.Service).Msg("database connection is established")

// 	if err := db.AutoMigrate(&models.Order{}, &models.Holding{}); err != nil {
// 		log.Fatal().
// 			Err(err).
// 			Str("service", constants.Service).
// 			Msg("failed to run automigrate")

// 	}
// 	redis.InitRedisClient()
// 	log.Info().Msg("database migration complete")

// 	app := fiber.New()
// 	router.SetupRoutes(app, db)
// 	log.Info().Msgf("starting server on port %s", PORT)
// 	log.Fatal().Err(app.Listen(":" + PORT)).Msg("server exited")

// 	// app.Post("/admin/nav", func(c *fiber.Ctx) error {
// 	// 	type Req struct {
// 	// 		SchemeCode string  `json:"scheme_code"`
// 	// 		NAV        float64 `json:"nav"`
// 	// 	}
// 	// 	var req Req
// 	// 	if err := c.BodyParser(&req); err != nil {
// 	// 		return fiber.ErrBadRequest
// 	// 	}
// 	// 	err := redisdb.SetNAV(req.SchemeCode, req.NAV)
// 	// 	if err != nil {
// 	// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	// 	}
// 	// 	return c.SendString("NAV set")
// 	// })

// }
package main

import (
	"flag"
	"orders/constants"
	"orders/database"
	"orders/kafka"
	"orders/models"
	"orders/redis"
	"orders/router"
	workerbo "orders/worker_bo"
	"os"
	"strings"

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
		DSN = `host=localhost user=app password=app123 dbname=orderdb port=5433 sslmode=disable`
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	SEEDS = os.Getenv("KAFKA_BROKERS")
	if SEEDS == "" {
		SEEDS = "localhost:29092"
	}
	//brokers := []string{"localhost:29092"}

	brokers := strings.Split(SEEDS, ",")
	topic := "orders.placed"

	orderProducer := kafka.NewOrderProducer(brokers, topic)
	log.Info().Strs("brokers", brokers).Str("topic", topic).Msg("Starting BO consumer")

	go workerbo.StartBOConsumer(brokers, topic)

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msgf("unable to connect to the database %s", constants.Service)
	}
	log.Info().Str("service", constants.Service).Msg("database connection is established")

	if err := db.AutoMigrate(&models.Order{}, &models.Holding{}); err != nil {
		log.Fatal().
			Err(err).
			Str("service", constants.Service).
			Msg("failed to run automigrate")
	}
	log.Info().Msg("database migration complete")

	redis.InitRedisClient()

	app := fiber.New()
	router.SetupRoutes(app, db, orderProducer)

	log.Info().Msgf("starting server on port %s", PORT)
	log.Fatal().Err(app.Listen(":" + PORT)).Msg("server exited")
}
