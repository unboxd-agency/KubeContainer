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

## Platform Capabilities

Cross-cutting requirements the platform must satisfy, with their governing
standards:

- **Real-time** — event-driven, not poll-driven: intent changes, telemetry,
  and confirmation verdicts propagate as they happen (watches, streams,
  webhooks), and routing decisions (JIT delivery) are made against *current*
  state. Kubernetes' watch/reconcile machinery is the proven pattern; batch
  is a fallback, never the architecture.
- **Multi-model** — "from all providers" made operational: every capability
  that touches intelligence must work across heterogeneous models and
  providers simultaneously — normalized evals for comparability, per-task
  routing and failover, no API surface that assumes a single vendor's
  semantics.
- **Multi-tenant** — many principals share the platform without sharing
  fate: isolation of workloads (namespaces, quotas), of identity and
  authorization (per-tenant OpenFGA relationship graphs), of policy
  (per-tenant OPA constraints), of metering and billing, and of blast radius.
  Tenancy is enforced at the control plane, not promised by convention.
- **Policy & authorization, out of the box** — OPA (CNCF graduated) for
  organizational policy-as-code; OpenFGA (CNCF incubating: supported, never
  required) for relationship-based authorization and delegation chains. See
  the corresponding sections in `docs/DESIGN.md` for the concrete
  KubeContainer implementation.
- **All languages, all speaker pairs** — the platform carries meaning across
  every combination of participants, each with its appropriate language
  class:
  - *Human ↔ Human* — natural languages: full internationalization and
    localization; no privileged human language.
  - *Human ↔ Machine* — intent languages: natural-language goals in, legible
    status out; declarative specs (YAML/CRDs) as the precise dialect.
  - *Machine ↔ Machine* — agent and service protocols: MCP, A2A-class
    protocols, gRPC/REST with typed contracts.
  - *System ↔ System* — interoperability standards: OCI, OpenAPI/JSON
    Schema, CloudEvents — the schemas that let platforms compose.
  - *System ↔ Information* — data languages: query (SQL/GraphQL),
    serialization (JSON/Protobuf), and semantics (RDF-class) layers.

  One rule binds them all: every boundary speaks a *published contract*
  (grammar + schema + semantics), so translation between any two parties is
  a platform service, not a per-integration negotiation. Tooling-wise this
  is the Language Server Protocol insight generalized: define the language
  once, get every editor/consumer for free.
- **Where code is configuration** — the inversion of "configuration as
  code", and the platform's deepest operating principle: generic, certified
  engines execute; *what they do* is entirely declared. Users do not program
  the platform — they parameterize it with versioned, diffable, reviewable
  declarations, and control loops make the declarations true. Consequences:
  every change is a data change (auditable by construction, evaluable by
  OPA, reversible by `git revert`); behavior is inspectable without reading
  source; and the blast radius of change is bounded by schema validation
  before anything executes. Where logic must live in the declaration, it is
  embedded as constrained expression languages (CEL, Rego) — code admitted
  into configuration on configuration's terms: sandboxed, terminating,
  side-effect-free. KubeContainer is the working example: a running,
  scaled, exposed workload is produced by twelve lines of YAML and zero
  lines of user code.

- **Data as intelligence** — the companion inversion to "code is
  configuration": data on this platform is not inert record but the active
  substance intelligence is made of, in both directions:
  - *Data → intelligence (grounding):* models are frontier-capable but
    context-blind; the platform's data — documents, state, telemetry —
    delivered just-in-time (retrieval, context assembly) is what turns
    general capability into situated intelligence. The same JIT discipline
    applies: the right data, at the moment of need, under authorization
    (OpenFGA relations gate what each principal's context may contain).
  - *Exhaust → intelligence (the flywheel):* every confirmation artifact the
    platform produces — eval scores, traces, metering, reliability history,
    reputation — is itself data that drives the platform's own decisions:
    routing, certification, pricing. The platform is its own first
    decision-intelligence customer.

  Governance corollary: if data is intelligence, then data governance *is*
  intelligence governance — provenance, consent, and tenancy isolation on
  data are upstream controls on what any agent can know or leak.

  Two properties make this trustworthy:
  - *Data is real* — data is the platform's anchor to reality: models
    generate claims, data carries evidence. Anything asserted without a
    data lineage behind it is opinion, and the confirmation machinery only
    accepts evidence. Provenance is therefore not metadata garnish; it is
    what makes a datum admissible.
  - *Data is temporal* — every datum carries its time: when the event
    happened, when it was recorded, and how long it remains valid
    (bitemporality: what was true vs. what was known). Intelligence decays —
    yesterday's cluster state answers no question about now — so freshness
    is a first-class quality dimension, JIT delivery includes *temporal*
    fitness, and the system of record is an immutable event log from which
    current state is a derived view (the watch/reconcile pattern again:
    status is always a projection of events, never a hand-edited fact).
  - *Data is geospatial* — every datum also carries its *where*: the
    location it describes, where it was produced, and where it may reside.
    Three obligations follow: location as a queryable dimension (geo-indexed
    retrieval, OGC/GeoJSON-class standards, not free-text place names);
    place-aware delivery (route work to the data and the nearest capacity —
    locality is the spatial half of JIT, latency its measure); and
    residency as governance (sovereignty and jurisdiction are *where*
    constraints — tenancy isolation has a geography, and policy must be
    able to say "this data does not leave this region").

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
