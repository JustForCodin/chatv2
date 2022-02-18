package user

type RoomUserDTO struct {
	RoomID int64 `gorm:"references:rooms.ID"`
	UserID int64 `gorm:"references:users.ID"`
}
