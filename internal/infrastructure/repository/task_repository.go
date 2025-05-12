package repository

import (
	"errors"
	"time"
	"todo-service/internal/domain"
)

var tasks []domain.Task

type taskRepository struct {
}

func NewTaskRepository() domain.TaskRepository {
	return &taskRepository{}
}

func (r *taskRepository) GetAll() ([]domain.Task, error) {
	return tasks, nil
}

func (r *taskRepository) GetByID(id int64) (*domain.Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *taskRepository) Insert(task domain.Task) (domain.Task, error) {
	task.ID = time.Now().UnixNano()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	tasks = append(tasks, task)
	return task, nil
}

func (r *taskRepository) Update(task *domain.Task) error {
	for i, t := range tasks {
		if t.ID == task.ID {
			t.IsDone = task.IsDone
			t.Title = task.Title
			t.Description = task.Description
			t.UpdatedAt = time.Now()
			tasks[i] = t
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *taskRepository) Delete(id int64) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return nil
}
