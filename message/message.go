package message

import (
	"time"

	"github.com/JustForCodin/chatv2/room"
	"github.com/JustForCodin/chatv2/user"
)

type Message struct {
	ID        int64
	Text      string
	RoomID    int64
	Room      room.Room
	CreatedBy user.UserDto
	CreatedAt time.Time
}
