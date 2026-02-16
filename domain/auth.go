package domain

import "context"

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthUseCase interface {
	Register(c context.Context, name, email, password string) (*User, error)
	Login(c context.Context, email, password string) (*User, error)
	CreateAccessAndRefreshToken(c context.Context, user *User) (string, string, error)
	ExtractUserFromRefreshToken(c context.Context, requestToken string) (*User, error)
}
