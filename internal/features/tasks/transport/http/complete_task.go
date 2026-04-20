package tasks_transport_http

import (
	core_logger "checklist/internal/core/logger"
	core_http_request "checklist/internal/core/request"
	core_http_response "checklist/internal/core/response"
	"net/http"
)

type CompleteTaskRequest struct {
	Complete bool `json:"completed"`
}

type CompleteTaskResponse TasksDTOResponse

func (h *TasksHTTPHandler) CompleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get int path value",
		)
		return
	}

	var request CompleteTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)
		return
	}
	taskCompleteDomain, err := h.taskService.CompleteTask(ctx, taskID, request.Complete)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to complete task",
		)
		return
	}

	response := CompleteTaskResponse(taskCompleteDomain)

	responseHandler.JSONResponse(response, http.StatusOK)
}
