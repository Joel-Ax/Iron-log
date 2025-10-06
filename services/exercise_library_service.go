package services

import "github.com/Joel-Ax/go-fiber-postgres/models"

type ExerciseLibraryService interface {
	CreateExercise(exercise *models.ExerciseLibrary) error
	GetAllExercises() ([]models.User, error)
	GetExerciseID(id string) (*models.ExerciseLibrary, error)
	DeleteExercise(id string) error
}
