package usecase

import (
	"context"
	"time"

	"github.com/janghanul090801/spine-clean-architecture/domain"
)

type profileUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUseCase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUseCase {
	return &profileUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUseCase) GetProfileByID(c context.Context, userID *domain.ID) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{Name: user.Name, Email: user.Email}, nil
}
