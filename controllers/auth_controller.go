package controllers

import (
	"net/http"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	if err := c.service.Register(user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user registered successfully",
	})
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&loginReq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	token, err := c.service.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login successful",
		"token":   token,
	})
}
