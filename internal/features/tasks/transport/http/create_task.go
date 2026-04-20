package tasks_transport_http

import (
	"checklist/internal/core/domain"
	core_logger "checklist/internal/core/logger"
	core_http_request "checklist/internal/core/request"
	core_http_response "checklist/internal/core/response"
	"net/http"
)

type CreateTaskRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=100" example:"Домашнее задание"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000" example:"Сделать дз до завтра"`
}

type CreateTaskResponse TasksDTOResponse

func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
	}

	taskDomain := domain.NewTaskUnitialized(
		request.Title,
		request.Description,
	)

	taskDomain, err := h.taskService.CreateTask(ctx, taskDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create task",
		)
		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))
	responseHandler.JSONResponse(response, http.StatusCreated)
}
