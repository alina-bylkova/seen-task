package main

import (
	"fmt"

	"github.com/alinabylkova/seen-task/config/env"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := env.NewConfig()
	if err != nil {
		fmt.Printf("error while loading config %v", err)
	}

	r := gin.Default()

	r.Run()
}
