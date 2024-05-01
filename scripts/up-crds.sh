#!/bin/zsh

for file in ./pkg/crds/*.yaml; do
  kubectl apply -f $file
done
