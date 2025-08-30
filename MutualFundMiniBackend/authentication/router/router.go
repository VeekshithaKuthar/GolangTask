package router

import (
	"authentication/business"
	"authentication/constants"
	"authentication/handlers"
	"authentication/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	loginRepo := repositories.NewLoginRepository(db)
	loginService := business.NewLoginService(loginRepo)
	loginHandler := handlers.NewLoginHandler(loginService)

	app.Get(constants.HealthCheck, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	})
	app.Post(constants.UserLogin, loginHandler.HandleCreateUser)
}
