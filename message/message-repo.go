package message

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MessageRepo interface {
	GetMessages(userID int64) ([]Message, error)
	CreateMessage(userID int64, message Message) (*Message, error)
	UpdateMessage(userID int64, message Message) (*Message, error)
	DeleteMessage(userID, messageID int64) (*Message, error)
}

type MessageRepoImpl struct {
	db  *gorm.DB
	err error
}

func (r *MessageRepoImpl) GetMessages(userID int64) ([]Message, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	user_exists := r.db.Raw("SELECT * FROM user WHERE user.id=%d LIMIT 1", userID)
	if user_exists == nil {
		panic("User " + string(userID) + "is not allowed to see messages")
	}

	var messages []Message
	r.db.Find(&messages)

	return messages, r.err
}

func (r *MessageRepoImpl) CreateMessage(userID int64, message Message) (*Message, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}
	r.db.Create(message)
	fmt.Printf("New message created by user %d", userID)
	return &message, r.err
}

func (r *MessageRepoImpl) UpdateMessage(userID int64, message Message) (*Message, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	var messageUpdate Message
	r.db.Where("text = ?", message.Text).Find(&messageUpdate)
	messageUpdate.Text = message.Text
	r.db.Save(&message)
	return &messageUpdate, r.err
}

func (r *MessageRepoImpl) DeleteMessage(userID, messageID int64) (*Message, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	user_exists := r.db.Raw("SELECT * FROM user WHERE user.id=%d LIMIT 1", userID)
	if user_exists == nil {
		panic("User " + string(userID) + "is not allowed to see messages")
	}

	var message Message
	r.db.Where("id = ?", message.ID).Find(&message)
	r.db.Delete(&message)
	return &message, r.err
}

func NewMessageRepo(db *gorm.DB) MessageRepo {
	return &MessageRepoImpl{
		db: db,
	}
}
