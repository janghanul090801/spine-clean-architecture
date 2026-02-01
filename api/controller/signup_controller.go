package controller

import (
	"context"
	"fmt"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	signupUsecase domain.SignupUsecase
}

func NewSignupController(usecase domain.SignupUsecase) *SignupController {
	return &SignupController{
		signupUsecase: usecase,
	}
}

func (sc *SignupController) Signup(ctx context.Context, request *domain.SignupRequest) httpx.Response[domain.SignupResponse] {
	_, err := sc.signupUsecase.GetUserByEmail(ctx, request.Email)
	if err == nil {
		return httpx.Response[domain.SignupResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusConflict, // User already exists with the given email
			},
		}
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.SignupResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.signupUsecase.Create(ctx, &user)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.SignupResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	accessToken, err := sc.signupUsecase.CreateAccessToken(&user, config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.SignupResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	refreshToken, err := sc.signupUsecase.CreateRefreshToken(&user, config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		fmt.Println(err.Error())
		return httpx.Response[domain.SignupResponse]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return httpx.Response[domain.SignupResponse]{
		Body: signupResponse,
	}
}
