package controller

import (
	"context"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	loginUsecase domain.LoginUsecase
}

func NewLoginController(usecase domain.LoginUsecase) *LoginController {
	return &LoginController{
		loginUsecase: usecase,
	}
}

func (lc *LoginController) Login(ctx context.Context, request *domain.LoginRequest) httpx.Response[domain.LoginResponse] {

	user, err := lc.loginUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return httpx.Response[domain.LoginResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusNotFound, // User not found with the given email
			},
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return httpx.Response[domain.LoginResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // Invalid credentials
			},
		}
	}

	accessToken, err := lc.loginUsecase.CreateAccessToken(user, config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.LoginResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	refreshToken, err := lc.loginUsecase.CreateRefreshToken(user, config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.LoginResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return httpx.Response[domain.LoginResponse]{
		Body: loginResponse,
	}
}
