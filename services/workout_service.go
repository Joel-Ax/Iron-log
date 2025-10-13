package services

import (
	"errors"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/repositories"
)

type WorkoutService interface {
	CreateWorkout(workout *models.Workout) error
	GetAllWorkouts() ([]models.Workout, error)
	GetWorkoutById(id string) (*models.Workout, error)
	DeleteWorkout(id string) error
}

type workoutService struct {
	repo repositories.WorkoutRepository
}

func NewWorkoutService(repo repositories.WorkoutRepository) WorkoutService {
	return &workoutService{repo: repo}
}

func (s workoutService) CreateWorkout(workout *models.Workout) error {
	if workout.UserID == 0 {
		return errors.New("user id cannot be empty")
	}
	if workout.Name == nil || *workout.Name == "" {
		return errors.New("workout name cannot be empty")
	}
	if workout.WorkoutDate == nil || workout.WorkoutDate.IsZero() {
		return errors.New("workout date cannot be empty")
	}
	return s.repo.Create(workout)
}

func (s *workoutService) GetAllWorkouts() ([]models.Workout, error) {
	return s.repo.FindAll()
}

func (s workoutService) GetWorkoutById(id string) (*models.Workout, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}
	return s.repo.FindByID(id)
}

func (s *workoutService) DeleteWorkout(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	return s.repo.Delete(id)
}
