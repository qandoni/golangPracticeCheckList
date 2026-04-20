package tasks_postgres_repository

import "checklist/internal/core/domain"

type TaskModel struct {
	ID          int
	Title       string
	Description *string
	Completed   bool
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
	)
}

func taskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(taskModels))
	for i, model := range taskModels {
		domains[i] = taskDomainFromModel(model)
	}
	return domains
}
