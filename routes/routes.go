package routes

import (
	"github.com/Joel-Ax/go-fiber-postgres/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userController *controllers.UserController, exerciseLibraryController *controllers.ExerciseLibraryController, exerciseSetsController *controllers.ExerciseSetsController, workoutController *controllers.WorkoutController) {
	api := app.Group("/api")

	api.Post("/users", userController.CreateUser)
	api.Get("/users/:id", userController.GetUserByID)
	api.Get("/users", userController.GetUser)
	api.Delete("/users/:id", userController.DeleteUser)

	api.Post("/workouts", workoutController.CreateWorkout)
	api.Get("/workouts/:id", workoutController.GetWorkoutById)
	api.Get("/workouts", workoutController.GetWorkout)
	api.Delete("/workouts/:id", workoutController.DeleteWorkout)

	api.Post("/exercise-library", exerciseLibraryController.CreateExercise)
	api.Get("/exercise-library/:id", exerciseLibraryController.GetExerciseByID)
	api.Get("/exercise-library", exerciseLibraryController.GetExercise)
	api.Delete("/exercise-library/:id", exerciseLibraryController.DeleteExercise)

	api.Post("/exercise-sets", exerciseSetsController.CreateExercise)
	api.Get("/exercise-sets/:id", exerciseSetsController.GetExerciseSetByID)
	api.Get("/exercise-sets", exerciseSetsController.GetExerciseSet)
	api.Delete("/exercise-sets/:id", exerciseSetsController.DeleteExerciseSet)

}
