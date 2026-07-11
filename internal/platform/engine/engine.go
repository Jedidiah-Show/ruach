package engine

import (
	"log/slog"

	"github.com/Jedidiah-Show/ruach/internal/platform/config"
)

type Engine struct {
	config *config.Config
	logger *slog.Logger
}

func New(cfg *config.Config, logger *slog.Logger) *Engine {
	return &Engine{
		config: cfg,
		logger: logger,
	}
}
