package domain

import (
	"context"
	"time"
)

type Task struct {
	ID        ID        `json:"id"`
	Title     string    `form:"title" binding:"required" json:"title"`
	UserID    ID        `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID *ID) ([]*Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID *ID) ([]*Task, error)
}
