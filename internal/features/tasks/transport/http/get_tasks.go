package tasks_transport_http

import (
	core_logger "checklist/internal/core/logger"
	core_http_response "checklist/internal/core/response"
	"net/http"
)

type GetTasksResponse []TasksDTOResponse

func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	tasksDomains, err := h.taskService.GetTasks(ctx)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get tasks",
		)
		return
	}
	response := GetTasksResponse(taskDTOFromDomains(tasksDomains))

	responseHandler.JSONResponse(response, http.StatusOK)
}
