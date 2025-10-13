package controllers

import (
	"net/http"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/services"
	"github.com/gofiber/fiber/v2"
)

type ExerciseSetsController struct {
	service services.ExerciseSetsService
}

func NewExerciseSetsController(service services.ExerciseSetsService) *ExerciseSetsController {
	return &ExerciseSetsController{service: service}
}

func (c *ExerciseSetsController) CreateExercise(ctx *fiber.Ctx) error {
	exerciseSet := new(models.ExerciseSet)

	if err := ctx.BodyParser(exerciseSet); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "request failed",
		})
	}

	if err := c.service.CreateExerciseSet(exerciseSet); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "exercise set has been created",
	})

}
func (c *ExerciseSetsController) GetExerciseSet(ctx *fiber.Ctx) error {
	ExerciseSet, err := c.service.GetAllExerciseSets()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get exercise",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Exercise sets fetched successfully",
		"data":    ExerciseSet,
	})
}

func (c *ExerciseSetsController) GetExerciseSetID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	exerciseSet, err := c.service.GetExerciseSetID(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"messsage": "could not get the exercise",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "exercise set fetched successfully",
		"data":    exerciseSet,
	})
}

func (c ExerciseSetsController) DeleteExerciseSet(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.service.DeleteExerciseSet(id); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "exercise set deleted successfully",
	})
}
