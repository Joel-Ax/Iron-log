package models

import (
	"time"

	"gorm.io/gorm"
)

type Workout struct {
	ID          uint       `gorm:"primary key;autoIncrement" json:"id"`
	UserID      uint       `gorm:"not null;index;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_id"`
	Name        *string    `gorm:"type:varchar(255);not null" json:"name"`
	WorkoutDate *time.Time `gorm:"type:date;not null" json:"workout_date"`
	StartTime   time.Time  `gorm:"type:date;time" json:"start_time"`
	EndTime     time.Time  `gorm:"type:date;time" json:"end_time"`
	Notes       string     `gorm:"type:text" json:"notes"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoCreateTime" json:"updated_at"`
}

func MigrateWorkout(db *gorm.DB) error {
	err := db.AutoMigrate(&Workout{})
	return err
}
