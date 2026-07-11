package bootstrap

import (
	"github.com/Jedidiah-Show/ruach/internal/platform/config"
	"github.com/Jedidiah-Show/ruach/internal/platform/engine"
	"github.com/Jedidiah-Show/ruach/internal/platform/logging"
)

func Build() (*engine.Engine, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	logger := logging.New(cfg.Logging)

	return engine.New(cfg, logger), nil
}
