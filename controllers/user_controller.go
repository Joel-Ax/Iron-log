package controllers

import (
	"net/http"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "request failed",
		})
	}

	if err := c.service.CreateUser(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user has been created",
	})
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	User, err := c.service.GetAllUsers()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get user",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User fetched successfully",
		"data":    User,
	})
}

func (c *UserController) GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := c.service.GetUserByID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get the user",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user fetched successfully",
		"data":    user,
	})
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.service.DeleteUser(id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}
