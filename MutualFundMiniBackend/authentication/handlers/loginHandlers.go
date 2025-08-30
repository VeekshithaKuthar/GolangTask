package handlers

import (
	"authentication/business"
	"authentication/constants"
	"authentication/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ILoginHandler interface {
	HandleCreateUser(c *fiber.Ctx) error
}

type LoginHandler struct {
	service business.ILoginService
}

func NewLoginHandler(service business.ILoginService) ILoginHandler {
	return &LoginHandler{service: service}
}

func (controller *LoginHandler) HandleCreateUser(c *fiber.Ctx) error {
	var userRequest models.LoginRequestModel
	if err := c.BodyParser(&userRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.BadRequest)
	}
	if err := userRequest.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	userRequest.CreatedAt = time.Now()

	err := controller.service.CreateUser(&userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": constants.SuccessMessage,
	})

}
