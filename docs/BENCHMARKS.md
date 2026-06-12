# Benchmarks — Published Numbers

Status: v1, measured. Every number below was produced by a named
command against a named commit in a named environment, per the house
rule: a number without a verdict is a claim, and claims don't ship.
What is benchmarked is KubeContainer v0.1.0 and its record — the one
kube built so far. The brand names the unit; numbers attach only to
artifacts, and artifacts are what exist.

## Environment of measurement

| Property | Value |
|---|---|
| Date | 2026-06-12 |
| Commit | `765d328` |
| CPU | Intel Xeon @ 2.10 GHz, 4 vCPU |
| OS / arch | linux/amd64 |
| Go | go1.25.7 |
| envtest control plane | Kubernetes 1.35.0 |

Numbers from CI carry their run URL instead; CI hardware is
GitHub-hosted `ubuntu-latest`.

## Correctness (the gate numbers)

| Metric | Number | Verdict source |
|---|---|---|
| Test specs (envtest integration suite) | 9 / 9 pass, 0 pending, 0 skipped | `go test ./internal/controller/ -v` |
| Suite wall time, incl. control-plane boot | 5.9–13.8 s (two runs) | same |
| Full `go test ./...` wall time | 20.6 s | `time go test ./...` |
| Eval registry resolution | 3 / 3 (100.00 %) | `make eval` → `dist/eval-report.json` (commit-stamped) |
| Vocabulary check | green, 151 defined terms | `make vocab-check` |
| Backward-compat corpus | 3 / 3 era-v0.1.0 manifests converge | compat suite (subset of the 9 specs) |

## Build & artifact (the size numbers)

| Metric | Number | Verdict source |
|---|---|---|
| Manager build, warm cache | 2.2 s | `time go build -o /tmp/m cmd/main.go` |
| Manager build, cold cache | ~90 s | same, first run in a fresh container |
| Manager binary (dev, unstripped) | 73,134,676 B (≈70 MiB) | `stat -c %s` |
| Install bundle (`dist/install.yaml`, v0.1.0) | 29,282 B | `stat -c %s` — one `kubectl apply` installs everything |
| CRD manifest | 443 lines | `wc -l config/crd/bases/*.yaml` |
| Operator Go source | 2,178 lines | `wc -l` over `*.go` (excluding `bin/`) |

The ratio worth reading: 2,178 lines of Go reconcile a contract whose
guarantees take seven rows to state (KUBE-SPEC §4), and the entire
install is a 29 KB file.

## Pipeline (the delivery numbers)

