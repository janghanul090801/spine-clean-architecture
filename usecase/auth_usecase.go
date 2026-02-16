package usecase

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"github.com/janghanul090801/spine-clean-architecture/internal/token"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewAuthUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.AuthUseCase {
	return &authUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (u *authUseCase) Register(c context.Context, name, email, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	_, err := u.userRepository.GetByEmail(ctx, email)
	if err == nil {
		return nil, domain.NewBadRequestError(err)
	}

	encrypted, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, domain.NewBadRequestError(err)
	}

	user, err := u.userRepository.Create(ctx, &domain.User{
		Name:     name,
		Email:    email,
		Password: string(encrypted),
	})
	if err != nil {
		return nil, domain.NewInternalServerError(err)
	}

	return user, nil
}

func (u *authUseCase) Login(c context.Context, email, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, domain.NewUnauthorizedError(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, domain.NewBadRequestError(errors.New("invalid credentials"))
	}

	return user, nil
}

func (u *authUseCase) CreateAccessAndRefreshToken(c context.Context, user *domain.User) (string, string, error) {
	access, err := token.CreateAccessToken(user, config.E.AccessTokenSecret, config.E.AccessTokenExpiryHour)
	if err != nil {
		return "", "", domain.NewInternalServerError(err)
	}

	refresh, err := token.CreateRefreshToken(user, config.E.RefreshTokenSecret, config.E.RefreshTokenExpiryHour)
	if err != nil {
		return "", "", domain.NewInternalServerError(err)
	}

	return access, refresh, nil
}

func (u *authUseCase) ExtractUserFromRefreshToken(c context.Context, requestToken string) (*domain.User, error) {
	id, err := token.ExtractIDFromToken(requestToken, config.E.RefreshTokenSecret)
	if err != nil {
		return nil, domain.Error{
			StatusCode: http.StatusBadRequest,
			Err:        err,
		}
	}

	user, err := u.userRepository.GetByID(c, id)
	if err != nil {
		return nil, domain.NewUnauthorizedError(err)
	}

	return user, nil
}
