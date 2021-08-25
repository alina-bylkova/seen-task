package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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

	router := gin.Default()

	authorized := router.Group("/api", gin.BasicAuth(gin.Accounts{
		config.AuthUser: config.AuthPassword,
	}))

	authorized.GET("recipients", api.GetRecipients(db))
	authorized.GET("recipients/:id", api.GetRecipientById(db))
	authorized.POST("recipients", api.PostRecipient(db))
	authorized.POST("events", api.PostEvent(db))

	server := &http.Server{
		Addr:    config.ServerAddress,
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Serer is listening on port: %s", config.ServerAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("Shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
