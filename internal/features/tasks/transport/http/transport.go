package tasks_transport_http

import (
	"checklist/internal/core/domain"
	core_http_server "checklist/internal/core/server"
	"context"
	"net/http"
)

type TasksHTTPHandler struct {
	taskService TaskService
}

type TaskService interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
	GetTasks(
		ctx context.Context,
	) ([]domain.Task, error)
	GetTask(
		ctx context.Context,
		taskID int,
	) (domain.Task, error)
	DeleteTask(
		ctx context.Context,
		taskID int,
	) error
	CompleteTask(
		ctx context.Context,
		taskID int,
		completed bool,
	) (domain.Task, error)
}

func NewTasksHTTPHandler(
	taskService TaskService,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		taskService: taskService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks",
			Handler: h.GetTasks,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/tasks/{id}",
			Handler: h.DeleteTask,
		},
		{
			Method:  http.MethodGet,
			Path:    "/tasks/{id}",
			Handler: h.GetTask,
		},
		{
			Method:  http.MethodPut,
			Path:    "/tasks/{id}",
			Handler: h.CompleteTask,
		},
	}
}
