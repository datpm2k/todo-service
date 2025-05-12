package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-service/internal/domain"
	usecase "todo-service/internal/usecase/task"
)

type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks, err := h.taskUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    tasks,
	})
}

func (h *TaskHandler) Insert(c *gin.Context) {
	request := &domain.Task{}
	err := c.BindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	task, err := h.taskUseCase.Insert(*request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    task,
	})
}

func (h *TaskHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	task := &domain.Task{}
	err = c.BindJSON(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	task, err = h.taskUseCase.Update(int64(id), task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    task,
	})
}

func (h *TaskHandler) DeleteById(c *gin.Context) {
	idQuery := c.Param("id")

	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.taskUseCase.Delete(int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
