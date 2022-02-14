package user

type RoomUserDTO struct {
	RoomID int64
	UserID int64 `gorm:"primary_key;auto_increment;not_null"`
}
