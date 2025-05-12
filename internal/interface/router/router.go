package router

import (
	"github.com/gin-gonic/gin"
	"todo-service/internal/infrastructure/repository"
	"todo-service/internal/interface/handler"
	usecase "todo-service/internal/usecase/task"
)

func RouteTasks(r *gin.Engine) {
	taskRepo := repository.NewTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	r.POST("/api/v1/tasks", taskHandler.Insert)
	r.PUT("/api/v1/tasks/:id", taskHandler.Update)
	r.GET("/api/v1/tasks", taskHandler.GetAll)
	r.DELETE("/api/v1/tasks/:id", taskHandler.DeleteById)
}
