package logging

import (
	"github.com/bluest-eel/cli/components/config"
	"github.com/geomyidia/zylog/logger"
	log "github.com/sirupsen/logrus"
)

// Setup ...
func Setup(cfg *config.Config) {
	logger.SetupLogging(cfg.Logging)
}

// Load pretends that the global is more functional in nature ...
func Load(cfg *config.Config) *log.Logger {
	Setup(cfg)
	return log.StandardLogger()
}
