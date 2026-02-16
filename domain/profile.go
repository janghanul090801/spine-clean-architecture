package domain

import "context"

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileUseCase interface {
	GetProfileByID(c context.Context, userID *ID) (*Profile, error)
}
