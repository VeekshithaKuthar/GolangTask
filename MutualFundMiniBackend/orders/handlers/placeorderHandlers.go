package handlers

import (
	"orders/business"
	"orders/constants"
	"orders/models"

	"github.com/gofiber/fiber/v2"
)

type PlaceOrderHandler struct {
	Handler business.IPlaceOrderService
}

type IPlaceOrderHandler interface {
	HandlePlaceOrder(c *fiber.Ctx) error
}

func NewPlaceOrderHandler(controller business.IPlaceOrderService) IPlaceOrderHandler {
	return &PlaceOrderHandler{Handler: controller}
}

func (controller *PlaceOrderHandler) HandlePlaceOrder(c *fiber.Ctx) error {
	var orderRequest models.PlaceOrderRequest
	if err := c.BodyParser(&orderRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.BadRequest)
	}

	if err := orderRequest.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	order, err := controller.Handler.PlaceOrder(&orderRequest)
	if err != nil {
		return err
	}
	return c.JSON(order)

}
