package main

import (
	"flag"
	"github.com/glieske/kube-nanny/internal/dto"
	log "github.com/sirupsen/logrus"
)

func main() {
	onInitLoadLoggingConfiguration()
	log.Info("Kube Nanny initializing...")
	cfg := parseConfig()
	loadLoggingConfiguration(cfg)
}

func onInitLoadLoggingConfiguration() {
	log.SetFormatter(&log.JSONFormatter{})
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

func parseConfig() *dto.MainConfig {
	cfg := &dto.MainConfig{}
	cfg.LoggingLevel = flag.String("v", "info", "Log level: debug, info, warn, error, fatal, panic")

	return cfg
}
