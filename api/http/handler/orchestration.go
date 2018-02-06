package handler

import (
	"dolphin/api"
	httperror "dolphin/api/http/error"
	"dolphin/api/http/security"

	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"

	"fmt"
)

// OrchestrationHandler represents an HTTP API handler for managing Orchestration
type OrchestrationHandler struct {
	*mux.Router
	Logger                            *log.Logger
	authorizeOrchestrationManagement  bool
	OrchestrationService              dockm.OrchestrationService
	EndpointService					  dockm.EndpointService
}

const (
	ErrOrchestrationManagementDisabled = dockm.Error("Endpoint management is disabled")
)

type (
	postOrchestrationRequest struct {
		EndPointId int `valid:"required"`
		Output string `json:"required"`
	}

	postOrchestrationResponse struct {
		Output string `json:"Output"`
	}
)

// NewOrchestrationHandler returns a new instance of OrchestrationHandler.
func NewOrchestrationHandler(bouncer *security.RequestBouncer, authorizeOrchestrationManagement bool, ) *OrchestrationHandler {
	h := &OrchestrationHandler{
		Router: mux.NewRouter(),
		Logger: log.New(os.Stderr, "", log.LstdFlags),
		authorizeOrchestrationManagement: authorizeOrchestrationManagement,
	}
	h.Handle("/orchestration",
		bouncer.AdministratorAccess(http.HandlerFunc(h.handlePostOrchestration))).Methods(http.MethodPost)
	return h

}

// handlePostOrchestration handles POST requests on /AppToContainer
func (handler *OrchestrationHandler) handlePostOrchestration(w http.ResponseWriter, r *http.Request) {

	var req postOrchestrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.WriteErrorResponse(w, ErrInvalidJSON, http.StatusBadRequest, handler.Logger)
		return
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		httperror.WriteErrorResponse(w, ErrInvalidRequestFormat, http.StatusBadRequest, handler.Logger)
		return
	}

	otos := &dockm.OToS{
		EndPointId:req.EndPointId,
	}

	endpoint, err := handler.EndpointService.Endpoint(dockm.EndpointID(otos.EndPointId))

	if err == dockm.ErrEndpointNotFound {
		httperror.WriteErrorResponse(w, err, http.StatusNotFound, handler.Logger)
		return
	} else if err != nil {
		httperror.WriteErrorResponse(w, err, http.StatusInternalServerError, handler.Logger)
		return
	}

	err , output := handler.OrchestrationService.BuildOrchestration(otos,endpoint)
	fmt.Println(output)
	if err != nil {
		httperror.WriteErrorResponse(w, err, http.StatusBadRequest, handler.Logger)
		return
	}

	encodeJSON(w, &postOrchestrationResponse{Output: output}, handler.Logger)
}

