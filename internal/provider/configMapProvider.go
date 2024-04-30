package provider

import (
	"context"
	"github.com/glieske/kube-nanny/internal/dto"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetRunningConfigFromConfigMap(ctx context.Context, kcs *kubernetes.Clientset) *dto.RunningConfig {
	log.Debug("Getting running config from ConfigMap...")
	ns := GetMyNamespace()
	cm, err := kcs.CoreV1().ConfigMaps(ns).Get(ctx, "kube-nanny", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return parseConfigMapData(cm)
}

func parseConfigMapData(cm *v1.ConfigMap) *dto.RunningConfig {
	log.Debug("Parsing ConfigMap data...")
	cfg := dto.RunningConfig{}

	return &cfg
}
