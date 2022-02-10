package main

import (
	"fmt"

	"github.com/JustForCodin/chatv2/config"
	"github.com/JustForCodin/chatv2/message"
	"gorm.io/gorm"
)

func main() {
	config := config.GetConfig()

	dbClient, err := gorm.Open(config.MysqlDSN)
	if err != nil {
		panic(err)
	}

	messageRepo := message.NewMessageRepo(dbClient)

	message, err := messageRepo.GetMessages(1)
	fmt.Println(message)

}
