package main

import (
	"github.com/JustForCodin/chatv2/config"
	"gorm.io/gorm"
)

func main() {
	config := config.GetConfig()

	dbClient, err := gorm.Open(config.MysqlDSN)
	if err != nil {
		panic(err)
	}

	messageRepo := NewMessageRepo(dbClient)

	message, err := messageRepo.GetMessages(1)

}
