package repository

import (
	"context"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/ent"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/ent/user"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) domain.UserRepository {
	return &userRepository{
		client: client,
	}
}

func (r *userRepository) Create(c context.Context, user *domain.User) error {
	_, err := r.client.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		Save(c)

	return err
}

func (r *userRepository) Fetch(c context.Context) ([]*domain.User, error) {
	u, err := r.client.User.Query().All(c)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(u))
	for i, v := range u {
		users[i] = domain.NewUserFromEnt(v)
	}

	return users, err
}

func (r *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	u, err := r.client.User.Query().Where(
		user.EmailEQ(email),
	).Only(c)
	if err != nil {
		return nil, err
	}

	return domain.NewUserFromEnt(u), err
}

func (r *userRepository) GetByID(c context.Context, id *domain.ID) (*domain.User, error) {
	u, err := r.client.User.Get(c, *id)
	if err != nil {
		return nil, err
	}

	return domain.NewUserFromEnt(u), err
}
