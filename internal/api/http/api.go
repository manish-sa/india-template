package http

import (
	"net/http"

	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/app/provider"

	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/api/http/oapi"
	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/config"
	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/info"
)

var _ oapi.StrictServerInterface = (*API)(nil)

type API struct {
	cfg              config.Config
	serviceProviders *provider.ServiceProvider
}

func NewAPI(
	cfg config.Config,
	serviceProviders *provider.ServiceProvider,
) *API {
	return &API{
		cfg:              cfg,
		serviceProviders: serviceProviders,
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
