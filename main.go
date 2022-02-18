package main

import (
	"fmt"
	"log"

	"github.com/JustForCodin/chatv2/config"
	"github.com/JustForCodin/chatv2/message"
	"github.com/JustForCodin/chatv2/room"
	"github.com/JustForCodin/chatv2/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := config.GetConfig()

	dbClient, err := gorm.Open(mysql.Open(config.MysqlDSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	dbClient.AutoMigrate(&user.User{}, &message.Message{}, &room.Room{},
		&user.RoomUser{}, &user.UserCredentials{})

	messageRepo := message.NewMessageRepo(dbClient)
	roomRepo := room.NewRoomRepository(dbClient)
	userRepo := user.NewUserRepository(dbClient)

	newMessage, err := messageRepo.CreateMessage(1, message.Message{Text: "Hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newMessage)

	msg, err := messageRepo.GetMessages(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
	messageRepo.UpdateMessage(1, message.Message{Text: "New Msg"})
	messageRepo.GetMessages(1)
	fmt.Println(msg)

	if newRoom, err := roomRepo.CreateRoom(1, room.Room{Name: "New Room"}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(newRoom)
	}
	roomRepo.UpdateRoom(1, room.Room{Name: "Updated Room"})

	if newUser, err := userRepo.CreateUser(1, user.User{Name: "John", Email: "email@mail.com", Status: 1}); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(newUser)
	}
	userRepo.GetUser(1, 1)
}
