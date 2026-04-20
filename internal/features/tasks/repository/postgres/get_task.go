package tasks_postgres_repository

import (
	"checklist/internal/core/domain"
	core_errors "checklist/internal/core/errors"
	core_postgres_pool "checklist/internal/core/pool"
	"context"
	"errors"
	"fmt"
)

func (r *TasksRepository) GetTask(
	ctx context.Context,
	id int,
) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, description, completed
	FROM checklist.tasks
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var taskModel TaskModel

	err := row.Scan(
		&taskModel.ID,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"task with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.Task{}, fmt.Errorf("scan error: %w", err)
	}
	taskDomain := taskDomainFromModel(taskModel)
	return taskDomain, nil
}
