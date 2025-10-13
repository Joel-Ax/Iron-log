package controllers

import (
	"net/http"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/services"
	"github.com/gofiber/fiber/v2"
)

type ExerciseLibraryController struct {
	service services.ExerciseLibraryService
}

func NewExerciseLibraryController(service services.ExerciseLibraryService) *ExerciseLibraryController {
	return &ExerciseLibraryController{service: service}
}

func (c *ExerciseLibraryController) CreateExercise(ctx *fiber.Ctx) error {
	exercise := new(models.ExerciseLibrary)

	if err := ctx.BodyParser(exercise); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "request failed",
		})
	}

	if err := c.service.CreateExercise(exercise); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "exercise has been created",
	})
}

func (c *ExerciseLibraryController) GetExercise(ctx *fiber.Ctx) error {
	Exercise, err := c.service.GetAllExercises()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get exercise",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Exercises fetched successfully",
		"data":    Exercise,
	})
}

func (c *ExerciseLibraryController) GetExerciseID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	exercise, err := c.service.GetExerciseID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get the exercise",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "exercise fetched successfully",
		"data":    exercise,
	})
}

func (c ExerciseLibraryController) DeleteExercise(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.service.DeleteExercise(id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "exercise deleted succesfully",
	})
}
