package message

import (
	"time"

	"github.com/JustForCodin/chatv2/room"
)

type Message struct {
	ID        int64
	Text      string
	RoomID    int64
	Room      room.Room
	CreatedBy time.Time
	CreatedAt time.Time
}
