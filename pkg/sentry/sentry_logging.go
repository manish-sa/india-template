package sentry

import (
	"context"
	"fmt"
	"time"

	sentry "github.com/getsentry/sentry-go"
)

const (
	ctxKVKey  = "sentry_kv"
	ctxTagKey = "sentry_tag"
)

func WithKV(ctx context.Context, key string, value any) context.Context {
	if ctx.Value(ctxKVKey) == nil {
		ctx = context.WithValue(ctx, ctxKVKey, make([]any, 0))
	}

	kv, ok := ctx.Value(ctxKVKey).([]any)
	if !ok {
		return ctx
	}

	return context.WithValue(ctx, ctxKVKey, append(kv, key, value))
}

func WithTag(ctx context.Context, key string, value any) context.Context {
	if ctx.Value(ctxTagKey) == nil {
		ctx = context.WithValue(ctx, ctxTagKey, make([]string, 0))
	}

	tags, ok := ctx.Value(ctxTagKey).([]string)
	if !ok {
		return ctx
	}

	return context.WithValue(ctx, ctxTagKey, append(tags, key, fmt.Sprint(value)))
}

func Error(ctx context.Context, msg string) {
	sendEventCtx(ctx, msg, sentry.LevelError)
}

func Warning(ctx context.Context, msg string) {
	sendEventCtx(ctx, msg, sentry.LevelWarning)
}

func sendEventCtx(ctx context.Context, msg string, level sentry.Level) {
	e := sentry.Event{}
	e.Level = level
	e.Message = msg
	e.Threads = []sentry.Thread{
		{
			Stacktrace: getStackTrace(),
			Current:    true,
		},
	}
	e.Tags = make(map[string]string)
	e.Extra = map[string]interface{}{
		"full_msg": msg,
	}

	if tagsCtx := ctx.Value(ctxTagKey); tagsCtx != nil {
		tags, ok := tagsCtx.([]string)
		if ok {
			for i := 0; i < len(tags); i += 2 {
				e.Tags[tags[i]] = tags[i+1]
			}
		}
	}

	if kvCtx := ctx.Value(ctxKVKey); kvCtx != nil {
		kv, ok := kvCtx.([]any)
		if ok {
			for i := 0; i < len(kv); i += 2 {
				e.Extra[fmt.Sprint(kv[i])] = kv[i+1]
			}
		}
	}

	sentry.CaptureEvent(&e)
}

func getStackTrace() *sentry.Stacktrace {
	s := sentry.NewStacktrace()
	s.Frames = s.Frames[:len(s.Frames)-3] // 3 - means exclude package internal calls

	return s
}

func Flush() {
	const flushTimeout = 2 * time.Second
	sentry.Flush(flushTimeout)
}
