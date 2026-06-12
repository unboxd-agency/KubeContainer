#!/bin/sh
# Deploy KubeContainer on any VPS — one command:
#   curl -sfL https://raw.githubusercontent.com/unboxd-cloud/KubeContainer/main/deploy/deploy.sh | sudo sh
# Installs k3s (native runtime), the latest released operator, and the
# arithmetic kube. Self-resolves the newest release; never goes stale.
set -e
TAG=$(curl -sfL https://api.github.com/repos/unboxd-cloud/KubeContainer/releases/latest | grep '"tag_name"' | cut -d'"' -f4)
echo "deploying release: $TAG"
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--disable traefik" sh -
sleep 25
k3s kubectl apply -f "https://github.com/unboxd-cloud/KubeContainer/releases/download/$TAG/install.yaml"
sleep 10
k3s kubectl apply -f https://raw.githubusercontent.com/unboxd-cloud/KubeContainer/main/deploy/arithmetic-kube.yaml
sleep 20
echo "================ VERDICT ================"
k3s kubectl get nodes,kubecontainers
