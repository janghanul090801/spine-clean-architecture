package domain

import (
	"context"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/ent"
	"time"
)

type User struct {
	ID        ID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUserFromEnt(entity *ent.User) *User {
	return &User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
	}
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	GetByID(c context.Context, id *ID) (*User, error)
}
