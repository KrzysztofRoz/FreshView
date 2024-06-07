package handler

import "go.uber.org/zap"

type Handler struct {
	logger *zap.Logger
}

func NewHandler() *Handler {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &Handler{
		logger: logger,
	}
}
