package controller

import (
	"context"
	"errors"

	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/janghanul090801/spine-clean-architecture/domain"
)

type AuthController struct {
	service domain.AuthUseCase
}

func NewAuthController(service domain.AuthUseCase) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (c *AuthController) Login(ctx context.Context, request *domain.LoginRequest) httpx.Response[domain.AuthResponse] {
	var errInfo domain.Error

	user, err := c.service.Login(ctx, request.Email, request.Password)

	accessToken, refreshToken, err := c.service.CreateAccessAndRefreshToken(ctx, user)
	if err != nil {
		if ok := errors.As(err, &errInfo); ok {
			return httpx.Response[domain.AuthResponse]{
				Options: httpx.ResponseOptions{
					Status: errInfo.StatusCode, // err.Error()
				},
			}
		}
	}

	response := domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return httpx.Response[domain.AuthResponse]{
		Body: response,
	}
}

func (c *AuthController) RefreshToken(ctx context.Context, request *domain.RefreshTokenRequest) httpx.Response[domain.AuthResponse] {
	var errInfo domain.Error

	user, err := c.service.ExtractUserFromRefreshToken(ctx, request.RefreshToken)
	if err != nil {
		if ok := errors.As(err, &errInfo); ok {
			return httpx.Response[domain.AuthResponse]{
				Options: httpx.ResponseOptions{
					Status: errInfo.StatusCode, // user not found
				},
			}
		}
	}

	accessToken, refreshToken, err := c.service.CreateAccessAndRefreshToken(ctx, user)
	if err != nil {
		if ok := errors.As(err, &errInfo); ok {
			return httpx.Response[domain.AuthResponse]{
				Options: httpx.ResponseOptions{
					Status: errInfo.StatusCode, // err.Error()
				},
			}
		}
	}

	response := domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return httpx.Response[domain.AuthResponse]{
		Body: response,
	}
}

func (c *AuthController) Signup(ctx context.Context, request *domain.SignupRequest) httpx.Response[domain.AuthResponse] {
	var errInfo domain.Error

	user, err := c.service.Register(ctx, request.Name, request.Email, request.Password)
	if err != nil {
		if ok := errors.As(err, &errInfo); ok {
			return httpx.Response[domain.AuthResponse]{
				Options: httpx.ResponseOptions{
					Status: errInfo.StatusCode, // err.Error()
				},
			}
		}
	}

	accessToken, refreshToken, err := c.service.CreateAccessAndRefreshToken(ctx, user)
	if err != nil {
		if ok := errors.As(err, &errInfo); ok {
			return httpx.Response[domain.AuthResponse]{
				Options: httpx.ResponseOptions{
					Status: errInfo.StatusCode, // err.Error()
				},
			}
		}
	}

	response := domain.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return httpx.Response[domain.AuthResponse]{
		Body: response,
	}
}
