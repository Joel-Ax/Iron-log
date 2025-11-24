package models

import "gorm.io/gorm"

type User struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Username *string `json:"username"`
	Email    string  `gorm:"uniqueIndex;not null" json:"email"`
	Password string  `json:"password,omitempty"`

	GoogleID     *string `gorm:"uniqueIndex" json:"google_id,omitempty"`
	Provider     string  `gorm:"default:'local'" json:"provider"`
	ProfileImage *string `json:"profile_image,omitempty"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
