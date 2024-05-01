#!/bin/zsh

#KO_DOCKER_REPO=ghcr.io/glieske/kube-nanny ko build -P --bare .
KO_DOCKER_REPO=ko.local ko build -P --bare ./
