# How to Use — Collated from the Official Documentation

Part of the instruction manual, under the same sourcing law as the
FAQ: collated, never fabricated. Every step below is taken from the
named official documentation of its component — for this house's own
components the official documentation is this repository (the
Makefile, the tool usage lines, the specs), and each step is the
documented command verbatim; for upstream components the official
source is linked and the steps are its own quickstart, cited not
restyled. Nothing here is invented; everything resolves to its
source.

## KubeContainer (official source: this repository)

Install (README / release page):

    kubectl apply -f https://github.com/unboxd-agency/KubeContainer/releases/download/v0.1.0/install.yaml

Declare and converge (docs/KUBE-SPEC.md §5 — the only required
interface):

    kubectl apply -f deploy/arithmetic-kube.yaml
    kubectl get kubecontainers
    kubectl describe kubecontainer arithmetic   # conditions carry reasons

Exit (KUBE-SPEC §3 step 6):

    kubectl delete kubecontainer arithmetic     # children collected by owner refs

## The development gauntlet (official source: Makefile, CLAUDE.md)

    make build        # manifests, deepcopy, fmt, vet, compile
    make test         # envtest suite
    make lint         # golangci-lint
    make vocab        # rebuild the vocabulary index
    make vocab-check  # P2: bold is coinage, coinage requires definition
    make eval         # the evaluation registry -> dist/eval-report.json

## The tools (official source: docs/TOOLS.md, the usage lines in source)

    go build ./cmd/codecompiler && ./codecompiler
        # verdict: does the code compile, and does it conform

    go build ./cmd/structuredinstructions
    ./structuredinstructions <declaration.json>          # validate
    ./structuredinstructions <declaration.json> -submit  # register

## The rehearsal deployment (official source: hack/deployrehearsal)

    export KUBEBUILDER_ASSETS="$PWD/bin/k8s/1.35.0-linux-amd64"
    go run ./hack/deployrehearsal
        # control plane up, loop live, declaration admitted, children converge

## The deployment ladder (official sources, each step theirs)

- **minikube** (https://minikube.sigs.k8s.io/docs/start/):
  `minikube start` — then the KubeContainer install above, unchanged.
- **k3s** (https://docs.k3s.io/quick-start):
  `curl -sfL https://get.k3s.io | sh -` — then the install above;
  the full walk at deploy/VPS.md.
- **OpenStack** (https://docs.openstack.org/install-guide/): the
  estate's walk is the official install guide's, series-pinned at
  deploy time; this house's stack picks and reasons at
  deploy/STACK.md.
- **kubectl itself** (https://kubernetes.io/docs/reference/kubectl/):
  every command above is its official usage; the kube adds no CLI
  of its own by policy.

## The rule of this manual

Each section names its official source in its heading; a step that
cannot be traced to its source does not enter; and when an official
upstream document changes, the collation re-checks against the
living anchor (the references table at deploy/REFERENCES.md) — the
manual follows the documentation it collates, never the other way
around, except for this house's own components, where the repository
is the official documentation and this manual is its index.
