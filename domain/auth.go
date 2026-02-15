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
	Create(c context.Context, name, email, password string) error
	GetUserByEmail(c context.Context, email string) (*User, error)
	GetUserByID(c context.Context, id *ID) (*User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromRefreshToken(requestToken string, secret string) (*ID, error)
}
