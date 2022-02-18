package room

func FromDto(dto RoomDTO) Room {
	return Room{ID: dto.ID, Name: dto.Text, CreatedBy: dto.CreatedBy, CreatedAt: dto.CreatedAt}
}

func ToDto(room Room) RoomDTO {
	return RoomDTO{ID: room.ID, Text: room.Name, CreatedBy: room.CreatedBy, CreatedAt: room.CreatedAt}
}
