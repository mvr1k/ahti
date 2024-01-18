package reminders

import (
	"ahti/app/internal/web/api"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

const (
	ModuleName = "reminder"
)

type ApiController struct {
	svc    *service
	logger *zap.SugaredLogger
}

func NewApiController(logger *zap.SugaredLogger) *ApiController {
	return &ApiController{
		svc:    newService(logger),
		logger: logger,
	}
}

func (ctrl *ApiController) GetRoutingList() []api.RouteInfo {
	return []api.RouteInfo{
		api.NewRouteInfo(ModuleName, http.MethodGet, ctrl.GET),
		api.NewRouteInfo(ModuleName, http.MethodPost, ctrl.POST),
		api.NewRouteInfo(ModuleName, http.MethodPut, ctrl.PUT),
		api.NewRouteInfo(ModuleName, http.MethodDelete, ctrl.DELETE),
	}
}

func (ctrl *ApiController) ModuleName() string {
	return ModuleName
}

func (ctrl *ApiController) POST(writer http.ResponseWriter, request *http.Request) {
	req, err := NewReminderFromHttpRequest(request)
	if err != nil {
		errStr := fmt.Sprintf("Error while parsing request: %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusInternalServerError)
		return
	}

	err = req.Validate()
	if err != nil {
		errStr := fmt.Sprintf("Error while validating request: %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusBadRequest)
		return
	}

	err = ctrl.svc.createReminder(req)
	if err != nil {
		errStr := fmt.Sprintf("Error while Creating Reminder : %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusInternalServerError)
		return
	}

	RespondToWriter(writer, true, "Reminder Created Successfully :) ", nil, http.StatusOK)
}
func (ctrl *ApiController) PUT(writer http.ResponseWriter, request *http.Request) {

	id := NewIdFromHttpRequest(request)

	err := id.Validate()
	if err != nil {
		errStr := fmt.Sprintf("Error while validating request: %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusBadRequest)
		return
	}

	req, err := NewReminderFromHttpRequest(request)
	if err != nil {
		errStr := fmt.Sprintf("Error while parsing request: %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusInternalServerError)
		return
	}

	err = req.Validate()
	if err != nil {
		errStr := fmt.Sprintf("Error while validating request: %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusBadRequest)
		return
	}

	err = ctrl.svc.updateReminder(id, req)
	if err != nil {
		errStr := fmt.Sprintf("Error while Creating Reminder : %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusInternalServerError)
		return
	}

	RespondToWriter(writer, true, "Reminder Updated Successfully :) ", nil, http.StatusOK)
}
func (ctrl *ApiController) DELETE(writer http.ResponseWriter, request *http.Request) {
	id := NewIdFromHttpRequest(request)

	err := id.Validate()
	if err != nil {
		errStr := fmt.Sprintf("Error while validating request: %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusBadRequest)
		return

	}

	err = ctrl.svc.deleteReminder(id)
	if err != nil {
		errStr := fmt.Sprintf("Error while Deleting Reminder : %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusInternalServerError)
		return
	}

	RespondToWriter(writer, true, "Reminder Deleted Successfully :( ", nil, http.StatusOK)

}
func (ctrl *ApiController) GET(writer http.ResponseWriter, request *http.Request) {
	id := NewIdFromHttpRequest(request)

	list, err := ctrl.svc.getReminder(id)
	if err != nil {
		errStr := fmt.Sprintf("Error while Fetching Reminder : %v", err)
		ctrl.logger.Error(errStr)
		RespondToWriter(writer, false, errStr, nil, http.StatusInternalServerError)
		return
	}

	RespondToWriter(writer, true, "Reminder Fetched Successfully", list, http.StatusOK)
}
