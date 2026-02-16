package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"github.com/janghanul090801/spine-clean-architecture/infra/model"
	"github.com/uptrace/bun"
)

type taskRepository struct {
	db bun.IDB
}

func NewTaskRepository(db bun.IDB) domain.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(c context.Context, task *domain.Task) (*domain.Task, error) {

	taskModel := &model.TaskModel{
		ID:     uuid.New(),
		Title:  task.Title,
		UserID: task.UserID,
	}

	_, err := r.db.NewInsert().Model(taskModel).Exec(c)
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:        taskModel.ID,
		Title:     taskModel.Title,
		UserID:    taskModel.UserID,
		CreatedAt: taskModel.CreatedAt,
	}, nil
}

func (r *taskRepository) FetchByUserID(c context.Context, userID *domain.ID) ([]*domain.Task, error) {
	var taskModels []model.TaskModel
	err := r.db.NewSelect().Model(&taskModels).Where("user_id = ?", userID).Scan(c)
	if err != nil {
		return nil, err
	}

	tasks := make([]*domain.Task, len(taskModels))
	for i, t := range taskModels {
		tasks[i] = &domain.Task{
			ID:        t.ID,
			Title:     t.Title,
			UserID:    t.UserID,
			CreatedAt: t.CreatedAt,
		}
	}

	return tasks, err
}
