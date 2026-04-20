package tasks_postgres_repository

import (
	"checklist/internal/core/domain"
	core_errors "checklist/internal/core/errors"
	"context"
	"fmt"
)

func (r *TasksRepository) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO checklist.tasks (title, description, completed)
	VALUES ($1, $2, $3)
	RETURNING id, title, description, completed;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
	)

	var (
		taskModel TaskModel
	)

	err := row.Scan(
		&taskModel.ID,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
	)

	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"scan error: %d: %w", err, core_errors.ErrNotFound,
		)
	}

	taskDomain := domain.NewTask(
		taskModel.ID,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
	)
	return taskDomain, nil
}
