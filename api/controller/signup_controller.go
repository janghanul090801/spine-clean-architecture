package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/bootstrap"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	signupUsecase domain.SignupUsecase
	env           *bootstrap.Env
}

func NewSignupController(usecase domain.SignupUsecase, env *bootstrap.Env) *SignupController {
	return &SignupController{
		signupUsecase: usecase,
		env:           env,
	}
}

func (sc *SignupController) Signup(c *fiber.Ctx) error {
	ctx := c.Context()
	var request domain.SignupRequest

	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	_, err = sc.signupUsecase.GetUserByEmail(ctx, request.Email)
	if err == nil {
		return c.Status(http.StatusConflict).JSON(domain.ErrorResponse{Message: "User already exists with the given email"})
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.signupUsecase.Create(ctx, &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	accessToken, err := sc.signupUsecase.CreateAccessToken(&user, sc.env.AccessTokenSecret, sc.env.AccessTokenExpiryHour)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := sc.signupUsecase.CreateRefreshToken(&user, sc.env.RefreshTokenSecret, sc.env.RefreshTokenExpiryHour)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.Status(http.StatusOK).JSON(signupResponse)
}
