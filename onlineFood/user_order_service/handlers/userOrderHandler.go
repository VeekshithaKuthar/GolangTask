package handlers

import (
	"time"
	"userorders/business"
	"userorders/constants"
	"userorders/messaging"
	"userorders/models"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	business.IOrderService
}

type IOrderHandler interface {
	HandleCreateOrder(msg *messaging.Messaging) func(c *fiber.Ctx) error
}

func NewOrderHandler(iorderServicedb business.IOrderService) IOrderHandler {
	return &OrderHandler{iorderServicedb}
}

func (controller *OrderHandler) HandleCreateOrder(msg *messaging.Messaging) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var userOrderRequest models.UserOrders
		if err := c.BodyParser(&userOrderRequest); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, constants.BadRequest)
		}
		if err := userOrderRequest.Validate(); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		userOrderRequest.CreatedAt = time.Now()

		order, err := controller.CreateOrder(&userOrderRequest)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		msg.ChMessaging <- order.ToBytes()
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"order_id": order.OrderID,
		})
	}
}
