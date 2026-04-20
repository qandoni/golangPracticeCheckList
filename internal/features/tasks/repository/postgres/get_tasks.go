package tasks_postgres_repository

import (
	"checklist/internal/core/domain"
	"context"
	"fmt"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, title, description, completed
	FROM checklist.tasks
	ORDER BY id ASC;
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return []domain.Task{}, fmt.Errorf("select tasks: %w", err)
	}
	defer rows.Close()

	var taskModels []TaskModel

	for rows.Next() {
		var taskModel TaskModel
		err := rows.Scan(
			&taskModel.ID,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
		)
		if err != nil {
			return nil, fmt.Errorf("scan tasks: %w", err)
		}
		taskModels = append(taskModels, taskModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("scan tasks: %w", err)
	}
	taskDomains := taskDomainsFromModels(taskModels)
	return taskDomains, nil
}
