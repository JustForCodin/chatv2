package room

type RoomRepo interface {
	GetRooms(userID int64) ([]Room, error)
	CreateRoom(userID int64, message Room) (*Room, error)
	UpdateRoom(userID int64, message Room) (*Room, error)
	DeleteRoom(userID, messageID int64) (*Room, error)
}
