package domain

import (
	"context"
	"time"
)

type User struct {
	ID        ID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserRepository interface {
	Create(c context.Context, user *User) (*User, error)
	Fetch(c context.Context) ([]*User, error)
	GetByEmail(c context.Context, email string) (*User, error)
	GetByID(c context.Context, id *ID) (*User, error)
}
