package http

import (
	"context"

	"gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/api/http/oapi"
)

//nolint:revive
func (api *API) GetPing(
	_ context.Context, _ oapi.GetPingRequestObject,
) (oapi.GetPingResponseObject, error) {
	return oapi.GetPing200JSONResponse{Status: "ok"}, nil
}
