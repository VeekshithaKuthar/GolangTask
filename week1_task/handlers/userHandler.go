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

type UserHandler struct {
	repositories.IuserDB
}

type IUserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUserBy(c *fiber.Ctx) error
}

func NewUserHandler(iuserDb repositories.IuserDB) IUserHandler {
	return &UserHandler{iuserDb}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	err = user.Validate()
	if err != nil {
		return err
	}
	user.LastModified = time.Now().Unix()

	user, err = uh.Create(user)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (uh *UserHandler) GetUserBy(c *fiber.Ctx) error {
	id := c.Params("id") // Retrieves the value of ":id"

	_id, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid id 22")
	}

	user, err := uh.GetBy(uint(_id))
	if err != nil {
		log.Err(err).Msg("data might not be available or some sql issue")
		return errors.New("something went wrong or no data available with that id")
	}

	return c.JSON(user)
}
