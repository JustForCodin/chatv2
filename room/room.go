package room

import (
	"time"

	"github.com/JustForCodin/chatv2/user"
)

type Room struct {
	ID        int64
	Name      string
	CreatedBy user.User
	CreatedAt time.Time
}
