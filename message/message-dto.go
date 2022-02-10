package message

import "time"

type MessageDTO struct {
	ID int64
	Text string
	RoomID int64
	CreatedBy time.Time
	CreatedAt time.Time
}