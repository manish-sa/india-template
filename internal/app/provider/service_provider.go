package provider

import (
	"context"

	"github.com/manish-sa/india-template/internal/client/gmail"
	"github.com/manish-sa/india-template/internal/service/healthcheck"
	"gitlab.dyninno.net/go-libraries/dyninnogorm"
	"gorm.io/gorm"

	"github.com/manish-sa/india-template/internal/config"
)

type Clients struct {
	GmailClient gmail.GmailServiceInterface
}

type Services struct {
	HealthCheckService healthcheck.HealthcheckInterface
}

type ServiceProvider struct {
	ctx context.Context
	cfg config.Config
	db  *gorm.DB
	Clients
	Services
}

func NewServiceProviders(ctx context.Context, cfg config.Config) *ServiceProvider {
	sp := &ServiceProvider{
		ctx: ctx,
		cfg: cfg,
	}

	sp.db = sp.initDbClient()

	clients := Clients{
		GmailClient: sp.initGmailClient(),
	}

	services := Services{
		HealthCheckService: sp.initHealthcheckServiceInstance(),
	}

	sp.Clients = clients
	sp.Services = services

	return sp
}

func (sp *ServiceProvider) initHealthcheckServiceInstance() healthcheck.HealthcheckInterface {
	if sp.HealthCheckService == nil {
		sp.HealthCheckService = healthcheck.NewHealthcheckService(
			sp.ctx,
			sp.initDbClient(),
			sp.initGmailClient(),
		)
	}

	return sp.HealthCheckService
}

func (sp *ServiceProvider) initGmailClient() gmail.GmailServiceInterface {
	if sp.GmailClient == nil {
		sp.GmailClient = gmail.NewGmailClient(sp.ctx)
	}

	return sp.GmailClient
}

func (sp *ServiceProvider) initDbClient() *gorm.DB {
	if sp.db == nil {
		err := dyninnogorm.Init(nil)
		if err != nil {
			panic(err)
		}

		sp.db = dyninnogorm.Inst(sp.ctx)
	}

	return sp.db
}
