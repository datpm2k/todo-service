package domain

import "time"

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type TaskRepository interface {
	GetAll() ([]Task, error)
	GetByID(id int64) (*Task, error)
	Insert(task Task) (Task, error)
	Update(task *Task) error
	Delete(id int64) error
}
