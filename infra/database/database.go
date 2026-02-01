package database

import (
	"context"
	"crypto/tls"
	"database/sql"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/internal/logger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.uber.org/zap"
	"time"
)

// NewDB returns an orm client
func NewDB() (*bun.DB, error) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: config.E.DBTlsSkipVerify,
	}

	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(config.E.DBHost+":"+config.E.DBPort),
		pgdriver.WithTLSConfig(tlsConfig),
		pgdriver.WithUser(config.E.DBUser),
		pgdriver.WithPassword(config.E.DBPass),
		pgdriver.WithDatabase(config.E.DBName),
		pgdriver.WithApplicationName(config.E.DBApplicationName),
		pgdriver.WithTimeout(10*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(10*time.Second),
		pgdriver.WithWriteTimeout(10*time.Second),
		pgdriver.WithInsecure(true),
	)

	sqldb := sql.OpenDB(pgconn)

	sqldb.SetMaxOpenConns(25)
	sqldb.SetMaxIdleConns(5)
	sqldb.SetConnMaxLifetime(5 * time.Minute)
	sqldb.SetConnMaxIdleTime(10 * time.Minute)

	db := bun.NewDB(sqldb, pgdialect.New())

	if config.E.Debug {
		db.AddQueryHook(&debugHook{})
	}

	return db, nil
}

type debugHook struct{}

func (h *debugHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *debugHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	getLogger := logger.GetLogger()
	getLogger.Debug("Database query",
		zap.String("query", event.Query),
		zap.Duration("duration", time.Since(event.StartTime)),
	)
}
