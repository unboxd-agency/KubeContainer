# Changelog

All notable changes to this project are documented here. The format
follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/); the
project adheres to semantic versioning once released. Per the axiom, this
file summarizes — the git history remains the record.

## [Unreleased]

### Changed

- Organization renamed: unboxd-cloud → **unboxd-agency** — cloud
  described where the work ran; agency describes who the work serves.
  Public references (site, README, NOTICE, manifests) aligned. The Go
  module path is unchanged for now (GitHub redirects); the API group
  `kubecontainer.unboxd.cloud` is unchanged by design — it is the DNS
  domain, and the compatibility corpus protects it.

### Added

- Product website (`site/`, GitHub Pages) with the v0.1.0 install
  command and release links.
- NOTICE: copyright claim, the record as proof, trademark reservation
  (unboxd, KubeContainer, the Fabric, Kube), and the naming rationale.
- Registry-defines-the-contract doctrine (eval/README).
- Kube product specification (`docs/KUBE-SPEC.md`); the box doctrine and
  agent equation on the front door; declared-release mechanism
  (`release/REQUEST`).

## [0.1.0] — 2026-06-10

### Added — operator

- `KubeContainer` CRD (`kubecontainer.unboxd.cloud/v1alpha1`): image,
  port, env, resources, scaling (fixed replicas or autoscale), expose
  (ClusterIP / LoadBalancer / Ingress), HTTP health checks.
- Reconciler managing owned Deployment, Service, Ingress, and HPA via
  `CreateOrUpdate`, with orphan cleanup, HPA-owned replica counts,
  `Ready`/`Progressing`/`Degraded` conditions, events, and a computed
  endpoint in status.
- CEL validation: replicas/autoscale exclusivity; Ingress host
  requirement. Printer columns for image, availability, readiness,
  endpoint.
- envtest integration suite (9 specs) and a kind-based e2e suite with a
  real-workload gate: a declared workload must converge to `Ready=True`
  and serve HTTP 200 from inside the cluster.
- Backward-compatibility golden corpus
  (`internal/controller/testdata/compat/`, frozen, append-only) with
  tests proving era manifests remain valid and convergent forever.
- Install bundle `dist/install.yaml`; README quickstart; Apache 2.0
  LICENSE.

### Added — platform & governance

- `docs/FOUNDING-PRINCIPLES.md`: the charter — dedication, purpose,
  twenty-four principles, the five-clause axiom, the promise, epilogue.
- `docs/AGENT-PLATFORM.md`: the agent ladder, platform capabilities,
  the agent lexicon (130+ terms), reality & drift vocabulary, the seven
  anti-drift protocols, constitutional interpretation rules.
- `docs/DESIGN.md`: architecture, OPA compliance and OpenFGA alignment,
  vendor-neutral distribution & supply-chain policy, roadmap.
- `docs/SOLID-STATE-DATABASE.md`: first product brief (five candidate
  names). `docs/GO-TO-MARKET.md`: enterprise motion — sell trust, not
  lock-in.
- Alignment assessments with scored rubric (conformance / convergence /
  drift / distance): Adobe XDM (78/86/22/0.31) and SWE-bench
  (64/80/25/0.42) under `docs/assessments/`.
- Evaluation registry (`eval/`): append-only task corpus with
  world-tests and provenance (LLM usage schema included), harness, and
  `make eval` evidence reports.
- Vocabulary system: generated index (`docs/VOCABULARY.md`,
  `make vocab`) and enforcement (`make vocab-check`) — bold is coinage,
  coinage requires definition.
- Scheduled drift-audit workflow (weekly): tests, lint, vocabulary
  check, evaluation registry, published evidence artifact.

### Fixed

- Migrated from deprecated `GetEventRecorderFor` to the supported
  events API (staticcheck SA1019).
- e2e CI: kind install pinned to v0.32.0 and sourced from GitHub
  releases (kind.sigs.k8s.io was unreachable from runners; `latest`
  violated the pinning policy).
- Generated eval reports kept out of git; custody is CI artifacts.

### Infrastructure

- Go pinned to an exact patch release in `go.mod` (single source of
  truth); matching pinned `golang` image tags in Dockerfile and
  devcontainer.
- CI: lint, envtest, and kind e2e workflows on every push and PR.
