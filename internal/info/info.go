package info

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

var (
	serviceName     string
	namespace       string
	version         string
	gitlabProjectID string
	buildDate       string
	gitLog          string
	gitHash         string
	gitBranch       string
	goVersion       string

	info = Info{
		ServiceName:     serviceName,
		Namespace:       namespace,
		Version:         version,
		GitlabProjectID: gitlabProjectID,
		BuildDate:       getBuildDate(),
		GitLog:          getGitLog(),
		GitHash:         gitHash,
		GitBranch:       gitBranch,
		GoVersion:       getGoVersion(),
	}

	marshalOnce sync.Once
	jsonInfo    []byte
)

type Info struct {
	ServiceName     string
	Namespace       string
	Version         string
	GitlabProjectID string
	BuildDate       time.Time
	GitLog          string
	GitHash         string
	GitBranch       string
	GoVersion       string
}

func GetInfo() Info {
	return info
}

func GetJSONInfo() []byte {
	marshalOnce.Do(func() {
		b, err := json.Marshal(info)
		if err != nil {
			// TODO log error
			return
		}

		jsonInfo = b
	})

	return jsonInfo
}

func getBuildDate() time.Time {
	t, err := time.Parse(time.RFC3339, buildDate)
	if err != nil {
		fmt.Println(err) // TODO: wrong date format log
		return time.Time{}.UTC()
	}

	return t.UTC()
}

func getGoVersion() string {
	if len(goVersion) > 0 {
		return goVersion
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goVersion = buildInfo.GoVersion
		return goVersion
	}

	fmt.Println("can't fill golang version") // TODO: can't fill golang version log

	return ""
}

func getGitLog() string {
	if len(gitLog) == 0 {
		return ""
	}

	b, err := base64.StdEncoding.DecodeString(gitLog)
	if err != nil {
		fmt.Println(err) // TODO: git log must be base64 log
		return ""
	}

	return string(b)
}
