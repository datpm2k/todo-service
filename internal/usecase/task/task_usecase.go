package usecase

import (
	"todo-service/internal/domain"
)

type TaskUseCase interface {
	GetAll() ([]domain.Task, error)
	Insert(domain.Task) (domain.Task, error)
	Update(int64, *domain.Task) (*domain.Task, error)
	Delete(id int64) error
}

type taskUseCase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) TaskUseCase {
	return &taskUseCase{taskRepo: taskRepo}
}

func (t *taskUseCase) GetAll() ([]domain.Task, error) {
	return t.taskRepo.GetAll()
}

func (t *taskUseCase) Insert(task domain.Task) (domain.Task, error) {
	return t.taskRepo.Insert(task)
}

func (t *taskUseCase) Update(id int64, task *domain.Task) (*domain.Task, error) {
	oldTask, err := t.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	oldTask.Title = task.Title
	oldTask.Description = task.Description
	oldTask.IsDone = task.IsDone
	err = t.taskRepo.Update(oldTask)
	if err != nil {
		return nil, err
	}
	return oldTask, nil
}

func (t *taskUseCase) Delete(id int64) error {
	return t.taskRepo.Delete(id)
}
