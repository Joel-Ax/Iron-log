package models

import (
	"time"

	"gorm.io/gorm"
)

type ExerciseLibrary struct {
	ID                 uint      `gorm:"primaryKey:autoIncrement" json:"id"`
	Name               string    `gorm:"type:varchar(255);not null" json:"name"`
	Category           string    `gorm:"type:varchar(255);not null" json:"category"`
	PrimaryMuscleGroup string    `gorm:"type:varchar(255);not null" json:"primary_muscle_group"`
	CreatedAt          time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func MigrateExerciseLibrary(db *gorm.DB) error {
	err := db.AutoMigrate(&ExerciseLibrary{})
	return err
}
