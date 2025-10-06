package routes

import (
	"github.com/Joel-Ax/go-fiber-postgres/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.UserController, exerciseLibraryController *controllers.ExerciseLibraryController) {
	api := app.Group("/api")

	api.Post("/users", userController.CreateUser)
	api.Get("/users/:id", userController.GetUserByID)
	api.Get("/users", userController.GetUser)
	api.Delete("/users/:id", userController.DeleteUser)

	api.Post("/workouts")
	api.Get("/workouts/:id")
	api.Get("/workouts")
	api.Delete("/workouts/:id")

	api.Post("/exercise-library")
	api.Get("/exercise-library/:id")
	api.Get("/exercise-library")
	api.Delete("/exercise-library/:id")

	api.Post("/exercise-sets")
	api.Get("/exercise-sets/:id")
	api.Get("/exercise-sets")
	api.Delete("/exercise-sets/:id")

}
