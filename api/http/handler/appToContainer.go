package handler//click2cloud-apptocontainer

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

// AppToContainerHandler represents an HTTP API handler for managing AppToContainer
type AppToContainerHandler struct {
	*mux.Router
	Logger                            *log.Logger
	authorizeAppToContainerManagement bool
	AppToContainerService             dockm.AppToContainerService
	EndpointService					  dockm.EndpointService
}

const (
	// ErrEndpointManagementDisabled is an error raised when trying to access the endpoints management endpoints
	// when the server has been started with the --external-endpoints flag
	ErrapptocManagementDisabled = dockm.Error("Endpoint management is disabled")
)

type (
 postAppToContainerRequest struct {
	BaseImage string `valid:"required"`
	GitUrl string `valid:"required"`
	ImageName string `valid:"required"`
	EndPointId int `valid:"required"`
	 Output string `json:"required"`
}

	postAppToContainerResponse struct {
		Output string `json:"Output"`
	}
)

// NewEndpointHandler returns a new instance of EndpointHandler.
func NewAppToContainerHandler(bouncer *security.RequestBouncer, authorizeAppToContainerManagement bool, ) *AppToContainerHandler {
	h := &AppToContainerHandler{
		Router: mux.NewRouter(),
		Logger: log.New(os.Stderr, "", log.LstdFlags),
		authorizeAppToContainerManagement: authorizeAppToContainerManagement,
	}
	h.Handle("/apptocontainer",
		bouncer.AdministratorAccess(http.HandlerFunc(h.handlePostAppToContainer))).Methods(http.MethodPost)
	//h.Handle("/apptocontainer",
	//	bouncer.RestrictedAccess(http.HandlerFunc(h.handleGetAppToContainer))).Methods(http.MethodGet)
	return h

}

// handlePostAppToContainer handles POST requests on /AppToContainer
func (handler *AppToContainerHandler) handlePostAppToContainer(w http.ResponseWriter, r *http.Request) {
	//if !handler.authorizeAppToContainerManagement {
	//	httperror.WriteErrorResponse(w, ErrapptocManagementDisabled, http.StatusServiceUnavailable, handler.Logger)
	//	return
	//}

	var req postAppToContainerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.WriteErrorResponse(w, ErrInvalidJSON, http.StatusBadRequest, handler.Logger)
		return
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		httperror.WriteErrorResponse(w, ErrInvalidRequestFormat, http.StatusBadRequest, handler.Logger)
		return
	}

	atoc := &dockm.AToC{
		BaseImage:req.BaseImage,
		GitUrl:req.GitUrl,
		ImageName:req.ImageName,
		EndPointId:req.EndPointId,

	}

	endpoint, err := handler.EndpointService.Endpoint(dockm.EndpointID(atoc.EndPointId))

	if err == dockm.ErrEndpointNotFound {
		httperror.WriteErrorResponse(w, err, http.StatusNotFound, handler.Logger)
		return
	} else if err != nil {
		httperror.WriteErrorResponse(w, err, http.StatusInternalServerError, handler.Logger)
		return
	}

	err , output := handler.AppToContainerService.BuildAppToContainer(atoc,endpoint)
	fmt.Println(output)
	if err != nil {
		httperror.WriteErrorResponse(w, err, http.StatusBadRequest, handler.Logger)
		return
	}

	encodeJSON(w, &postAppToContainerResponse{Output: output}, handler.Logger)
}

