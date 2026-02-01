package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users"`

	ID        uuid.UUID `bun:"id,pk,type:uuid"`
	Name      string    `bun:"name,notnull"`
	Email     string    `bun:"email,unique"`
	Password  string    `bun:"password"`
	CreatedAt time.Time `bun:"default:now()"`
}

func (u *UserModel) BeforeInsert(
	ctx context.Context,
	q *bun.InsertQuery,
) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
