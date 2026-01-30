package domain

import (
	"context"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/ent"
	"time"
)

type Task struct {
	ID        ID        `json:"id"`
	Title     string    `form:"title" binding:"required" json:"title"`
	UserID    ID        `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTaskFromEnt(entity *ent.Task) *Task {
	return &Task{
		ID:        entity.ID,
		Title:     entity.Title,
		UserID:    entity.Edges.Owner.ID,
		CreatedAt: entity.CreatedAt,
	}
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID *ID) ([]*Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID *ID) ([]*Task, error)
}
