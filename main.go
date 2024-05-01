package main

import (
	"context"
	"github.com/glieske/kube-nanny/internal/core/env"
	"github.com/glieske/kube-nanny/internal/core/initLoader"
	"github.com/glieske/kube-nanny/internal/core/preInitLoader"
	"github.com/glieske/kube-nanny/internal/dto"
	"github.com/glieske/kube-nanny/internal/namespaceWatcher"
	"github.com/glieske/kube-nanny/internal/provider"
	log "github.com/sirupsen/logrus"
	"net/http"
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

	go namespaceWatcher.Watcher(ctx, kcs, runningConfig)

	runHealthCheckSrv()
}

func parseConfig() *dto.MainConfig {
	cfg := &dto.MainConfig{}
	logLevel := env.GetEnv("LOG_LEVEL", "info")
	cfg.LoggingLevel = &logLevel

	return cfg
}

func runHealthCheckSrv() {
	log.Debug("Starting health check server...")
	http.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, e := w.Write([]byte("ok")); e != nil {
			log.Fatal(e)
		}
	}))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
