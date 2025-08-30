package router

import (
	"paymenst/business"
	"paymenst/constants"
	"paymenst/handlers"
	"paymenst/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	paymentDb := repositories.NewPaymentDB(db)
	service := business.NewPaymentService(paymentDb)
	paymentDbHandler := handlers.NewPaymentHandler(service)

	app.Get(constants.HealthCheck, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	})
	app.Post(constants.CreatePaymenst, paymentDbHandler.HandleCreatPayment)
}
