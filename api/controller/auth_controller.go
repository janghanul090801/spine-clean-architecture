package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"golang.org/x/crypto/bcrypt"
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

	user, err := c.service.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusNotFound, // User not found with the given email
			},
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // Invalid credentials
			},
		}
	}

	accessToken, err := c.service.CreateAccessToken(user, config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	refreshToken, err := c.service.CreateRefreshToken(user, config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
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

	id, err := c.service.ExtractIDFromRefreshToken(request.RefreshToken, config.E.RefreshTokenSecret)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // user not found
			},
		}
	}

	user, err := c.service.GetUserByID(ctx, id)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // user not found
			},
		}
	}

	accessToken, err := c.service.CreateAccessToken(user, config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	refreshToken, err := c.service.CreateRefreshToken(user, config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
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
	_, err := c.service.GetUserByEmail(ctx, request.Email)
	if err == nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusConflict, // User already exists with the given email
			},
		}
	}

	err = c.service.Create(ctx, request.Name, request.Email, request.Password)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	user, err := c.service.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	accessToken, err := c.service.CreateAccessToken(user, config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	refreshToken, err := c.service.CreateRefreshToken(user, config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.AuthResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
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
