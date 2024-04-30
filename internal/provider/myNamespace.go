package provider

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func GetMyNamespace() string {
	if ns, ok := os.LookupEnv("POD_NAMESPACE"); ok {
		return ns
	}

	if data, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}

	log.Warn("Namespace not found, using default namespace: kube-nanny")
	return "kube-nanny"
}
