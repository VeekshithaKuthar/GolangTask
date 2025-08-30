// package main

// import (
// 	"context"
// 	"flag"
// 	"math/rand"
// 	"ordersApi/database"
// 	"ordersApi/handlers"
// 	"ordersApi/models"
// 	"ordersApi/repositories"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"syscall"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/rs/zerolog"
// 	"github.com/rs/zerolog/log"
// 	"gorm.io/gorm"
// )

// var (
// 	Jobs        chan repositories.Job
// 	DSN         string
// 	PORT        string
// 	debug       bool
// 	wg          sync.WaitGroup
// 	JOBS_BUFFER = 10
// )

// func main() {
// 	service := "Orders"
// 	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
// 	flag.Parse()
// 	// Default level for this example is info, unless debug flag is present
// 	zerolog.SetGlobalLevel(zerolog.InfoLevel)
// 	if debug {
// 		zerolog.SetGlobalLevel(zerolog.DebugLevel)
// 	}
// 	DSN = os.Getenv("DSN")
// 	if DSN == "" {
// 		DSN = `host=localhost user=app password=app123 dbname=usersdb port=5433 sslmode=disable`
// 		log.Info().Msg(DSN)
// 	}

// 	PORT := os.Getenv("PORT")
// 	if PORT == "" {
// 		PORT = "8081"
// 	}

// 	db, err := database.GetConnection(DSN)
// 	if err != nil {
// 		//log.Fatal().Msg("unable to connect to the database..." + err.Error())
// 		log.Fatal().
// 			Err(err).
// 			Str("service", service).
// 			Msgf("unable to connect to the database %s", service)
// 	}
// 	log.Info().Str("service", service).Msg("database connection is established")
// 	Init(db)

// 	Jobs = make(chan repositories.Job, JOBS_BUFFER)
// 	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
// 	defer stop()
// 	StartWorkers(db, 3)

// 	userDb := repositories.NewUserDB(db)           //hired a caretaker
// 	userHandler := handlers.NewUserHandler(userDb) //appointed a manager and Here’s your caretaker (userDb).
// 	// You can ask them for help 4
// 	// whenever you need to serve guests.
// 	orderDb := repositories.NewOrderDB(db, Jobs)
// 	orderHandler := handlers.NewOrderHandler(orderDb)

// 	app := fiber.New()
// 	user_group := app.Group("/api/v1/users")
// 	user_group.Post("/", userHandler.CreateUser)
// 	user_group.Get("/:id", userHandler.GetUserBy)

// 	order_group := app.Group("/api/v1/orders")
// 	order_group.Post("/", orderHandler.CreatOrder)
// 	order_group.Get("/:id", orderHandler.GetOrderBy)
// 	order_group.Post("/:id/confirm", orderHandler.ConfirmOrderById)

// 	go func() {
// 		if err := app.Listen(":" + PORT); err != nil {
// 			log.Fatal().Err(err).Msg("fiber server stopped unexpectedly")
// 		}
// 	}()
// 	<-ctx.Done()
// 	log.Info().Msg("shutdown signal received")

// 	_ = app.Shutdown()
// 	close(Jobs)
// 	wg.Wait()
// 	log.Info().Msg("all workers finished, shutdown complete")

// }

// func Init(db *gorm.DB) {
// 	db.AutoMigrate(&models.User{}, &models.Order{})
// }

// func StartWorkers(db *gorm.DB, num int) {
// 	for i := 0; i < num; i++ {
// 		wg.Add(1)
// 		go func(workerID int) {
// 			defer wg.Done()
// 			for job := range Jobs {
// 				order := new(models.Order)
// 				if err := db.First(order, job.OrderId).Error; err != nil {
// 					log.Error().Err(err).Int("worker", workerID).Msg("order not found")
// 					continue
// 				}

// 				time.Sleep(2 * time.Second)
// 				status := "failed"
// 				if rand.Intn(2) == 0 {
// 					status = "confirmed"
// 				}

//					db.Model(order).Where("id = ?", job.OrderId).Update("status", status)
//					log.Info().
//						Int("worker", workerID).
//						Uint("order_id", job.OrderId).
//						Str("status", status).
//						Msg("job processed")
//				}
//				log.Info().Int("worker", workerID).Msg("worker exiting")
//			}(i)
//		}
//	}
package main

import (
	"context"
	"flag"
	"math/rand"
	"ordersApi/database"
	"ordersApi/handlers"
	"ordersApi/models"
	"ordersApi/repositories"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	Jobs        chan repositories.Job
	DSN         string
	PORT        string
	debug       bool
	wg          sync.WaitGroup
	JOBS_BUFFER = 10
)

func main() {
	service := "Orders"
	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
	flag.Parse()
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	DSN = os.Getenv("DSN")
	if DSN == "" {
		DSN = `host=localhost user=app password=app123 dbname=usersdb port=5433 sslmode=disable`
		log.Info().Msg(DSN)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}

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

	Jobs = make(chan repositories.Job, JOBS_BUFFER)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	StartWorkers(db, 3)

	userDb := repositories.NewUserDB(db)           //hired a caretaker
	userHandler := handlers.NewUserHandler(userDb) //appointed a manager and Here’s your caretaker (userDb).
	// You can ask them for help 4
	// whenever you need to serve guests.
	orderDb := repositories.NewOrderDB(db, Jobs)
	orderHandler := handlers.NewOrderHandler(orderDb)

	app := fiber.New()
	user_group := app.Group("/api/v1/users")
	user_group.Post("/", userHandler.CreateUser)
	user_group.Get("/:id", userHandler.GetUserBy)

	order_group := app.Group("/api/v1/orders")
	order_group.Post("/", orderHandler.CreatOrder)
	order_group.Get("/:id", orderHandler.GetOrderBy)
	order_group.Post("/:id/confirm", orderHandler.ConfirmOrderById)

	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			log.Fatal().Err(err).Msg("fiber server stopped unexpectedly")
		}
	}()
	<-ctx.Done()
	log.Info().Msg("shutdown signal received")

	_ = app.Shutdown()
	close(Jobs)
	wg.Wait()
	log.Info().Msg("all workers finished, shutdown complete")

}

func Init(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Order{})
}

func StartWorkers(db *gorm.DB, num int) {
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range Jobs {
				order := new(models.Order)
				if err := db.First(order, job.OrderId).Error; err != nil {
					log.Error().Err(err).Int("worker", workerID).Msg("order not found")
					continue
				}

				time.Sleep(2 * time.Second)
				status := "failed"
				if rand.Intn(2) == 0 {
					status = "confirmed"
				}

				db.Model(order).Where("id = ?", job.OrderId).Update("status", status)
				log.Info().
					Int("worker", workerID).
					Uint("order_id", job.OrderId).
					Str("status", status).
					Msg("job processed")
			}
			log.Info().Int("worker", workerID).Msg("worker exiting")
		}(i)
	}
}
