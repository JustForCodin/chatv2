package room

import (
	"fmt"
	"time"

	"github.com/JustForCodin/chatv2/user"
	"gorm.io/gorm"
)

type RoomRepo interface {
	GetRooms(userID int64) ([]Room, error)
	CreateRoom(userID int64, room Room) (*Room, error)
	UpdateRoom(userID int64, room Room) (*Room, error)
	DeleteRoom(userID, roomID int64) (*Room, error)
}

type RoomRepoImpl struct {
	db  *gorm.DB
	err error
}

func (r *RoomRepoImpl) GetRooms(roomID int64) ([]Room, error) {
	var rooms []Room
	r.db.Where("id=?", roomID).Find(&rooms)
	return rooms, r.err
}

func (r *RoomRepoImpl) CreateRoom(userID int64, room Room) (*Room, error) {
	room.CreatedAt = time.Now()
	room.CreatedBy = user.User{ID: userID}
	r.db.Create(room)
	fmt.Println("New room created by user ", userID)
	return &room, r.err
}

func (r *RoomRepoImpl) UpdateRoom(userID int64, room Room) (*Room, error) {
	var roomUpdate Room
	r.db.Where("id = ?, name = ?", roomUpdate.ID, roomUpdate.Name).Find(&roomUpdate)
	room.Name = roomUpdate.Name
	r.db.Save(&roomUpdate)
	return &roomUpdate, r.err
}

func (r *RoomRepoImpl) DeleteRoom(userID, roomID int64) (*Room, error) {
	var room Room
	r.db.Where("id = ?", room.ID).Find(&room)
	r.db.Delete(&room)
	return &room, r.err
}

func NewRoomRepository(db *gorm.DB) RoomRepo {
	return &RoomRepoImpl{
		db: db,
	}
}