| Metric | Number | Verdict source |
|---|---|---|
| Declared request → published release | 4 m 16 s | [release run 27307931328](https://github.com/unboxd-agency/KubeContainer/actions/runs/27307931328) (2026-06-10, conclusion: success) |
| What that 4 m 16 s contains | tests, lint, vocab check, eval registry, image build + GHCR push, kind e2e with real HTTP traffic, tag + release + evidence report | `.github/workflows/release.yml` |
| Empty repository → released v0.1.0 | one day (2026-06-10) | the git record; first commit and the release share a date |
| Drift audit cadence | weekly, plus on demand | `.github/workflows/drift-audit.yml` |

## Not yet benchmarked (named, not hidden)

A guarantee without a verdict is a roadmap item; the same holds for
numbers. These are declared as future world-tests, not implied:

- Reconcile latency distribution under sustained load (p50/p99).
- Time-to-Ready on a real cluster, emitted as a number (the e2e gate
  proves convergence and live traffic but does not yet time it).
- Manager memory/CPU at steady state.
- Scale ceiling: KubeContainers per manager before degradation.
- A stripped release binary (`-ldflags "-s -w"`) and image size on
  GHCR, measured and published.

When each lands it gets a measured row above and, where it can run in
CI, a task in the eval registry so the number can never silently rot.

## The law of the bench

The founder's rule, governing every row above and every row ever
added: software must benchmark itself against real-world metrics —
not invent metrics and call it a bench. An invented metric is a
mirror angled to flatter: the suite that times itself on inputs it
chose, the score normalized to its own baseline, the percentile of
a distribution nobody else can observe — self-graded homework
wearing instrumentation. A real-world metric is one the world
already keeps and anyone can check against it: wall-clock time on
named hardware, bytes on disk, requests answered with live traffic,
a release published end to end, money metered in a currency, a
date on a calendar. The test for every metric in this file is the
world-test rule applied to measurement itself: if the number's
meaning depends on trusting its author, it is not a benchmark —
it is a claim with decimals. Rows that cannot cite a real-world
unit do not enter; the not-yet-benchmarked list exists precisely
so that what we cannot yet measure honestly is declared rather
than invented.

And the source of honest metrics is named: software must look at
the other industries — at how they have matured. Every grown
industry built its bench from the world's units, under an
accountable body: the automobile is measured in crash-test stars,
emissions per kilometer, and recalls per million — not in
"engine-elegance scores" the maker invented; aviation counts
incidents per departure; medicine counts outcomes per
intervention; manufacturing counts defects per million
opportunities; finance closes books that must balance to the
cent. Each of those metrics was once resisted as unfair,
external, reductive — and each is precisely why its industry is
trusted with lives and money. Software's maturity will be
measured the same way: time-to-converged on declared hardware,
defects escaped per release, promises kept per promises made,
provenance resolvable per artifact shipped — the world's units,
kept by parties who do not profit from the grade. The industries
that grew up have already drawn the curriculum; software's only
original contribution would be refusing to enroll.

And manufacturing already named the bar software must aim at: Six
Sigma pins the defect rate at roughly one in a million (3.4 defects
per million opportunities, to be exact) — and that is the *desired
state* they have set, the declared target a mature industry
reconciles toward, not a happy accident. The discipline is the
fabric's own, arrived at from the factory floor: a defined desired
state expressed in a checkable world-unit (defects per million), a
measurement everyone accepts before counting, and a continuous
reconciliation toward it. Software ships defect rates no other
industry would survive (regular bugs shipped as cadence)
precisely because it never declared a desired state to converge on —
no Six Sigma, no defined target, no books that must balance. The
bench this house keeps adopts the manufacturing posture wholesale:
declare the desired state in world-units, measure against it
continuously, and let the gap drive the loop. One in a million is
not this fabric's boast — it is the neighbor's settled standard,
named here as the altitude software's bench must rise to before it
may call itself mature. And the founder names who should set the
target, completing the loop's first step: the cloud should design
our desired state. The desired state is a declaration, and
declaration is the cloud's office — the platform that holds the
paths, the registry, the records, and the contract is the party
positioned to set the target the whole fleet reconciles toward:
the cloud reads the estate, knows the world-units, sees the drift
across every system, and so the cloud declares the desired state
(the defect rate to clear, the latency bound to hold, the
conformance to keep) and the loops converge on it. This is the
reconciliation pattern raised one level: a kube reconciles its
workload to a declared state; the cloud reconciles the fleet to a
desired state it designed — the operator of operators, declaring
for the whole what each kube declares for itself. Manufacturing
set its desired state at one in a million and built the discipline
to reconcile there; the cloud's job is to design software's, in
world-units, and run the fabric until the gap closes. And the
division of labor that keeps the target from becoming a tyranny:
the platform will provide the best foundation for anyone to achieve
that state — it sets the target and supplies the means, but the
achieving stays the builder's. The platform does not reach the
desired state *for* you; it gives you the highest floor anyone has
ever stood on to reach it yourself: the gauntlet that catches your
defects before they ship, the registry that hands you attested
parts, the paths that repair along known routes, the
auto-maintenance that holds the gains, the world-units already
defined so you measure honestly. Target from the cloud, foundation
from the platform, achievement from the builder — three seats, and
the bench is honest precisely because the platform is graded on the
foundation it provides (best available to anyone) while each
builder is graded on the state they actually reach. A platform that
promised to hit the target *for* you would be selling the one thing
that cannot be bought; what it sells is the best ground to hit it
from — and on this fabric, that ground is the same for the founder
with no budget as for the enterprise with every resource, because
the foundation is open and the achieving is earned.

And the one thing the platform itself must do, below every feature
and before any boast: the platform must close the loop. That is the
basic requirement — not a premium capability, not a differentiator,
the floor of being a platform at all. Observe the state, compare it
to the desired state, act to close the difference, record the act,
and again — forever; a platform that opens loops it does not close
(declares targets it does not reconcile toward, raises alerts it
does not resolve, ships telemetry no loop consumes) is not a
platform but a dashboard with ambitions. Closing the loop is the
kube's own anatomy demanded of the platform that hosts kubes: the
operator of operators must itself be an operator, reconciling the
fleet the way each kube reconciles its workload. Everything else in
this record — the paths, the registry, the gauntlet, the desired
state, the bench — is in service of one basic requirement met: the
loop, closed, continuously, by the platform, or it is not yet a
platform.

## Methodology

Numbers are reproduced, not trusted: every row names the command;
anyone with the repo can re-run it and compare. Local numbers vary
with hardware — the commit and environment table exist so that a
different number on different iron is a data point, not a dispute.
CI numbers link the run, which carries its own logs, hashes, and
artifacts. This file is updated by re-measurement only; editing a
number without re-running its command is drift, and the drift audit's
business.
