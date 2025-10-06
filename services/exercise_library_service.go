package services

import (
	"errors"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/repositories"
)

type ExerciseLibraryService interface {
	CreateExercise(exercise *models.ExerciseLibrary) error
	GetAllExercises() ([]models.ExerciseLibrary, error)
	GetExerciseID(id string) (*models.ExerciseLibrary, error)
	DeleteExercise(id string) error
}

type exerciseLibraryService struct {
	repo repositories.ExerciseLibraryRepository
}

func NewExerciseLibraryService(repo repositories.ExerciseLibraryRepository) ExerciseLibraryService {
	return &exerciseLibraryService{repo: repo}
}

func (s *exerciseLibraryService) CreateExercise(exercise *models.ExerciseLibrary) error {
	//Validation logic to create if needed
	if exercise.Name == nil || *exercise.Name == "" {
		return errors.New("name cannot be empty")
	}
	if exercise.Category == nil || *exercise.Category == "" {
		return errors.New("category cannot be emtpy")
	}
	if exercise.PrimaryMuscleGroup == nil || *exercise.PrimaryMuscleGroup == "" {
		return errors.New("muscle group cannot be empty")
	}
	return s.repo.Create(exercise)
}

func (s *exerciseLibraryService) GetAllExercises() ([]models.ExerciseLibrary, error) {
	return s.repo.FindAll()
}

func (s exerciseLibraryService) GetExerciseID(id string) (*models.ExerciseLibrary, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	return s.repo.FindByID(id)
}

func (s exerciseLibraryService) DeleteExercise(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	return s.repo.Delete(id)
}
