package constants

type (
	contextKey string
	Source     string
)

const (
	ProductName             = "Lost Baggage Claim"
	ProductType             = "lbc"
	CreatedBy               = 123
	AllowedRoleName         = "ancillaryProductRefundManager"
	DefaultRefundableAmount = 0
	MaxBagCount             = 2
	CompensationPaid        = "compensationPaid"
	XLogId                  = "X-Log-Id"
	Curl                    = contextKey("curl")
	LogDataContextKey       = contextKey("logData")
)
