package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-service/internal/config"
	"todo-service/internal/interface/router"
	"todo-service/pkg/logger"
)

func main() {
	r := gin.Default()

	config.LoadConfig()

	router.RouteTasks(r)

	ctx := context.Background()

	err := r.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
	if err != nil {
		panic(err)
	} else {
		logger.Info(ctx, fmt.Sprintf("Server listening on port: %d\n", config.AppConfig.Port))
	}
}
