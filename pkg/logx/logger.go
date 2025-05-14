package logx

import (
	"context"
	"go.uber.org/zap"
	"os"
	"todo-service/pkg/utils"
)

var logger *zap.Logger

func Init() {
	logger, _ = zap.NewProduction()
	if os.Getenv("APP_ENV") == "dev" {
		logger, _ = zap.NewDevelopment()
	}
	defer func() { _ = logger.Sync() }()
}

func L() *zap.Logger {
	return logger
}

func Info(ctx context.Context, message string) {
	sugar := L().Sugar()
	sugar.Infow(
		message,
		keysAndValues(ctx)...,
	)
}

func Debug(ctx context.Context, message string) {
	sugar := L().Sugar()
	sugar.Infow(
		message,
		keysAndValues(ctx)...,
	)
}

func keysAndValues(ctx context.Context) []interface{} {
	return []interface{}{
		"userId", utils.OrElseEmpty(ctx.Value("userId")),
		"requestId", utils.OrElseEmpty(ctx.Value("requestId")),
	}
}
