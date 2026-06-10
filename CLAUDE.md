# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

KubeContainer is a Kubernetes operator (Go, Kubebuilder/controller-runtime) that manages
the lifecycle of containerized workloads through a single `KubeContainer` CRD
(`kubecontainer.unboxd.cloud/v1alpha1`). The operator reconciles each CR into owned
Deployment, Service/Ingress, and HPA resources. The full architecture, CRD schema,
reconcile-loop design, and roadmap live in `docs/DESIGN.md` — read it before making
changes.

## Commands

- `make build` — generate manifests/deepcopy, fmt, vet, and compile the manager.
- `make test` — run unit/integration tests under envtest (downloads control-plane
  binaries to `bin/k8s/` on first run). Note: in some environments the
  `-coverprofile` step errors with `no such tool "covdata"` on packages without
  tests; the test results above that error are still valid. To run without
  coverage: `KUBEBUILDER_ASSETS="$(bin/setup-envtest use 1.35.0 -p path)" go test ./...`
- Single test: add `FIt`/`FDescribe` (Ginkgo focus), or
  `KUBEBUILDER_ASSETS=... go test ./internal/controller/ -v --ginkgo.focus="<It name>"`
- `make manifests generate` — regenerate CRDs and deepcopy after editing
  `api/v1alpha1/kubecontainer_types.go`. Always run before committing type changes.
- `make lint` / `make lint-fix` — golangci-lint.
- `make test-e2e` — kind-based e2e suite (requires a running Docker daemon).

## Layout & conventions

Standard Kubebuilder v4 layout: types in `api/v1alpha1/`, reconciler in
`internal/controller/kubecontainer_controller.go`, Kustomize manifests in `config/`.

- The reconciler is level-triggered and idempotent; children are managed with
  `controllerutil.CreateOrUpdate` and cleaned up via owner references (no finalizers).
- When `spec.scaling.autoscale` is set, the HPA owns the Deployment's replica
  count — the reconciler must not write `spec.replicas` (see the HPA test).
- Spec invariants (replicas/autoscale exclusivity, Ingress host requirement) are
  enforced with CEL `XValidation` markers on the types, not webhooks.
- envtest runs no Deployment controller or garbage collector: tests assert
  `Ready=False/Progressing=True` and delete children explicitly in `AfterEach`.
