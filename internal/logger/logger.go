package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	constants "github.com/manish-sa/india-template/internal/constant"
	"github.com/manish-sa/india-template/internal/helper"
	"github.com/manish-sa/india-template/pkg/sentry"
	"gitlab.dyninno.net/go-libraries/log"
	"moul.io/http2curl"
)

type LogData struct {
	RequestBody  map[string]interface{} `json:"requestBody,omitempty"`
	ResponseBody interface{}            `json:"responseBody,omitempty"`
	URL          string                 `json:"url,omitempty"`
	Params       url.Values             `json:"params,omitempty"`
	PathParams   string                 `json:"pathParams,omitempty"`
	Method       string                 `json:"method,omitempty"`
	Source       string                 `json:"source,omitempty"`
	Error        []interface{}          `json:"error,omitempty"`
}

func LogInfo(ctx context.Context, message string, response interface{}) {
	logData, ok := ctx.Value(constants.LogDataContextKey).(LogData)
	if !ok {
		log.FromCtx(ctx).Debug("Error asserting type of logData from ctx", "error", "")
	}

	logData.ResponseBody = response
	logData.Source = helper.GetStackTraceFrame().Function

	jsonData, err := json.Marshal(logData)
	if err != nil {
		log.FromCtx(ctx).Error("Error marshaling data to JSON", "error", err)
	}

	jsonString := string(jsonData)
	log.FromCtx(ctx).Info(message, "data", jsonString)
}

// LogWarn logs structured data at the warn level.
func LogWarn(ctx context.Context, message string, args ...interface{}) {
	logData, ok := ctx.Value(constants.LogDataContextKey).(LogData)
	if !ok {
		log.FromCtx(ctx).Debug("Error asserting type of logData from ctx", "error", "")
	}

	logData.Error = args
	logData.Source = helper.GetStackTraceFrame().Function

	jsonString := marshalToJSON(ctx, logData)

	fmt.Println(message, "warning details", jsonString)
	log.FromCtx(ctx).Warn(message, "warning details", jsonString)
}

// LogError logs structured data at the error level.
func LogError(ctx context.Context, message string, args ...interface{}) {
	logData, ok := ctx.Value(constants.LogDataContextKey).(LogData)
	if !ok {
		fmt.Println("Error asserting type of logData from ctx")
		log.FromCtx(ctx).Debug("Error asserting type of logData from ctx")
	}

	curl, ok := ctx.Value(constants.Curl).(*http2curl.CurlCommand)
	if !ok {
		fmt.Println("Error asserting curl from ctx")
		log.FromCtx(ctx).Debug("Error asserting curl from ctx")
	} else {
		ctx = sentry.WithKV(ctx, "curl", curl.String())
	}

	ctx = sentry.WithKV(ctx, "request data", logData)
	ctx = sentry.WithKV(ctx, "error", args[0])
	ctx = sentry.WithKV(ctx, "additional data", args[1:])

	logData.Error = args
	logData.Source = helper.GetStackTraceFrame().Function

	jsonString := marshalToJSON(ctx, logData)

	for i, s := range strings.Split(logData.Source, ".") {
		ctx = sentry.WithTag(ctx, fmt.Sprintf("source%d", i), s)
	}

	sentry.Error(ctx, message)
	fmt.Println(message, "error details", jsonString)
	log.FromCtx(ctx).Error(message, "error details", jsonString)
}

func marshalToJSON(ctx context.Context, args ...interface{}) string {
	jsonData, err := json.MarshalIndent(args, "", " ")
	if err != nil {
		log.FromCtx(ctx).Error("Error marshaling data to JSON", "error", err)
		return "[]"
	}

	return string(jsonData)
}

func (ld LogData) MarshalJSON() ([]byte, error) {
	type Alias LogData

	errorSerialized := struct {
		Error []interface{} `json:"error,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(&ld),
	}

	for _, v := range ld.Error {
		switch e := v.(type) {
		case string:
			errorSerialized.Error = append(errorSerialized.Error, e)
		case []string:
			if len(e) > 0 {
				errorSerialized.Error = append(errorSerialized.Error, e)
			} else {
				errorSerialized.Error = nil
			}
		default:
			errorSerialized.Error = nil
		}
	}

	xb, err := json.Marshal(errorSerialized)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	return xb, nil
}
