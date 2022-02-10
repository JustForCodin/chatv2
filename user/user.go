package user

import "time"

type User struct {
	ID int64
	Name string
	Email string
	Status int8
	RegisteredAt time.Time
}