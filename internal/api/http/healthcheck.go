package http

import (
    "net/http"
    
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/service/healthcheck"
)

func (api *API) Healthcheck(
    _ http.ResponseWriter,
    r *http.Request,
) map[string]interface{} {
    queryParams := r.URL.Query()
    var healthCheckType string
    
    for key, values := range queryParams {
        if len(values) > 0 {
            healthCheckType = key
            break
        }
    }
    
    healthcheckService := api.serviceProviders.HealthCheckService
    
    statusFuncs := map[string]healthcheck.StatusFunc{
        healthcheck.CLAIMDB:    healthcheckService.GetDBStatus,
        healthcheck.CLAIMApp:   healthcheckService.GetAppStatus,
        healthcheck.CLAIMRedis: healthcheckService.GetRedisStatus,
        healthcheck.CLAIMGmail: healthcheckService.GetGmailClientStatus,
    }
    
    dataMap := healthcheck.HealthcheckResponse{}
    
    if statusFunc, ok := statusFuncs[healthCheckType]; ok {
        dataMap[healthCheckType] = statusFunc()
    } else if healthCheckType == healthcheck.CLAIMServiceList {
        var list []string
        for key := range statusFuncs {
            list = append(list, key)
        }
        dataMap[healthCheckType] = list
    } else {
        for key, statusFunc := range statusFuncs {
            dataMap[key] = statusFunc()
        }
    }
    
    return healthcheck.HealthcheckResponse{
        "success": true,
        "message": []string{"Healthcheck response"},
        "data":    dataMap,
    }
}
