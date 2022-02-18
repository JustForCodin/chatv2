package user

import "time"

type UserDto struct {
	ID           int64  `gorm:"primary_key;auto_increment;not_null"`
	Name         string `gorm:"unique;not_null"`
	Email        string `gorm:"unique;not_null"`
	Status       int8
	RegisteredAt time.Time
}
