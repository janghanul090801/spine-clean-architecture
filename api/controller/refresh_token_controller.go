package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/bootstrap"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"net/http"
)

type RefreshTokenController struct {
	refreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func NewRefreshTokenController(usecase domain.RefreshTokenUsecase, env *bootstrap.Env) *RefreshTokenController {
	return &RefreshTokenController{
		refreshTokenUsecase: usecase,
		Env:                 env,
	}
}

func (rtc *RefreshTokenController) RefreshToken(c *fiber.Ctx) error {
	ctx := c.Context()

	var request domain.RefreshTokenRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	id, err := rtc.refreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "User not found"})
	}

	user, err := rtc.refreshTokenUsecase.GetUserByID(ctx, id)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "User not found"})
	}

	accessToken, err := rtc.refreshTokenUsecase.CreateAccessToken(user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := rtc.refreshTokenUsecase.CreateRefreshToken(user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(http.StatusOK).JSON(refreshTokenResponse)
}
