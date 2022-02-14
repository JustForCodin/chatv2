package message

import "time"

type MessageDTO struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	Text      string
	RoomID    int64
	CreatedBy time.Time
	CreatedAt time.Time
}
