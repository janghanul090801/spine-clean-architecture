package bootstrap

import (
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/ent"
	_ "github.com/lib/pq"
)

// New returns data source name
func New(env *Env) string {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost,
		env.DBPort,
		env.DBUser,
		env.DBPass,
		env.DBName,
	)

	return dsn
}

// NewClient returns an orm client
func NewClient(env *Env) (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	dsn := New(env)

	return ent.Open(dialect.Postgres, dsn, entOptions...)
}
