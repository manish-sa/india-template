package helper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"
)

// GetStackTraceFrame prints the calling function name
func GetStackTraceFrame() (res sentry.Frame) {
	s := sentry.NewStacktrace()
	res = s.Frames[len(s.Frames)-3]

	return
}

func ExtractPathParams(path string) string {
	segments := strings.Split(strings.Trim(path, "/"), "/")

	for i, segment := range segments {
		if segment == "lbc" && i+1 < len(segments) {
			return segments[i+1]
		}
	}

	return ""
}

// PrintPretty below function for dev testing purpose
func PrintPretty(resp interface{}) {
	jsonData, _ := json.MarshalIndent(resp, "", "	") //nolint:errchkjson
	fmt.Println(string(jsonData))
}
