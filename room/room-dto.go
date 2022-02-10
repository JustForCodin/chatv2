package room

import "time"

type RoomDTO struct {
	ID int64
	Text string
	CreatedBy time.Time
}