package main

import (
	"flag"
	"github.com/glieske/kube-nanny/internal/core/initLoader"
	"github.com/glieske/kube-nanny/internal/core/preInitLoader"
	"github.com/glieske/kube-nanny/internal/dto"
	log "github.com/sirupsen/logrus"
)

func main() {
	preInitLoader.ExecutePreInit()
	log.Info("Kube Nanny initializing...")
	cfg := parseConfig()
	initLoader.ExecuteInit(cfg)
	log.Info("Kube Nanny initialized!")

}

func parseConfig() *dto.MainConfig {
	cfg := &dto.MainConfig{}
	cfg.LoggingLevel = flag.String("v", "info", "Log level: debug, info, warn, error, fatal, panic")

	return cfg
}
