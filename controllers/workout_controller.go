package controllers

import (
	"net/http"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/services"
	"github.com/gofiber/fiber/v2"
)

type WorkoutController struct {
	service services.WorkoutService
}

func NewWorkoutController(service services.WorkoutService) *WorkoutController {
	return &WorkoutController{service: service}
}

func (c *WorkoutController) CreateWorkout(ctx *fiber.Ctx) error {
	workout := new(models.Workout)

	if err := ctx.BodyParser(workout); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "request failed",
		})
	}

	if err := c.service.CreateWorkout(workout); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "workout has been created",
	})
}

func (c *WorkoutController) GetWorkout(ctx *fiber.Ctx) error {
	Workout, err := c.service.GetAllWorkouts()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get workouts",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Workout fetched successfully",
		"data":    Workout,
	})
}

func (c *WorkoutController) GetWorkoutById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	workout, err := c.service.GetWorkoutById(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get the workout",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "workout fetched successfully",
		"data":    workout,
	})
}

func (c *WorkoutController) DeleteWorkout(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.service.DeleteWorkout(id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "workout deleted successfully",
	})
}
