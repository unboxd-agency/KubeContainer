# The Fabric

An intelligent operating system of work: declared outcomes, kept by
autonomous loops, woven from contracts, with provenance attached to
everything it delivers. The fabric is built from kubes — whole,
indivisible units of kept promises — and this repository contains its
first one, plus the constitution the whole weave answers to.

**KubeContainer** is that first kube: a Kubernetes operator where twelve
lines of YAML become a running, scaled, exposed, self-healing workload —
the reference implementation of the
[Kube product specification](docs/KUBE-SPEC.md), governed by the
[founding principles](docs/FOUNDING-PRINCIPLES.md). You declare what you
want; the kube makes it true, keeps it true, proves it was true, and
answers for it.

And KubeContainer is the box — the plain, graspable thing every
metaphor here reduces to. A declared boundary with work inside it,
whole and indivisible, stackable and packable, opaque about its
internals and contractual at its surface. The shipping container
organized the physical economy by making every cargo the same shape to
every crane; the KubeContainer does the same for work — one schema to
every fabric, one contract to every operator. unboxd, because the value
was never the box: it is what declaring the box sets free.

```yaml
apiVersion: kubecontainer.unboxd.cloud/v1alpha1
kind: KubeContainer
metadata:
  name: my-app
spec:
  image: ghcr.io/acme/my-app:1.4.2
  port: 8080
  scaling:
    autoscale:
      minReplicas: 2
      maxReplicas: 10
  expose:
    type: Ingress
    host: my-app.example.com
  healthCheck:
    path: /healthz
```

## Proof, not promises

Nothing here asks to be believed; everything names its verdict:

| Claim | Check it |
|---|---|
| It ships | [Release v0.1.0](https://github.com/unboxd-cloud/KubeContainer/releases/tag/v0.1.0) — image, install bundle, sha256 digests |
| It works in real clusters | The e2e gate: a declared workload must converge and serve HTTP 200 in a live cluster — every push, in CI |
| It will not break its word | The [golden compatibility corpus](internal/controller/testdata/compat/): era manifests must stay valid forever, CI-enforced |
| It carries its evidence | `eval-report.json` ships as a release asset — the [evaluation registry](eval/README.md)'s verdicts, attached |
| It is governed | The [charter](docs/FOUNDING-PRINCIPLES.md), the [lexicon and protocols](docs/AGENT-PLATFORM.md), and a vocabulary check that has already caught its own authors |

## The document map

| Read | For |
|---|---|
| [KUBE-SPEC.md](docs/KUBE-SPEC.md) | What a kube is — anatomy, guarantees, conformance |
| [DESIGN.md](docs/DESIGN.md) | The operator architecture and roadmap |
| [FOUNDING-PRINCIPLES.md](docs/FOUNDING-PRINCIPLES.md) | The constitution: 24 principles, axiom, promise |
| [AGENT-PLATFORM.md](docs/AGENT-PLATFORM.md) | The agent ladder, lexicon, anti-drift protocols |
| [GO-TO-MARKET.md](docs/GO-TO-MARKET.md) | What is sold and to whom (provenance is the product) |
| [CHANGELOG.md](CHANGELOG.md) | What shipped, when |

## Features

- **One CRD for the common case** — Deployment + Service + optional
  Ingress/HPA from a dozen lines of YAML.
- **Self-healing** — drift in any managed child is reverted by the reconcile
  loop; deleting the CR garbage-collects everything via owner references.
- **Safe scaling semantics** — fixed `replicas` or an `autoscale` block
  (mutually exclusive, enforced by CEL validation); under autoscaling the HPA
  owns the replica count and the operator never fights it.
- **Observable** — `Ready`/`Progressing`/`Degraded` conditions, events on
  every child change, and a computed `endpoint` in status:

  ```
  $ kubectl get kubecontainers
  NAME     IMAGE                     AVAILABLE   READY   ENDPOINT
  my-app   ghcr.io/acme/my-app:1.4.2 2           True    my-app.example.com
  ```

## Quickstart

Prerequisites: a Kubernetes v1.30+ cluster, `kubectl`, and (to build) Go and
Docker.

```sh
# Install CRDs and deploy the operator
make docker-build docker-push IMG=<registry>/kubecontainer:v0.1.0
make deploy IMG=<registry>/kubecontainer:v0.1.0

# Run a workload
kubectl apply -k config/samples/
kubectl get kubecontainers

# Tear down
kubectl delete -k config/samples/
make undeploy
```

Or install from the prebuilt bundle (CRDs, RBAC, and manager in one file):

```sh
kubectl apply -f dist/install.yaml
```

Notes:

- `expose.type: Ingress` requires an ingress controller in the cluster;
  `scaling.autoscale` requires metrics-server. Everything else works on any
  conformant cluster — the operator uses only stable upstream APIs.
- HPA scaling decisions need CPU `resources.requests` on the workload.

## Development

```sh
make build      # generate, fmt, vet, compile
make test       # envtest-based unit/integration suite
make lint       # golangci-lint
make run        # run the manager locally against your kubeconfig
```

See [CLAUDE.md](CLAUDE.md) for conventions and
[docs/DESIGN.md](docs/DESIGN.md) for the architecture. Run `make help` for all
targets.

## License

Apache License 2.0 — see [LICENSE](LICENSE). Copyright 2026.
