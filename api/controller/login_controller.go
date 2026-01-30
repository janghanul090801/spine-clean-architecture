package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/bootstrap"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	loginUsecase domain.LoginUsecase
	env          *bootstrap.Env
}

func NewLoginController(usecase domain.LoginUsecase, env *bootstrap.Env) *LoginController {
	return &LoginController{
		loginUsecase: usecase,
	}
}

func (lc *LoginController) Login(c *fiber.Ctx) error {
	ctx := c.Context()

	var request domain.LoginRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	user, err := lc.loginUsecase.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(domain.ErrorResponse{Message: "User not found with the given email"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return c.Status(http.StatusUnauthorized).JSON(domain.ErrorResponse{Message: "Invalid credentials"})
	}

	accessToken, err := lc.loginUsecase.CreateAccessToken(user, lc.env.AccessTokenSecret, lc.env.AccessTokenExpiryHour)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := lc.loginUsecase.CreateRefreshToken(user, lc.env.RefreshTokenSecret, lc.env.RefreshTokenExpiryHour)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(http.StatusOK).JSON(loginResponse)
}
