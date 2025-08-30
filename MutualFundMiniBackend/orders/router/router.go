// package router

// import (
// 	"orders/business"
// 	"orders/constants"
// 	"orders/handlers"
// 	"orders/repositories"

// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/gorm"
// )

// func SetupRoutes(app *fiber.App, db *gorm.DB) {

// 	app.Get(constants.HealthCheck, func(c *fiber.Ctx) error {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
// 	})

// 	orderDb := repositories.NewPlaceOrderRepository(db)
// 	orderservice := business.NewPlaceOrderService(orderDb)
// 	orderDbHandler := handlers.NewPlaceOrderHandler(orderservice)

// 	orderBookRepo := repositories.NewOrderBookRepository(db)
// 	orderBookService := business.NewOrderBookService(orderBookRepo)
// 	orderBookHandler := handlers.NewOrderBookHandler(orderBookService)

// 	app.Post(constants.PlaceOrder, orderDbHandler.HandlePlaceOrder)
// 	app.Get(constants.OrderBook, orderBookHandler.HandleGetOrders)

// }
package router

import (
	"orders/business"
	"orders/constants"
	"orders/handlers"
	"orders/kafka"
	"orders/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, orderProducer *kafka.OrderProducer) {
	ordersRepo := repositories.NewPlaceOrderRepository(db)
	ordersService := business.NewPlaceOrderService(ordersRepo, orderProducer)
	ordersHandler := handlers.NewPlaceOrderHandler(ordersService)

	orderBookRepo := repositories.NewOrderBookRepository(db)
	orderBookService := business.NewOrderBookService(orderBookRepo)
	orderBookHandler := handlers.NewOrderBookHandler(orderBookService)

	app.Get(constants.HealthCheck, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	})
	app.Post(constants.PlaceOrder, ordersHandler.HandlePlaceOrder)
	app.Get(constants.OrderBook, orderBookHandler.HandleGetOrders)

	// app.Post("/admin/nav", func(c *fiber.Ctx) error {
	//  type Req struct {
	//   SchemeCode string  `json:"scheme_code"`
	//   NAV        float64 `json:"nav"`
	//  }
	//  var req Req
	//  if err := c.BodyParser(&req); err != nil {
	//   return fiber.ErrBadRequest
	//  }
	//  err := redisdb.SetNAV(req.SchemeCode, req.NAV)
	//  if err != nil {
	//   return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	//  }
	//  return c.SendString("NAV set")
	// })

}
