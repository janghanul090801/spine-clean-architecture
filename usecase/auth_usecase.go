package usecase

import (
	"context"
	"time"

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

func (u *authUseCase) Create(c context.Context, name, email, password string) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	return u.userRepository.Create(ctx, &domain.User{
		Name:     name,
		Email:    email,
		Password: string(encryptedPassword),
	})
}

func (u *authUseCase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByEmail(ctx, email)
}

func (u *authUseCase) GetUserByID(c context.Context, id *domain.ID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.userRepository.GetByID(ctx, id)
}

func (u *authUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return token.CreateAccessToken(user, secret, expiry)
}

func (u *authUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return token.CreateRefreshToken(user, secret, expiry)
}

func (u *authUseCase) ExtractIDFromRefreshToken(requestToken string, secret string) (*domain.ID, error) {
	return token.ExtractIDFromToken(requestToken, secret)
}
