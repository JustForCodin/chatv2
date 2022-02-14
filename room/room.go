package room

import "time"

type Room struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	Name      string
	CreatedBy time.Time
	CreatedAt time.Time
}
