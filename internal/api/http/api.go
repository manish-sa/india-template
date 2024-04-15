package http

import (
	"net/http"

	"github.com/manish-sa/india-template/internal/api/http/employee"
	"github.com/manish-sa/india-template/internal/app/provider"

	"github.com/manish-sa/india-template/internal/api/http/oapi"
	"github.com/manish-sa/india-template/internal/config"
	"github.com/manish-sa/india-template/internal/info"
)

var _ oapi.StrictServerInterface = (*API)(nil)

type API struct {
	cfg              config.Config
	serviceProviders *provider.ServiceProvider
	employee.APIEmployee
}

func NewAPI(
	cfg config.Config,
	serviceProviders *provider.ServiceProvider,
	empService employee.APIEmployee,
) *API {
	return &API{
		cfg:              cfg,
		serviceProviders: serviceProviders,
		APIEmployee:      empService,
	}
}

func (api *API) Live(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *API) Ready(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (api *API) Version(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(info.GetJSONInfo()); err != nil {
		// TODO: log
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
