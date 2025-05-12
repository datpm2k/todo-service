package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo-service/internal/interface/router"
)

func main() {
	r := gin.Default()

	router.RouteTasks(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	} else {
		log.Println("Server listening on port: 8080")
	}
}
