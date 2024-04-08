package gmail

import (
	"context"
	"fmt"

	gmail "gitlab.dyninno.net/go-libraries/client-component-gmail"
)

type GmailServiceInterface interface {
	Ping() error
}

type GmailService struct {
	client *gmail.Client
	ctx    context.Context
}

func NewGmailClient(ctx context.Context) GmailServiceInterface {
	return &GmailService{
		ctx:    ctx,
		client: gmail.NewClient(),
	}
}

func (gs *GmailService) Ping() error {
	err := gs.client.PingContext(gs.ctx)
	if err != nil {
		return fmt.Errorf("gmail client error: %w", err)
	}

	return nil
}
