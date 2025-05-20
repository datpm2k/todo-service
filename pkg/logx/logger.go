package logx

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"sync"
)

type ctxKey struct{}

type Fields map[string]interface{}

var (
	logger *zap.Logger
	once   sync.Once
)

// Init initializes the logger. Call once in main().
func Init() {
	once.Do(func() {
		var err error
		if os.Getenv("APP_ENV") == "dev" {
			logger, err = zap.NewDevelopment()
		} else {
			logger, err = zap.NewProduction()
		}
		if err != nil {
			panic(fmt.Sprintf("failed to initialize zap logger: %v", err))
		}
	})
}

// L returns the base zap logger instance.
func L() *zap.Logger {
	if logger == nil {
		Init()
	}
	return logger
}

// WithFields adds structured fields to the context.
func WithFields(ctx context.Context, fields Fields) context.Context {
	if fields == nil || len(fields) == 0 {
		return ctx
	}
	existing := getFields(ctx)
	merged := make(Fields, len(existing)+len(fields))
	for k, v := range existing {
		merged[k] = v
	}
	for k, v := range fields {
		merged[k] = v
	}
	return context.WithValue(ctx, ctxKey{}, merged)
}

// getFields retrieves structured log fields from context.
func getFields(ctx context.Context) Fields {
	if ctx == nil {
		return Fields{}
	}
	if val, ok := ctx.Value(ctxKey{}).(Fields); ok {
		return val
	}
	return Fields{}
}

// Helper to convert map to zap.Infow format: []interface{}{key1, val1, key2, val2}
func keysAndValues(ctx context.Context) []interface{} {
	fields := getFields(ctx)
	kv := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		kv = append(kv, k, v)
	}
	return kv
}

// Base internal function to log with level
func log(ctx context.Context, level string, msg string) {
	sugar := L().Sugar()
	args := keysAndValues(ctx)

	switch level {
	case "info":
		sugar.Infow(msg, args...)
	case "debug":
		sugar.Debugw(msg, args...)
	case "warn":
		sugar.Warnw(msg, args...)
	case "error":
		sugar.Errorw(msg, args...)
	}
}

func Info(ctx context.Context, msg string)  { log(ctx, "info", msg) }
func Debug(ctx context.Context, msg string) { log(ctx, "debug", msg) }
func Warn(ctx context.Context, msg string)  { log(ctx, "warn", msg) }
func Error(ctx context.Context, msg string) { log(ctx, "error", msg) }
