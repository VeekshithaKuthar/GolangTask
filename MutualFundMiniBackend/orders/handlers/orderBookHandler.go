package handlers

import (
	"orders/business"

	"github.com/gofiber/fiber/v2"
)

type IOrderBookHandler interface {
	HandleGetOrders(c *fiber.Ctx) error
}

type OrderBookHandler struct {
	service business.IOrderBookService
}

func NewOrderBookHandler(service business.IOrderBookService) IOrderBookHandler {
	return &OrderBookHandler{service: service}
}

func (controller *OrderBookHandler) HandleGetOrders(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "user_id is required")
	}
	orders, err := controller.service.GetOrders(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(orders)
}
