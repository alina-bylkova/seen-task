package main

import (
	"log"

	"github.com/alinabylkova/seen-task/api"
	"github.com/alinabylkova/seen-task/config/env"
	"github.com/alinabylkova/seen-task/db"
	"github.com/alinabylkova/seen-task/model"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := env.NewConfig()
	if err != nil {
		log.Printf("error while loading config %v", err)
	}

	db, err := db.New(config)
	if err != nil {
		log.Fatal("Failed initializing db")
	}
	db.Gorm.AutoMigrate(&model.Recipient{}, &model.Video{}, &model.Event{})

	r := gin.Default()

	r.GET("recipients", api.GetRecipients(db))
	r.GET("recipients/:id", api.GetRecipientById(db))
	r.POST("recipients", api.PostRecipient(db))

	r.Run()
}
