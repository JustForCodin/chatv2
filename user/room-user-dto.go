package user

type RoomUserDTO struct {
	RoomID int64 `gorm:"primary_key;unique;not_null"`
	UserID int64 `gorm:"primary_key;unique;not_null"`
}
