package main

import (
	"context"
	"flag"
	"github.com/glieske/kube-nanny/internal/core/initLoader"
	"github.com/glieske/kube-nanny/internal/core/preInitLoader"
	"github.com/glieske/kube-nanny/internal/dto"
	"github.com/glieske/kube-nanny/internal/namespaceWatcher"
	"github.com/glieske/kube-nanny/internal/provider"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.TODO()
	preInitLoader.ExecutePreInit()
	log.Info("Kube Nanny initializing...")
	cfg := parseConfig()
	initLoader.ExecuteInit(cfg)
	log.Info("Kube Nanny initialized!")

	kcs := provider.GetKubernetesClientSet()
	runningConfig := provider.GetRunningConfigFromConfigMap(ctx, kcs)

	namespaceWatcher.Watcher(ctx, kcs, runningConfig)

	select {}
}

func parseConfig() *dto.MainConfig {
	cfg := &dto.MainConfig{}
	cfg.LoggingLevel = flag.String("v", "info", "Log level: debug, info, warn, error, fatal, panic")

	return cfg
}
