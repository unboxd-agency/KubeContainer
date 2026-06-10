# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

KubeContainer is a Kubernetes operator (Go, Kubebuilder/controller-runtime) that manages
the lifecycle of containerized workloads through a single `KubeContainer` CRD
(`kubecontainer.unboxd.cloud/v1alpha1`). The operator reconciles each CR into owned
Deployment, Service/Ingress, and HPA resources. The full architecture, CRD schema,
reconcile-loop design, and roadmap live in `docs/DESIGN.md` — read it before making
changes.

## Repository Status

No source code has been scaffolded yet — only the design doc exists. The first
implementation step is the Kubebuilder scaffold described in the roadmap
(`docs/DESIGN.md`). Once code lands, update this file with build/test/lint commands
(`make build`, `make test`, envtest usage) and any deviations from the standard
Kubebuilder layout.
