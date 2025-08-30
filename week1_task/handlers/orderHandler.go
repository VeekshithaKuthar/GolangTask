package handlers

import (
	"errors"
	"ordersApi/models"
	"ordersApi/repositories"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type OrderHandler struct {
	repositories.IOrderDB
}
type IOrderHandler interface {
	CreatOrder(c *fiber.Ctx) error
	GetOrderBy(c *fiber.Ctx) error
	ConfirmOrderById(c *fiber.Ctx) error
}

func NewOrderHandler(iorderDb repositories.IOrderDB) IOrderHandler {
	return &OrderHandler{iorderDb}
}

func (uh *OrderHandler) CreatOrder(c *fiber.Ctx) error {
	order := new(models.Order)
	err := c.BodyParser(order)
	if err != nil {
		return err
	}

	err = order.Validate()
	if err != nil {
		return err
	}
	order.LastModified = time.Now().Unix()
	order.Status = "Pending"

	order, err = uh.CreateOrder(order)
	if err != nil {
		return err
	}
	return c.JSON(order)
}

func (uh *OrderHandler) GetOrderBy(c *fiber.Ctx) error {
	id := c.Params("id") // Retrieves the value of ":id"

	_id, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid id 11")
	}

	user, err := uh.GetOrderById(uint(_id))
	if err != nil {
		log.Err(err).Msg("data might not be available or some sql issue")
		return errors.New("something went wrong or no data available with that id")
	}

	return c.JSON(user)
}

func (uh *OrderHandler) ConfirmOrderById(c *fiber.Ctx) error {
	id := c.Params("id") // Retrieves the value of ":id"

	_id, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid id")
	}

	err = uh.ConfirmOrder(uint(_id))
	if err != nil {
		log.Err(err).Msg("data might not be available or some sql issue")
		return errors.New("something went wrong or no data available with that id")
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "order queued for confirmation",
		"orderId": _id,
	})

}
