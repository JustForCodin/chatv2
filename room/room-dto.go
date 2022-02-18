package room

import (
	"time"

	"github.com/JustForCodin/chatv2/user"
)

type RoomDTO struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	Text      string
	CreatedBy user.UserDto `gorm:"references:users.ID"`
	CreatedAt time.Time
}
