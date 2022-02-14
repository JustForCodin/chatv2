package room

import (
	"log"

	"gorm.io/driver/mysql"
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

func (r *RoomRepoImpl) GetRooms(userID int64) ([]Room, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	var rooms []Room
	r.db.Where("id=?", userID).Find(&rooms)
	return rooms, r.err
}

func (r *RoomRepoImpl) CreateRoom(userID int64, room Room) (*Room, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}
	r.db.Create(room)
	return &room, r.err
}

func (r *RoomRepoImpl) UpdateRoom(userID int64, room Room) (*Room, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	var roomUpdate Room
	r.db.Where("id = ?, name = ?", roomUpdate.ID, roomUpdate.Name).Find(&roomUpdate)
	room.Name = roomUpdate.Name
	r.db.Save(&roomUpdate)
	return &roomUpdate, r.err
}

func (r *RoomRepoImpl) DeleteRoom(userID, roomID int64) (*Room, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

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
