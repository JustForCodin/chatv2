package room

import "time"

type Room struct {
	ID int64
	Name string
	CreatedBy time.Time
	CreatedAt time.Time
}