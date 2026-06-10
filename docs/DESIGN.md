# KubeContainer — Design Document

KubeContainer is a Kubernetes operator that provides a single high-level abstraction —
the `KubeContainer` custom resource — for running a containerized workload. A user
declares *what* they want to run (image, ports, scaling, exposure); the operator
materializes and continuously manages the underlying Kubernetes primitives
(Deployment, Service, HorizontalPodAutoscaler, Ingress).

## Goals

- One small, opinionated CRD that covers the common case of "run this container,
  keep it healthy, scale it, and expose it".
- Full lifecycle automation: create, update (rollout), scale, expose, and clean up
  child resources when the CR is deleted.
- Self-healing: drift in any managed child resource is reverted on the next reconcile.
- Clear, observable state via status conditions and Kubernetes events.

## Non-Goals

- Replacing Helm or general-purpose templating — KubeContainer is deliberately opinionated.
- Managing stateful workloads (StatefulSets, PVC orchestration) in v1alpha1.
- Multi-cluster scheduling.

## Tech Stack

| Concern | Choice |
|---|---|
| Language | Go |
| Framework | Kubebuilder (controller-runtime) |
| API group | `kubecontainer.unboxd.cloud` |
| Initial version | `v1alpha1` |
| Testing | envtest (integration), kind (e2e) |
| Packaging | Multi-stage Docker image + Kustomize manifests under `config/` |

## Custom Resource

```yaml
apiVersion: kubecontainer.unboxd.cloud/v1alpha1
kind: KubeContainer
metadata:
  name: my-app
spec:
  image: ghcr.io/acme/my-app:1.4.2
  port: 8080                      # container port; also the Service target port
  env:                            # optional, standard corev1.EnvVar list
    - name: LOG_LEVEL
      value: info
  resources:                      # optional, standard corev1.ResourceRequirements
    requests: { cpu: 100m, memory: 128Mi }
  scaling:
    replicas: 2                   # fixed replica count, OR:
    autoscale:                    # mutually exclusive with replicas
      minReplicas: 2
      maxReplicas: 10
      targetCPUUtilization: 75
  expose:
    type: ClusterIP               # ClusterIP | LoadBalancer | Ingress
    host: my-app.example.com      # required when type=Ingress
  healthCheck:                    # optional; wired to liveness+readiness probes
    path: /healthz
status:
  observedGeneration: 3
  availableReplicas: 2
  endpoint: my-app.default.svc.cluster.local:8080
  conditions:
    - type: Ready                 # Ready | Progressing | Degraded
      status: "True"
      reason: DeploymentAvailable
```

Validation (CEL rules + webhook later if needed):
- `replicas` and `autoscale` are mutually exclusive.
- `expose.host` is required iff `expose.type == Ingress`.

## Architecture

```mermaid
flowchart LR
    U[Developer / GitOps] -->|apply CR| API[API Server + etcd]
    API -->|watch events| OP[KubeContainer Operator\ncontroller-runtime manager]
    OP -->|status update| API
    OP -->|owns| D[Deployment]
    OP -->|owns| S[Service / Ingress]
    OP -->|owns| H[HPA]
    D --> P[Pods]
    H -.scales.-> D
```

### Controller

A single reconciler registered with a controller-runtime manager:

- `For(&KubeContainer{})` — reconciles on CR changes.
- `Owns(Deployment)`, `Owns(Service)`, `Owns(Ingress)`, `Owns(HPA)` — child changes
  (including external drift or deletion) re-enqueue the parent CR.

### Reconcile loop

1. Fetch the `KubeContainer`; if gone, exit (children are garbage-collected via owner references).
2. Build the desired child objects from `spec`.
3. `CreateOrUpdate` each child, setting the CR as controller owner reference.
   - If `scaling.autoscale` is unset, delete any orphaned HPA (and vice versa: the
     HPA owns the replica count, so the Deployment's `replicas` field is left unmanaged).
   - Same orphan-cleanup applies when `expose.type` changes (e.g. Ingress → ClusterIP).
4. Compute status: mirror Deployment availability into `Ready`/`Progressing`/`Degraded`
   conditions, set `availableReplicas`, `endpoint`, and `observedGeneration`.
5. Patch status (server-side apply) only when it changed, to avoid hot-looping.

The loop is idempotent and level-triggered — every pass converges actual state toward
`spec` regardless of which event woke it. No finalizer is needed in v1alpha1 because
all children live in the same namespace and are cleaned up by owner references.

### Failure handling & observability

- Transient errors return an error from `Reconcile` so controller-runtime retries
  with exponential backoff.
- Terminal misconfigurations (e.g. invalid image ref rejected by the API) set
  `Degraded=True` with a reason instead of retrying forever.
- Emit Kubernetes events on child create/update and on degradation.
- Standard controller-runtime Prometheus metrics, plus `/healthz` and `/readyz` probes
  on the manager.

## Project Layout (Kubebuilder standard)

```
api/v1alpha1/          # KubeContainer types, deepcopy, CEL markers
internal/controller/   # reconciler + child-object builders
config/                # CRD, RBAC, manager manifests (Kustomize)
test/e2e/              # kind-based end-to-end tests
```

## Roadmap

1. **v1alpha1** — scaffold, CRD, reconciler for Deployment + Service, status conditions, envtest coverage.
2. **v1alpha2** — HPA + Ingress support, CEL validation, e2e suite.
3. **v1beta1** — conversion webhook, defaulting webhook, image update policy (auto-rollout on digest change).
