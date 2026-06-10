# Evaluation Registry

This is the platform's own registry of evaluated work — not a public
leaderboard. Each task in `corpus/` is a frozen provenance record of a
real issue: what was asked, against which commit, resolved by which
commit, proven by which **world-test** (the command whose pass/fail the
real world owns). The directory is append-only: tasks are never edited,
new eras add new tasks.

Run `make eval` to execute every task's world-test at HEAD and emit an
evidence record (`dist/eval-report.json`): commit, timestamp, per-task
results, resolution rate. Reports are provenance — the chain of custody
that lets the work be sold as confirmed, not claimed (see "Assurance as
a policy" and "Platform the provenance" in
`../docs/FOUNDING-PRINCIPLES.md`).

Roadmap: agent-execution mode (run an agent against `base_commit`,
apply its patch, run the world-test — a per-agent resolution rate over
the registry), per the gap analysis in
`../docs/assessments/SWEBENCH-ALIGNMENT.md`.

The registry is also the **discovery surface**: each task doubles as a
capability advertisement with evidence attached — what the system has
demonstrably done, queryable by id, era, and world-test. Buyers, partners,
and agents discover capability here the same way they verify it: by the
record (unique and independently identifiable, per the charter — you can
not sell what cannot be found, nor trust what cannot be checked).

## LLM usage in provenance

Where an LLM performed or assisted the work, the task's `provenance`
block must record the intelligence consumed — required for the cost
pillar (metering per agent per outcome), for reproducibility, and for
multi-model comparability:

```yaml
provenance:
  performed_by: <agent runtime / session id>
  llm:
    model: <provider>/<model>@<version>   # the configured model id
    role: author | assistant | none       # who did the work vs. judged it
    usage: {tokens_in: N, tokens_out: N, cost_usd: X}  # when metered
  evidence: <world-test receipts: CI run ids, report refs>
```

Rules: the LLM appears in provenance as worker, never as judge (verdicts
stay world-owned); `role: none` is valid and meaningful (human-only or
deterministic work); unmetered usage is recorded as `usage: unmetered`
rather than omitted — an absent field is ambiguity, a recorded absence
is a fact. Existing era-frozen tasks are not backfilled (append-only);
the schema binds new tasks from era v0.2.0 onward.
