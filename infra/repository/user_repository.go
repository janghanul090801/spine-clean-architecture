package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"github.com/janghanul090801/spine-clean-architecture/infra/model"
	"github.com/uptrace/bun"
)

type userRepository struct {
	db bun.IDB
}

func NewUserRepository(db bun.IDB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(c context.Context, user *domain.User) error {
	userModel := &model.UserModel{
		ID:       uuid.New(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	_, err := r.db.NewInsert().Model(userModel).Exec(c)

	user.ID = userModel.ID

	return err
}

func (r *userRepository) Fetch(c context.Context) ([]*domain.User, error) {
	var userModels []model.UserModel
	err := r.db.NewSelect().Model(&userModels).Scan(c)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(userModels))
	for i, m := range userModels {
		users[i] = &domain.User{
			ID:        m.ID,
			Name:      m.Name,
			Email:     m.Email,
			Password:  m.Password,
			CreatedAt: m.CreatedAt,
		}
	}

	return users, err
}

func (r *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	var userModel model.UserModel
	err := r.db.NewSelect().Model(&userModel).Where("email = ?", email).Limit(1).Scan(c)

	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
	}

	return user, err
}

func (r *userRepository) GetByID(c context.Context, id *domain.ID) (*domain.User, error) {
	var userModel model.UserModel
	err := r.db.NewSelect().Model(&userModel).Where("id = ?", *id).Limit(1).Scan(c)

	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:        userModel.ID,
		Name:      userModel.Name,
		Email:     userModel.Email,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
	}

	return user, err
}
