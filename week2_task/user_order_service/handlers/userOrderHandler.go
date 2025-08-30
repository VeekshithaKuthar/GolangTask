package handlers

import (
	"time"
	"userorders/business"
	"userorders/constants"
	"userorders/models"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	business.IOrderService
}

type IOrderHandler interface {
	HandleCreateOrder(c *fiber.Ctx) error
}

func NewOrderHandler(iorderServicedb business.IOrderService) IOrderHandler {
	return &OrderHandler{iorderServicedb}
}

func (controller *OrderHandler) HandleCreateOrder(c *fiber.Ctx) error {
	var userOrderRequest models.UserOrders
	if err := c.BodyParser(&userOrderRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.BadRequest)
	}
	if err := userOrderRequest.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	userOrderRequest.CreatedAt = time.Now()

	orderID, err := controller.CreateOrder(&userOrderRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"order_id": orderID,
	})
}
