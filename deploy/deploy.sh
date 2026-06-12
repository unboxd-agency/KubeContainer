#!/bin/sh
# One command: curl -sfL <this file> | sudo sh
set -e
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--disable traefik" sh -
sleep 25
k3s kubectl apply -f https://github.com/unboxd-cloud/KubeContainer/releases/download/v0.2.1/install.yaml
sleep 10
k3s kubectl apply -f https://raw.githubusercontent.com/unboxd-cloud/KubeContainer/main/deploy/arithmetic-kube.yaml
sleep 20
echo "================ VERDICT ================"
k3s kubectl get nodes,kubecontainers
