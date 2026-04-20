package tasks_service

import (
	"checklist/internal/core/domain"
	"context"
	"fmt"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
) ([]domain.Task, error) {
	tasks, err := s.tasksRepository.GetTasks(ctx)
	if err != nil {
		return []domain.Task{}, fmt.Errorf("get tasks from repository: %w", err)
	}
	return tasks, nil
}
