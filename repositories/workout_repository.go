package repositories

import (
	"github.com/Joel-Ax/go-fiber-postgres/models"
	"gorm.io/gorm"
)

type WorkoutRepository interface {
	Create(workout *models.Workout) error
	FindAll() ([]models.Workout, error)
	FindByID(id string) (*models.Workout, error)
	Delete(id string) error
}

type workoutRepository struct {
	db *gorm.DB
}

func NewWorkoutRespository(db *gorm.DB) WorkoutRepository {
	return &workoutRepository{db: db}
}

func (r *workoutRepository) Create(workout *models.Workout) error {
	return r.db.Create(workout).Error
}

func (r *workoutRepository) FindAll() ([]models.Workout, error) {
	var Workout []models.Workout
	err := r.db.Find(&Workout).Error
	return Workout, err
}

func (r *workoutRepository) FindByID(id string) (*models.Workout, error) {
	var Workout models.Workout
	err := r.db.Where("id = ?", id).First(&Workout).Error
	return &Workout, err
}

func (r *workoutRepository) Delete(id string) error {
	return r.db.Delete(&models.Workout{}, id).Error
}
