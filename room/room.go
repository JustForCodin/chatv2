package room

import (
	"time"

	"github.com/JustForCodin/chatv2/user"
)

type Room struct {
	ID        int64 `gorm:"primary_key;auto_increment;not_null"`
	Name      string
	CreatedBy user.User
	CreatedAt time.Time
}
