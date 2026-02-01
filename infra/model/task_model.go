package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"github.com/uptrace/bun"
	"time"
)

type TaskModel struct {
	bun.BaseModel `bun:"table:tasks"`

	ID        domain.ID `bun:"id,pk,type:uuid"`
	Title     string    `bun:"title,notnull"`
	UserID    domain.ID `bun:"user_id,notnull"`
	User      UserModel `bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt time.Time `bun:"default:now()"`
}

func (u *TaskModel) BeforeInsert(
	ctx context.Context,
	q *bun.InsertQuery,
) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
