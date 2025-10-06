package repositories

import (
	"github.com/Joel-Ax/go-fiber-postgres/models"
	"gorm.io/gorm"
)

type ExerciseLibraryRepository interface {
	Create(exercise *models.ExerciseLibrary) error
	FindAll() ([]models.ExerciseLibrary, error)
	FindByID(id string) (*models.ExerciseLibrary, error)
	Delete(id string) error
}

type exerciseLibraryRepository struct {
	db *gorm.DB
}

func NewExerciseLibraryRepository(db *gorm.DB) ExerciseLibraryRepository {
	return &exerciseLibraryRepository{db: db}
}

func (r *exerciseLibraryRepository) Create(exercise *models.ExerciseLibrary) error {
	return r.db.Create(exercise).Error
}

func (r *exerciseLibraryRepository) FindAll() ([]models.ExerciseLibrary, error) {
	var exercise []models.ExerciseLibrary
	err := r.db.Find(&exercise).Error
	return exercise, err
}

func (r *exerciseLibraryRepository) FindByID(id string) (*models.ExerciseLibrary, error) {
	var exercise models.ExerciseLibrary
	err := r.db.Where("id = ?", id).First(&exercise).Error
	return &exercise, err
}

func (r *exerciseLibraryRepository) Delete(id string) error {
	return r.db.Delete(&models.ExerciseLibrary{}, id).Error
}
