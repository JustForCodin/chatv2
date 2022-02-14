package user

import "time"

type UserDto struct {
	ID           int64 `gorm:"primary_key;auto_increment;not_null"`
	Name         string
	Email        string // # unique
	Status       int8
	RegisteredAt time.Time
}
