package tasks_postgres_repository

import (
	core_errors "checklist/internal/core/errors"
	"context"
	"fmt"
)

func (r *TasksRepository) DeleteTask(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM checklist.tasks
	WHERE id=$1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"task with id='%d': %w",
			id,
			core_errors.ErrNotFound,
		)
	}
	return nil
}
