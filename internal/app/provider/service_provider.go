package provider

import (
    "context"
    
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/client/gmail"
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/service/healthcheck"
    
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/config"
    "gorm.io/gorm"
)

type Clients struct {
    GmailClient gmail.GmailServiceInterface
}

type Services struct {
    HealthCheckService healthcheck.HealthcheckInterface
}

type ServiceProvider struct {
    *Clients
    *Services
}

func NewServiceProviders(ctx context.Context, cfg config.Config, db *gorm.DB) *ServiceProvider {
    clients := Clients{
        GmailClient: gmail.NewGmailClient(ctx),
    }
    
    serviceProvider := &ServiceProvider{
        Clients: &clients,
    }
    
    serviceProvider.HealthCheckService = serviceProvider.InitHealthcheckServiceInstance(ctx, db)
    
    return serviceProvider
}

// InitHealthcheckServiceInstance Method to initialize the healthcheck service instance
func (sp *ServiceProvider) InitHealthcheckServiceInstance(
    ctx context.Context,
    db *gorm.DB,
) healthcheck.HealthcheckInterface {
    if sp.HealthCheckService == nil {
        sp.HealthCheckService = healthcheck.NewHealthcheckService(
            ctx,
            db,
            sp.Clients.GmailClient,
        )
    }
    
    return sp.HealthCheckService
}
