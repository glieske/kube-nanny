package namespaceWatcher

import (
	"context"
	"github.com/glieske/kube-nanny/internal/dto"
	log "github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

func Watcher(ctx context.Context, kcs *kubernetes.Clientset, rcfg *dto.RunningConfig) {
	log.Info("Namespace Watcher initialized!")
	defer func() {
		log.Fatal("Namespace Watcher stopped!")
	}()
	for {
		log.Debug("Watcher loop starts...")
		namespaces := getNamespaces(kcs)
		for _, ns := range namespaces.Items {
			log.Debugf("Watching namespace: %s", ns.Name)

		}
		log.Debug("Watcher loop ends...")
		time.Sleep(time.Duration(rcfg.WatcherConfig.LoopDelay) * time.Second)
	}
}

func getNamespaces(kcs *kubernetes.Clientset) *v1.NamespaceList {
	log.Debug("Getting namespaces...")
	namespaces, err := kcs.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{
		LabelSelector: "nanny.letscode.cloud/watch=true",
	})
	if err != nil {
		log.Fatal(err)
	}
	return namespaces
}
