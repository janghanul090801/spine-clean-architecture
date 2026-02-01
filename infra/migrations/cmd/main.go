package main

import (
	"context"
	"fmt"
	"github.com/janghanul090801/spine-clean-architecture/config"
	"github.com/janghanul090801/spine-clean-architecture/infra/database"
	"github.com/janghanul090801/spine-clean-architecture/infra/migrations"
	_ "github.com/janghanul090801/spine-clean-architecture/infra/migrations"
	"github.com/uptrace/bun/migrate"
)

func main() {
	ctx := context.Background()

	config.NewEnv()

	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	migrator := migrate.NewMigrator(db, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		panic(err)
	}

	group, err := migrator.Migrate(ctx)

	if err != nil {
		fmt.Printf("Migration failed: %v\n", err)
		return
	}

	if group.IsZero() {
		fmt.Printf("there are no new migrations to run\n")
		return
	}
	fmt.Printf("migrated to %s\n", group)
}
