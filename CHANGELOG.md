# Changelog

All notable changes to this project are documented here. The format
follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/); the
project adheres to semantic versioning once released. Per the axiom, this
file summarizes — the git history remains the record.

## [Unreleased]

## [0.2.5] - 2026-06-12
### Added
- `spec.storage`: the kube declares what it keeps — size and path,
  reconciled into an owned PersistentVolumeClaim and mount. The claim
  is environment-neutral (the cluster's storage class supplies the
  backing: cloud disk in cloud, local path on metal). Dropping the
  clause unmounts but never deletes the claim — data is never blown.
- Compat corpus gains the storage era (`v0_2_5_storage.yaml`).

## [0.2.4] - 2026-06-12
### Added
- RecordGraph: the record extracted as a graph (146 nodes, internal
  references gated — a broken reference fails the build; 133 external
  anchors indexed), emitted as plain triples and schema.org JSON-LD.
- SchemaKeeper: the schema kept as a tool — pinned context, pinned
  types, no dangling nodes, host blueprints validated against the
  skeleton's declared taxonomies.
- The host skeleton (host-declaration/v1) and two blueprints:
  leap-micro (the verdict host, fixed release) and microos (rolling,
  dev boxes only); the Leap Micro decision record (`deploy/LEAPMICRO.md`).
- The desk, defined in the lexicon (the question rule applied).
### Fixed
- deploy.sh fails loudly when the latest-release tag cannot be
  resolved, and prepares transactional hosts (Leap Micro/MicroOS).

## [0.2.3] - 2026-06-12
### Added
- The site as an image: ghcr.io/unboxd-cloud/kubecontainer-site.
- deploy.sh self-resolves the newest release — one URL deploys on any
  VPS forever; proven end-to-end on bare metal (READY True in under
  two minutes).
### Fixed
- `.dockerignore` re-includes `site/` so the site image has a build
  context (the v0.2.3 first attempt died here).

## [0.2.2] - 2026-06-12
### Fixed
- HomeSetup field failures from the first VPS deploy: apt-lock race
  and the gitlab-runsvdir wedge.

## [0.2.1] - 2026-06-12
### Added
- HomeSetup binary (`-clean-slate`) shipped as a release asset — one
  download converges a VPS (eviction, GitLab, k0s, operator, kube).
- Release now ships test-report.json, security-report.json, SHA256SUMS.
### Fixed
- HomeSetup field-hardened: waits for apt locks, enables
  gitlab-runsvdir before reconfigure (the wedge seen on first deploy).

## [0.2.0] - 2026-06-12
### Added
- The Personal doctrine, the agent contract and flow, the registry
  (skeleton + blueprints), the tools (CodeCompiler, StructuredInstructions).
- The instruction manual (contracts, licenses, FAQ, how-to-use).
- The Primitives, the Doctrine Map, the deploy stack and references.
- release/PROMISE and the promise check: "promise delivered" is now a
  verdict the pipeline renders, never a flourish.
### Changed
- License of record: KubeContainer Research and Community Source License v1.0.


### Changed

- Organization renamed: unboxd-cloud → **unboxd-agency**. The reason,
  in the founder's words: *because this is the agent — unboxd had the
  agency*. "Cloud" was the wrong representation — it described where
  the work ran; "agency" describes who the work serves and what a box
  becomes when identity and governance are added to the artifact
  (artifact + identity + governance = agent). The platform was never
  about the cloud underneath; it was always about the agency conferred.
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
