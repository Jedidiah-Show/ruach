package workspace

import "log/slog"

type Service struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
