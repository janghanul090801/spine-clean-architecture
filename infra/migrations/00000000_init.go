package migrations

import (
	"context"
	"fmt"
	"github.com/janghanul090801/spine-clean-architecture/infra/model"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up]  ")
		_, err := db.NewCreateTable().Model((*model.UserModel)(nil)).Exec(ctx)
		_, err = db.NewCreateTable().Model((*model.TaskModel)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down]  ")
		_, err := db.NewDropTable().Model((*model.UserModel)(nil)).IfExists().Exec(ctx)
		_, err = db.NewDropTable().Model((*model.TaskModel)(nil)).IfExists().Exec(ctx)
		return err
	})
}
