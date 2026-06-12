# Deploying the kube on an Ubuntu VPS

The real ground, named by the founder. The rehearsal is green
(hack/deployrehearsal); this is the same declaration walked on real
metal — every step a command, every command checkable. Run as a user
with sudo on Ubuntu 22.04/24.04.

## 1. The substrate — conformant Kubernetes on one VPS

k3s: CNCF-conformant, single-binary, built for exactly this ground.
Pin the channel; the version you install is the version you record.

    curl -sfL https://get.k3s.io | INSTALL_K3S_CHANNEL=v1.35 sh -
    sudo k3s kubectl get nodes   # expect: Ready

## 2. The operator — KubeContainer v0.1.0 (one bundle, one apply)

    sudo k3s kubectl apply -f \
      https://github.com/unboxd-agency/KubeContainer/releases/download/v0.1.0/install.yaml
    sudo k3s kubectl -n kubecontainer-system get pods   # expect: Running

## 3. The declaration — the arithmetic kube

    sudo k3s kubectl apply -f \
      https://raw.githubusercontent.com/unboxd-agency/KubeContainer/main/deploy/arithmetic-kube.yaml
    sudo k3s kubectl get kubecontainers   # watch READY become True

## 4. The verdict — real traffic at the face

    sudo k3s kubectl get kubecontainer arithmetic -o jsonpath='{.status.endpoint}'
    sudo k3s kubectl run probe --image=curlimages/curl:8.10.1 --rm -it --restart=Never \
      -- curl -s -o /dev/null -w '%{http_code}' http://arithmetic.default.svc
    # expect: 200 — the kube serving at its face

## The honest seating

This session cannot reach the VPS (no SSH from the managed
environment — recorded constraint). Two lawful hands can:

- **Yours** — paste the four steps; the whole walk is ~3 minutes.
- **The platform's** — add the VPS as a self-hosted GitHub Actions
  runner and the pipeline becomes the delivery channel: the same
  gauntlet that gates the repo deploys to the metal, headless,
  through the one recorded channel — which is the doctrine's own
  preferred route (delivery at the ports, never by hand at the
  surface).

Exit clause, as always: `k3s-uninstall.sh` removes the substrate
whole; the declaration, the record, and the evidence stay in the
repo — exit real.
