package models

import (
	"time"

	"gorm.io/gorm"
)

type ExerciseSet struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	WorkoutID       uint      `gorm:"not null;index" json:"workout_id"`
	ExerciseID      uint      `gorm:"not null;index" json:"exercise_id"`
	SetNumber       int       `gorm:"not null" json:"set_number"`
	WeightKg        float64   `json:"weight_kg"`
	Reps            int       `json:"reps"`
	DurationSeconds int       `json:"duration_seconds"`
	Notes           string    `gorm:"type:text" json:"notes"`
	Completed       bool      `gorm:"default:false" json:"completed"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func MigrateExerciseSet(db *gorm.DB) error {
	err := db.AutoMigrate(&ExerciseSet{})
	return err
}
