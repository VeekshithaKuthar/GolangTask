package main

import (
	"flag"
	"os"
	"strings"
	"users-service/database"
	"users-service/handlers"
	"users-service/messaging"
	"users-service/models"

	fiberprometheus "github.com/ansrivas/fiberprometheus/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	DSN   string
	PORT  string
	debug bool
	//Seeds []string = []string{"kafka1:9092", "kafka2:9092", "kafka3:9092"}
	//SeedsD    []string = []string{"localhost:19092", "localhost:29092", "localhost:39092"}
	//KAFKASEED []string
	SEEDS string
)

func main() {
	service := "users-service"
	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
	flag.Parse()
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	DSN = os.Getenv("DSN")
	if DSN == "" {
		DSN = `host=pg user=app password=app123 dbname=usersdb port=5433 sslmode=disable`
		log.Info().Msg(DSN)
	}
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

	// SEEDS = os.Getenv("KAFKA_BROKERS")
	// if SEEDS == "" {
	// 	SEEDS = "localhost:19092, localhost:29092, localhost:39092"
	// }

	db, err := database.GetConnection(DSN)

	if err != nil {
		//log.Fatal().Msg("unable to connect to the database..." + err.Error())
		log.Fatal().
			Err(err).
			Str("service", service).
			Msgf("unable to connect to the database %s", service)
	}
	log.Info().Str("service", service).Msg("database connection is established")
	Init(db)

	msgUsersCreated := messaging.NewMessaging("omnenext.users.created", strings.Split(SEEDS, ","))
	go msgUsersCreated.ProduceRecords()

	app := fiber.New()

	prom := fiberprometheus.New(service)
	prom.RegisterAt(app, "/metrics") // exposes Prometheus metrics here
	app.Use(prom.Middleware)         // automatic request metrics

	app.Get("/", handlers.Root)
	app.Get("ping", handlers.Ping)
	app.Get("/health", handlers.Health)

	userHandler := handlers.NewUserHandler(database.NewUserDB(db))
	user_group := app.Group("/api/v1/users")
	user_group.Post("/", userHandler.CreateUser(msgUsersCreated))
	user_group.Get("/:id", userHandler.GetUserBy)
	user_group.Get("/all/:limit/:offset", userHandler.GetUsersByLimit)

	order_group := app.Group("/api/v1/users/orders")
	order_group.Post("/", userHandler.CreateOrder)

	app.Listen(":" + PORT)

}

func Init(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Order{})
}
