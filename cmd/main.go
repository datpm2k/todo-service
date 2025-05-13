package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"todo-service/internal/config"
	"todo-service/internal/interface/router"
)

func main() {
	r := gin.Default()

	config.LoadConfig()

	router.RouteTasks(r)

	err := r.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
	if err != nil {
		panic(err)
	} else {
		log.Printf("Server listening on port: %d\n", config.AppConfig.Port)
	}
}
