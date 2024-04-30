package initLoader

import (
	"github.com/glieske/kube-nanny/internal/dto"
	log "github.com/sirupsen/logrus"
)

type LoadLoggerInitHandler struct {
	next initHandler
}

func (s LoadLoggerInitHandler) execute(cfg *dto.MainConfig) {
	loadLoggingConfiguration(cfg)

	if s.next != nil {
		s.next.execute(cfg)
	}
}

func (s *LoadLoggerInitHandler) setNext(next initHandler) {
	s.next = next
}

func loadLoggingConfiguration(cfg *dto.MainConfig) {
	log.WithField("loggingLevel", *cfg.LoggingLevel).Debug("Loading logging configuration...")
	switch *cfg.LoggingLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.Debug("Logging configuration loaded!")
}
