package tasks_transport_http

import "checklist/internal/core/domain"

type TasksDTOResponse struct {
	ID          int     `json:"ID"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Completed   bool    `json:"completed"`
}

func taskDTOFromDomain(task domain.Task) TasksDTOResponse {
	return TasksDTOResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}
}

func taskDTOFromDomains(tasks []domain.Task) []TasksDTOResponse {
	dtos := make([]TasksDTOResponse, len(tasks))
	for i, task := range tasks {
		dtos[i] = taskDTOFromDomain(task)
	}
	return dtos
}
