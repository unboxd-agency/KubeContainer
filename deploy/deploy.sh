#!/bin/sh
# Deploy KubeContainer on any VPS — one command:
#   curl -sfL https://raw.githubusercontent.com/unboxd-cloud/KubeContainer/main/deploy/deploy.sh | sudo sh
# Installs k3s (native runtime), the latest released operator, and the
# arithmetic kube. Self-resolves the newest release; never goes stale.
set -e
TAG=$(curl -sfL https://api.github.com/repos/unboxd-cloud/KubeContainer/releases/latest | grep '"tag_name"' | cut -d'"' -f4)
if [ -z "$TAG" ]; then
  echo "FAIL: could not resolve the latest release tag from the GitHub API — re-run in a minute" >&2
  exit 1
fi
echo "deploying release: $TAG"
# Transactional hosts (openSUSE Leap Micro / MicroOS): the root filesystem is
# read-only, so the SELinux policy must be layered in before k3s can install.
if command -v transactional-update >/dev/null 2>&1; then
  if ! rpm -q k3s-selinux >/dev/null 2>&1; then
    echo "transactional host detected. Run these two first, then re-run this same command:" >&2
    echo "  transactional-update pkg install k3s-selinux" >&2
    echo "  reboot" >&2
    exit 1
  fi
  export INSTALL_K3S_SKIP_SELINUX_RPM=true
fi
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="--disable traefik" sh -
sleep 25
k3s kubectl apply -f "https://github.com/unboxd-cloud/KubeContainer/releases/download/$TAG/install.yaml"
sleep 10
k3s kubectl apply -f https://raw.githubusercontent.com/unboxd-cloud/KubeContainer/main/deploy/arithmetic-kube.yaml
sleep 20
echo "================ VERDICT ================"
k3s kubectl get nodes,kubecontainers
