package router

import (
	"userorders/business"
	"userorders/constants"
	"userorders/handlers"
	"userorders/messaging"
	"userorders/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, msgUsersCreated *messaging.Messaging) {
	service := business.NewOrderService(repository.NewOrderRepository(db))
	handler := handlers.NewOrderHandler(service)

	app.Get(constants.HealthCheck, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	})
	app.Post(constants.CreateOrders, handler.HandleCreateOrder(msgUsersCreated))
}
