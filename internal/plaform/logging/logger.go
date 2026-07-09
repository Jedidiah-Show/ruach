package logging

import (
	"log/slog"
	"os"

	"github.com/Jedidiah-Show/ruach/internal/platform/config"
)

func New(cfg config.LoggingConfig) *slog.Logger {
	var handler slog.Handler
	if cfg.Format == "json" {
    // JSON handler
	handler = slog.NewJSONHandler(os.Stdout, nil)
} else {
    // Text handler
	handler = slog.NewTextHandler(os.Stdout, nil)
}
	return slog.New(handler)
}
