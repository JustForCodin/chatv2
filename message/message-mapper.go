package message

import (
	"github.com/JustForCodin/chatv2/room"
)

func FromDto(dto MessageDTO) Message {
	return Message{ID: dto.ID, Text: dto.Text, RoomID: dto.RoomID, Room: room.Room{
		ID: dto.RoomID,
	}}
}

func ToDto(message Message) MessageDTO {
	return MessageDTO{ID: message.ID, Text: message.Text, RoomID: message.Room.ID}
}
