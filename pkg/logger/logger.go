package logger

import (
	"context"
	"go.uber.org/zap"
)

func Info(ctx context.Context, message string) {
	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	userId := ctx.Value("userId")
	if userId == nil {
		userId = ""
	}

	sugar := logger.Sugar()
	sugar.Infow(
		message,
		"userId", userId,
	)
}
