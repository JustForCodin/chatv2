package user

type RoomUser struct {
	RoomID int64 `gorm:"primary_key;unique;not_null"`
	UserID int64 `gorm:"primary_key;unique;not_null"`
}
