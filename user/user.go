package user

import "time"

type User struct {
	ID           int64 `gorm:"primary_key;auto_increment;not_null"`
	Name         string
	Email        string
	Status       int8
	RegisteredAt time.Time
}
