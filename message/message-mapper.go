package message

import (
	"time"

	"github.com/JustForCodin/chatv2/room"
)

func FromDto(dto MessageDTO) Message {
	return Message{ID: dto.ID, Text: dto.Text, RoomID: dto.RoomID, Room: room.Room{
		ID: dto.RoomID, CreatedBy: time.Now(), CreatedAt: time.Now(),
	}, CreatedBy: dto.CreatedBy, CreatedAt: dto.CreatedAt}
}

func ToDto(message Message) MessageDTO {
	return MessageDTO{ID: message.ID, Text: message.Text, RoomID: message.Room.ID,
		CreatedBy: message.CreatedBy, CreatedAt: message.CreatedAt,
	}
}
