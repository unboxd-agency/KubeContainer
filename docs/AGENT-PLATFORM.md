# The Agent Ladder: From Definition to Mission

A definitional ladder for agent systems, distilled from first principles. Each
rung builds on the one below it; the ladder ends in a mission statement for the
platform that makes trustworthy agents an economy. The KubeContainer operator
in this repository serves as the running worked example: Kubernetes controllers
are agent engineering avant la lettre, and most of the ladder's ideas were
proven there first.

## 1. Agent (the standard definition)

> An **agent** is a program that acts autonomously on behalf of a principal,
> deployed because the work must happen **where the principal isn't** (a remote
> host, a cluster, a user's machine) or must continue **when the principal is
> gone** (background loops, autonomous goal pursuit). Otherwise it would just
> be a function call.

Supporting vocabulary:

- **Actor** — the unit of concurrent computation: private state, asynchronous
  message passing, no shared memory. Reactive plumbing; works for no one.
  Agents are often *implemented* with actors, but an agent has a goal and a
  principal, and may act unprompted.
- **Contract** — mutual obligations at an interface: preconditions,
  postconditions, invariants. (Here: the CRD schema and its CEL rules.)
- **Assertion** — one executable claim that a condition holds at runtime; the
  enforcement instrument for contracts. (Here: the envtest suite.)

In one sentence: actors carry the messages, agents pursue the goals, contracts
define what everyone owes each other, assertions catch the moment anyone breaks
the deal.

## 2. Agent Architecture

The internal structure that lets a program act autonomously. Every architecture
must answer four questions: how the agent **perceives**, how it **decides**,
how it **acts**, and what it **remembers**.

Families: reactive (condition→action), deliberative (world model + planning,
e.g. BDI), hybrid/layered, the modern LLM stack (model + tools + loop + memory
+ guardrails), and the **operator/controller architecture** used here:
perception via watches and an informer cache, decisions in a level-triggered
reconcile loop, idempotent API mutations as actions, and memory externalized
into the environment itself (spec and status). *Stateless agent, stateful
environment* is the defining trick: the agent can crash and resume with
nothing lost.

## 3. Agent Engineering

The discipline of building agents whose autonomy is safe, reliable, and
verifiable. Six concerns:

1. **Goal & contract design** — specifying intent precisely enough to pursue
   and to grade.
2. **Tool/actuator design** — small, idempotent, well-described actions.
3. **Loop & orchestration** — retry, stop, and escalate-to-human decisions.
4. **Memory & context engineering** — what the agent carries vs. what lives in
   the environment.
5. **Guardrails & permissions** — autonomy granted per-action, never wholesale.
6. **Evaluation & observability** — evals for non-deterministic behavior,
   traces for every decision.

## 4. Agent Governance

The layer above engineering: which agents may exist, what each may do, on whose
authority, and who answers when one misbehaves. It is the systematic answer to
the principal-agent problem. Six functions: **identity**, **authorization**,
**policy**, **audit**, **lifecycle**, **accountability** — an agent can be
responsible for an action but never accountable for it.

Kubernetes is the most complete agent-governance system in production use:
ServiceAccounts, RBAC (see `config/rbac/role.yaml`), admission control, audit
logs, and lifecycle control, all per-controller.

## 5. Agent Excellence

The standard of practice the layers below aim at: autonomy trusted by default,
not managed as a risk. Six marks:

1. **Convergence, not just correctness** — reaches the intended end-state, even
   from partial failure.
2. **Calibrated judgment** — acts inside its competence, asks at the edge,
   refuses beyond.
3. **Spirit over letter** — pursues what the principal meant; never games the
   metric.
4. **Economy** — least privilege exercised, fewest actions, lowest cost.
5. **Legibility** — predictable before, observable during, reconstructable
   after.
6. **Graceful degradation** — fails informatively and reversibly.

## 6. The Platform

The **agent control plane**: the operational substrate where agent qualities
stop being claims and become continuously **confirmed** evidence.

| Pillar | Confirmed by |
|---|---|
| Engineering | CI gates: contract tests, regression evals, conformance suites |
| Intelligence | Benchmarks and task evals against golden sets, re-run on every change |
| Reliability | SLOs and error budgets: convergence rate, intervention rate, MTTR |
| Governance | Identity, policy-as-code, immutable audit trails |
| Cost efficiency | Metering per agent per outcome, with budgets and alerts |
| Excellence | The composite scorecard, sustained over time — certification you can lose |

A runtime runs agents; a platform can show you, at any moment, that its agents
deserve to be running.

## 7. The Agent Economy

Confirmation is what makes the economy possible: markets run on trust between
strangers, and the platform's evidence machinery is that trust. The value
chain it must support, end to end:

1. **Create** — SDKs, frameworks, contract/tool standards (MCP).
2. **Distribute** — registries and marketplaces with provenance and signing
   (the OCI-registry pattern applied to agents).
3. **Transact** — identity strong enough to contract on, delegation chains,
   pricing, machine-speed payment rails.
4. **Operate** — the control plane itself: run, govern, evaluate, meter.
5. **Consume** — humans and other agents hiring agents via declared intent.
6. **Settle & account** — metering becomes billing, audit becomes dispute
   resolution, the scorecard becomes reputation.

## 8. Delivery

- **Just-in-time** — the right intelligence routed to the work the instant the
  work exists; no idle inventory. (Autoscaling is JIT compute; this is JIT
  cognition.)
- **Trusted frontier intelligence, from all providers** — provider-neutral
  routing across the moving frontier; the platform's normalized evals are what
  make heterogeneous providers comparable, and JIT becomes arbitrage.
- **Best-in-class tools, from certified suppliers** — the actuator side.
  Verification is a technical fact about an artifact (signatures, provenance,
  SBOMs); **certification** is an institutional standing about a party —
  audited, renewable, revocable (the Certified Kubernetes precedent). The full
  trust stack: certified supplier → verified artifact → confirmed runtime
  behavior.
- **Through an open partner network** — membership gated by the certification
  bar and nothing else; partners deliver as principals in their own right; the
  platform's product is the trust fabric and routing, not the inventory. Open
  networks scale with the ecosystem (Kubernetes vs. its single-vendor
  predecessors is the existence proof).

## The Mission Statement

> **The platform is the control plane where agent engineering, intelligence,
> reliability, governance, cost-efficiency, and excellence are continuously
> converted from claims into confirmed evidence — supporting the agent economy
> across the entire value chain, from creation and distribution through
> transaction, operation, consumption, and settlement — by ensuring
> just-in-time delivery of trusted frontier intelligence from all providers,
> and best-in-class tools from certified suppliers, through an open partner
> network delivery model.**

What it is, whom it serves, how it wins, who delivers it.
