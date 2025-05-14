package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-service/internal/config"
	"todo-service/internal/interface/router"
	"todo-service/pkg/logx"
)

func main() {
	r := gin.Default()

	config.LoadConfig()
	logx.Init()

	router.RouteTasks(r)

	err := r.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
	if err != nil {
		panic(err)
	}
}
