package repositories

import (
	"github.com/Joel-Ax/go-fiber-postgres/models"
	"gorm.io/gorm"
)

type ExerciseSetsRepository interface {
	Create(exerciseSet *models.ExerciseSet) error
	FindAll() ([]models.ExerciseSet, error)
	FindByID(id string) (*models.ExerciseSet, error)
	Delete(id string) error
}

type exerciseSetsRepository struct {
	db *gorm.DB
}

func NewExerciseSetsRepository(db *gorm.DB) ExerciseSetsRepository {
	return &exerciseSetsRepository{db: db}
}

func (r *exerciseSetsRepository) Create(exerciseSet *models.ExerciseSet) error {
	return r.db.Create(exerciseSet).Error
}

func (r *exerciseSetsRepository) FindAll() ([]models.ExerciseSet, error) {
	var exerciseSet []models.ExerciseSet
	err := r.db.Find(&exerciseSet).Error
	return exerciseSet, err
}

func (r *exerciseSetsRepository) FindByID(id string) (*models.ExerciseSet, error) {
	var exerciseSet models.ExerciseSet
	err := r.db.Where("id = ?", id).First(&exerciseSet).Error
	return &exerciseSet, err
}

func (r *exerciseSetsRepository) Delete(id string) error {
	return r.db.Delete(&models.ExerciseSet{}, id).Error
}
