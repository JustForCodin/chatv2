package room

import "time"

type RoomDTO struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	Text      string
	CreatedBy time.Time
}
