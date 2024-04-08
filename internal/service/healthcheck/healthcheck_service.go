package healthcheck

import (
    "context"
    "net/http"
    "time"
    
    clientcomponentgbo "gitlab.dyninno.net/go-libraries/client-component-gbo"
    "gitlab.dyninno.net/trevolution/ancillaries/lbp/lbc-service/internal/client/gmail"
    
    "gorm.io/gorm"
)

type GBOClientInterface interface {
    Feed(ctx context.Context, containerID string) (*clientcomponentgbo.ObjectFeedResponse, error)
    Ping(ctx context.Context) (*clientcomponentgbo.PingResponse, error)
}

type HealthcheckInterface interface {
    GetDBStatus() HealthResponse
    GetAppStatus() HealthResponse
    GetRedisStatus() HealthResponse
    GetGmailClientStatus() HealthResponse
}

// HealthcheckService handles health check operations.
type HealthcheckService struct {
    ctx         context.Context
    db          *gorm.DB
    gmailClient gmail.GmailServiceInterface
}

// HealthResponse represents the response for a health check.
type HealthResponse struct {
    Status     bool      `json:"status"`
    StatusCode int       `json:"statusCode"`
    Message    string    `json:"message"`
    Date       time.Time `json:"date"`
}

// HealthModel represents the model used for health checks.
type HealthModel struct {
    gorm.Model
    Field int
}

const (
    CLAIMApp         = "lbc_app"
    CLAIMDB          = "lbc_db"
    CLAIMRedis       = "lbc_redis"
    CLAIMGmail       = "lbc_gmail"
    CLAIMServiceList = "serviceList"
)

type HealthcheckResponse map[string]interface{}

type StatusFunc func() HealthResponse

func NewHealthcheckService(
    ctx context.Context,
    db *gorm.DB,
    gmailClient gmail.GmailServiceInterface,
) HealthcheckInterface {
    return &HealthcheckService{
        ctx:         ctx,
        db:          db,
        gmailClient: gmailClient,
    }
}

func (s *HealthcheckService) GetDBStatus() HealthResponse {
    var result int
    err := s.db.Raw("SELECT 1").Scan(&result).Error
    
    if err == nil && result == 1 {
        return formatSuccessResponse()
    }
    
    return formatErrorResponse("fail")
}

func (s *HealthcheckService) GetAppStatus() HealthResponse {
    return formatSuccessResponse()
}

func (s *HealthcheckService) GetRedisStatus() HealthResponse {
    // add actual implantation
    return formatSuccessResponse()
}

func (s *HealthcheckService) GetGmailClientStatus() HealthResponse {
    if err := s.gmailClient.Ping(); err != nil {
        return formatErrorResponse(err.Error())
    }
    
    return formatSuccessResponse()
}

func formatSuccessResponse() HealthResponse {
    return HealthResponse{
        Status:     true,
        StatusCode: http.StatusOK,
        Message:    "success",
        Date:       time.Now().UTC(),
    }
}

func formatErrorResponse(message string) HealthResponse {
    return HealthResponse{
        Status:     false,
        StatusCode: http.StatusNotFound,
        Message:    message,
        Date:       time.Now().UTC(),
    }
}
