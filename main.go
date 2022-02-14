package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JustForCodin/chatv2/config"
	"github.com/JustForCodin/chatv2/message"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config := config.GetConfig()

	dbClient, err := gorm.Open(mysql.Open(config.MysqlDSN))
	if err != nil {
		panic(err)
	}

	messageRepo := message.NewMessageRepo(dbClient)

	newMessage, err := messageRepo.CreateMessage(1, message.Message{ID: 1, Text: "Hello", RoomID: 1,
		CreatedBy: time.Now(), CreatedAt: time.Now()})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newMessage)

	msg, err := messageRepo.GetMessages(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
	messageRepo.UpdateMessage(1, message.Message{Text: "New Msg", CreatedBy: time.Now(), CreatedAt: time.Now()})
	messageRepo.GetMessages(1)
	fmt.Println(msg)

}
