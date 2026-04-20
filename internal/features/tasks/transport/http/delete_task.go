package tasks_transport_http

import (
	core_logger "checklist/internal/core/logger"
	core_http_request "checklist/internal/core/request"
	core_http_response "checklist/internal/core/response"
	"net/http"
)

func (h *TasksHTTPHandler) DeleteTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task id path value",
		)
		return
	}
	if err := h.taskService.DeleteTask(ctx, taskID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete task",
		)
		return
	}
	responseHandler.NoContentResponse()
}
