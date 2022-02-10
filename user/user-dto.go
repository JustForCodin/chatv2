package user

import "time"

type UserDto struct {
	ID int64
	Name string
	Email string // # unique
	Status int8
	RegisteredAt time.Time
}