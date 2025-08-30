package handlers

import (
	"paymenst/business"
	"paymenst/constants"
	"paymenst/models"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	business.IPaymentService
}

type IPaymentHandler interface {
	HandleCreatPayment(c *fiber.Ctx) error
}

func NewPaymentHandler(ipaymenService business.IPaymentService) IPaymentHandler {
	return &PaymentHandler{ipaymenService}
}

func (controller *PaymentHandler) HandleCreatPayment(c *fiber.Ctx) error {
	var paymentRequest models.Payments
	//payment := new(models.Payments)
	if err := c.BodyParser(&paymentRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, constants.BadRequest)
	}

	if err := paymentRequest.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	payment, err := controller.CreatePayment(&paymentRequest)
	if err != nil {
		return err
	}
	// msg.ChMessaging <- payment.ToBytes()
	// fmt.Println(msg.ChMessaging)
	return c.JSON(payment)
}
