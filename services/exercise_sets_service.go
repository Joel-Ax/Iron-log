package services

import (
	"errors"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/repositories"
)

type ExerciseSetsService interface {
	CreateExerciseSet(exerciseSet *models.ExerciseSet) error
	GetAllExerciseSets() ([]models.ExerciseSet, error)
	GetExerciseSetID(id string) (*models.ExerciseSet, error)
	DeleteExerciseSet(id string) error
}

type exerciseSetsService struct {
	repo repositories.ExerciseSetsRepository
}

func NewExerciseSetsService(repo repositories.ExerciseSetsRepository) ExerciseSetsService {
	return &exerciseSetsService{repo: repo}
}

func (s *exerciseSetsService) CreateExerciseSet(exerciseSet *models.ExerciseSet) error {
	if exerciseSet.WorkoutID == 0 {
		return errors.New("workout id cannot be empty")
	}
	if exerciseSet.ExerciseID == 0 {
		return errors.New("exercise id cannot be empty")
	}
	if exerciseSet.SetNumber == 0 {
		return errors.New("number of sets cannot be empty")
	}
	return s.repo.Create(exerciseSet)
}

func (s *exerciseSetsService) GetAllExerciseSets() ([]models.ExerciseSet, error) {
	return s.repo.FindAll()
}

func (s exerciseSetsService) GetExerciseSetID(id string) (*models.ExerciseSet, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	return s.repo.FindByID(id)
}

func (s exerciseSetsService) DeleteExerciseSet(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	return s.repo.Delete(id)
}
