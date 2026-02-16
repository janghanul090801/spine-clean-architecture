package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/janghanul090801/spine-clean-architecture/domain"
	"github.com/janghanul090801/spine-clean-architecture/infra/model"
	"github.com/janghanul090801/spine-clean-architecture/infra/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"

	_ "github.com/lib/pq"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	dsn := "postgres://postgres:postgres@localhost:5434/test_db?sslmode=disable"
	sqldb, err := sql.Open("postgres", dsn)
	require.NoError(t, err)

	db := bun.NewDB(sqldb, pgdialect.New())

	_, err = db.NewCreateTable().
		Model((*model.UserModel)(nil)).
		IfNotExists().
		Exec(ctx)
	require.NoError(t, err)

	_, err = db.NewTruncateTable().
		Model((*model.UserModel)(nil)).
		Exec(ctx)
	require.NoError(t, err)

	repo := repository.NewUserRepository(db)

	_, err = repo.Create(ctx, &domain.User{
		Name:     "hanul",
		Email:    "hanul@gmail.com",
		Password: "123456",
	})
	require.NoError(t, err)

	u, err := repo.GetByEmail(ctx, "hanul@gmail.com")
	require.NoError(t, err)
	assert.Equal(t, "hanul", u.Name)

	t.Cleanup(func() {
		_, _ = db.NewTruncateTable().Table("users").Exec(ctx)
	})
}
