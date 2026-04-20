package tasks_service

import (
	"checklist/internal/core/domain"
	"context"
	"fmt"
)

func (s *TasksService) CompleteTask(
	ctx context.Context,
	id int,
	completed bool,
) (domain.Task, error) {
	task, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("get task: %w", err)
	}
	task.SetCompleted(completed)

	patchedTask, err := s.tasksRepository.CompleteTask(ctx, id, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("patch user: %w", err)
	}
	return patchedTask, nil
}
