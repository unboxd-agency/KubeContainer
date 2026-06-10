# Alignment Assessment: Platform Confirmation Machinery vs. SWE-bench

**Subject:** our system's evidence machinery — the golden compatibility
corpus, the kind-based real-world e2e gate, and the eval doctrine of
`docs/AGENT-PLATFORM.md` — assessed against **SWE-bench**
(swebench.com): the de-facto standard benchmark for AI agents performing
real software engineering, scored by fail-to-pass resolution on real
GitHub issues.

**Method & honesty note:** swebench.com refuses automated retrieval
(HTTP 403); this assessment is grounded in the published SWE-bench
methodology (real-repo task instances, fail-to-pass test gating,
resolution-rate scoring, Lite/Verified/Multimodal variants), not a live
crawl. Scores are argued against the rubric below, not machine-measured.
Same rubric as `XDM-ALIGNMENT.md` for comparability.

## Rubric

| Metric | Question | Scale |
|---|---|---|
| **Conformance** | How much of SWE-bench's eval methodology does our confirmation machinery already practice? | 0–100, higher better |
| **Convergence** | Are the two approaches evolving toward the same discipline? | 0–100, higher better |
| **Drift** | Expected divergence over time without active alignment | 0–100, *lower* better |
| **Distance** | Current methodological gap | 0.0–1.0 |

## Scores

| Area | Conformance | Notes |
|---|---|---|
| Real-world task grounding | 17/20 | SWE-bench uses real repos and real issues; our gates use real manifests (golden corpus) and a real cluster with real pods and traffic — same refusal of synthetic rehearsal as final proof |
| Binary world-test gating | 16/20 | Their fail-to-pass = our Ready=True + HTTP 200: in both, the world's own assertion must flip, and partial credit does not exist |
| Reproducible harness | 17/20 | Pinned toolchain, containerized runs, deterministic make targets — re-runnable by any party, which is SWE-bench's core demand of submissions |
| Agent-quality measurement | 8/20 | The decisive gap: SWE-bench scores *agents* over a task corpus (resolution rate); our gates score *the artifact*, not the agent that produced it — we have the eval doctrine on paper (lexicon, confirmation pillar) but no agent-eval harness, no task corpus, no resolution metric |
| Scale & statistical power | 6/20 | SWE-bench: hundreds to thousands of instances, leaderboard-grade statistics; ours: 3 corpus manifests + 1 e2e scenario — gates, not yet a benchmark |
| **Conformance total** | **64/100** | |

| Metric | Score | Reading |
|---|---|---|
| **Convergence** | **80/100** | Strong: both disciplines are converging on "real tasks, world-owned pass criteria, continuous public evidence, standings you can lose" — our charter mandates exactly this (engineering concern #6, the intelligence pillar); SWE-bench is the proof it works at ecosystem scale |
| **Drift** | **25/100** (low–moderate) | The doctrine won't drift — it's constitutional here; the risk is *practice* drift: our evals staying artifact-gates while the industry standardizes agent-resolution benchmarks, leaving our "intelligence confirmed by evals" pillar asserted but unmeasured |
| **Distance** | **0.42** | Larger than the XDM gap (0.31): we share the philosophy completely and the machinery partially — the entire agent-scoring layer (task corpus, harness, resolution rate, leaderboard) exists in our documents and not in our code |

## Composite

> **Alignment score: 64 conformant / 80 converging / 25 drifting / 0.42 distant**
> — same gating philosophy, working artifact-gates, but the agent-eval layer
> that is SWE-bench's essence remains doctrine without harness on our side.

## What would close the gap

1. **A task corpus** — era-stamped, real change-requests against this repo
   (issue text + repo snapshot + the world-test that must flip), the
   golden-corpus pattern applied to *work* instead of manifests.
2. **A resolution metric** — % of corpus tasks an agent resolves with CI
   green, measured per agent/model/config: our own resolution rate.
3. **A harness target** — `make eval` standing beside `make test`: run
   agent, apply patch, execute world-tests, emit scored evidence to the
   record.
4. **Standings over time** — scores recorded per the axiom (revisions,
   projections), making "certification you can lose" operational for the
   agents that work on this codebase.
