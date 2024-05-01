#!/bin/zsh

for file in ./tests/k8s/managed-ns/*.yaml; do
  kubectl apply -f $file
done
