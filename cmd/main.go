package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-service/internal/config"
	"todo-service/internal/interface/router"
	"todo-service/pkg/logx"
)

func main() {
	config.LoadConfig()

	r := gin.Default()
	router.RouteTasks(r)

	logx.Info(context.Background(), "Started application")

	err := r.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
	if err != nil {
		panic(err)
	}
}
