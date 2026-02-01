package controller

import (
	"context"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"net/http"
)

type RefreshTokenController struct {
	refreshTokenUsecase domain.RefreshTokenUsecase
}

func NewRefreshTokenController(usecase domain.RefreshTokenUsecase) *RefreshTokenController {
	return &RefreshTokenController{
		refreshTokenUsecase: usecase,
	}
}

func (rtc *RefreshTokenController) RefreshToken(ctx context.Context, request *domain.RefreshTokenRequest) httpx.Response[domain.RefreshTokenResponse] {

	id, err := rtc.refreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, config.E.RefreshTokenSecret)
	if err != nil {
		return httpx.Response[domain.RefreshTokenResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // user not found
			},
		}
	}

	user, err := rtc.refreshTokenUsecase.GetUserByID(ctx, id)
	if err != nil {
		return httpx.Response[domain.RefreshTokenResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // user not found
			},
		}
	}

	accessToken, err := rtc.refreshTokenUsecase.CreateAccessToken(user, config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.RefreshTokenResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	refreshToken, err := rtc.refreshTokenUsecase.CreateRefreshToken(user, config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		return httpx.Response[domain.RefreshTokenResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return httpx.Response[domain.RefreshTokenResponse]{
		Body: refreshTokenResponse,
	}
}
