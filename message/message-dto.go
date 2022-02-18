package message

import (
	"time"

	"github.com/JustForCodin/chatv2/user"
)

type MessageDTO struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	Text      string
	RoomID    int64 `gorm:"primary_key;auto_increment;not_null"`
	CreatedBy user.UserDto
	CreatedAt time.Time
}
