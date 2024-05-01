#!/bin/zsh

img=$(KO_DOCKER_REPO=ko.local ko build -P --bare ./)

helm upgrade --install kube-nanny -n kube-nanny --set image.fullOverride="${img}" --set nanny.logLevel=debug ./pkg/helm-charts/kube-nanny
