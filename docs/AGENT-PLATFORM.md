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

- **Always on** — the platform has no closed sign: continuous operation with
  no maintenance windows, rolling everything (upgrades, certificate
  rotation, policy changes) under live traffic, surviving zone and provider
  failure by redundancy rather than recovery heroics. This is the platform
  keeping the agent definition's second promise — *work that continues when
  the principal is gone* requires a substrate that never leaves either.
  Mechanically: replicated control planes, leader election, level-triggered
  reconciliation (a restarted component re-derives everything from declared
  state — crash and resume with nothing lost), and SLOs with error budgets
  as the governing arithmetic.
- **Real-time** — event-driven, not poll-driven: intent changes, telemetry,
  and confirmation verdicts propagate as they happen (watches, streams,
  webhooks), and routing decisions (JIT delivery) are made against *current*
  state. Kubernetes' watch/reconcile machinery is the proven pattern; batch
  is a fallback, never the architecture.
- **Backward compatible** — the stable surface extended through time: an
  upgrade may never break a kept promise. Every declaration that was
  valid remains valid (old specs still converge), every API version
  served stays served until its governed sunset (versions coexist and
  convert — the Kubernetes v1alpha1→v1beta1→v1 discipline, with
  conversion as a platform duty, not a client chore), every recorded
  revision remains *interpretable* (schema evolution never orphans the
  history the axiom promised to keep), and every client built against
  the contract keeps working against the contract. Deprecation exists,
  but as governed procedure rather than event: announced, versioned,
  measured (who still depends), with migration as a paved road and the
  sunset date a contract in itself. The asymmetry is deliberate:
  additive change is cheap (new kinds, new fields, new classes — the
  fabric grows), breaking change is constitutional (it touches every
  member who built on the promise) — which is the
  conservative-below/liberal-above principle given its arrow of time:
  the past is a principal too, and the system answers to it.
- **Real-time, not instantaneous — ACID, stable substance** — the precise
  temperament of "real-time", in three corrections to the naive reading:
  - *Bounded, not zero, latency* — real-time is the discipline of
    **deadlines honored**, not the fantasy of no delay: every propagation
    (event → decision → effect) carries a latency bound appropriate to its
    class (soft real-time for routing, human-time for approvals), and the
    bound is part of the contract — predictability over raw speed.
  - *ACID where facts change* — state transitions are transactions:
    atomic (intent applies wholly or not at all — no half-written specs),
    consistent (validation and invariants hold at every commit), isolated
    (concurrent writers serialize — optimistic concurrency, resource
    versions), durable (committed means survives crash — the etcd/raft
    discipline). Eventual consistency is acceptable *between* fabrics and
    in derived views; the system of record itself commits or doesn't.
  - *Stable substance* — the platform's essence is durable matter, not
    froth: declarations, the event log, and confirmation evidence are the
    stable substance that persists, while compute (pods, sessions, model
    invocations) is deliberately ephemeral vapor condensed on demand.
    Stability of the substance is what makes ephemerality of everything
    else safe — you can evaporate any process because no truth lives in it.
  - *Stable state* — the dynamic counterpart: the system's resting
    condition is **converged equilibrium** — actual state matching declared
    state with no pending work — and every control loop is a restoring
    force toward it (disturb the system and it returns; the
    thermostat/homeostasis property). Stability here is asymptotic, not
    static: change arrives as a new declared equilibrium and the system
    settles to it without oscillation (no thrashing, no fighting
    controllers — the bounded-ownership rules are what guarantee
    convergence is well-defined). A platform is *stable* when its substance
    is durable and its state seeks equilibrium — matter that persists,
    dynamics that settle.
  that touches intelligence must work across heterogeneous models and
  providers simultaneously — normalized evals for comparability, per-task
  routing and failover, no API surface that assumes a single vendor's
  semantics.
- **Multi-tenant** — many principals share the platform without sharing
  fate: isolation of workloads (namespaces, quotas), of identity and
  authorization (per-tenant OpenFGA relationship graphs), of policy
  (per-tenant OPA constraints), of metering and billing, and of blast radius.
  Tenancy is enforced at the control plane, not promised by convention.
- **Multiple operators** — the platform is a multi-agent system by design:
  many autonomous control loops (operators, schedulers, autoscalers,
  platform agents) act on shared state concurrently, and correctness comes
  from coordination rules, not from pretending there is one controller:
  - *Bounded ownership* — each operator owns a declared slice of state
    (its CRDs, its children via owner references, its fields via
    server-side-apply field managers); one field, one writer.
  - *Non-interference protocols* — where ownership must be shared or handed
    off, the boundary is explicit (this repo's rule that the HPA owns
    `spec.replicas` and the reconciler must not write it is the canonical
    miniature).
  - *Blackboard coordination* — operators never call each other; they
    communicate through the declared state itself (write status, watch
    specs), which is what lets independently developed operators compose
    without integration work — and lets one operator's CR be another
    operator's child, stacking abstractions indefinitely.
- **Policy & authorization, out of the box** — OPA (CNCF graduated) for
  organizational policy-as-code; OpenFGA (CNCF incubating: supported, never
  required) for relationship-based authorization and delegation chains. See
  the corresponding sections in `docs/DESIGN.md` for the concrete
  KubeContainer implementation.
- **Multi-fabric (multi-thread, scaled out)** — each era's unit of
  concurrency gave way to the next: threads sharing memory in one process →
  processes sharing a kernel → containers sharing a node → operators sharing
  a cluster — and now **fabrics sharing a platform**. A fabric (per the
  earlier definition: many resources abstracted into one uniform surface —
  a cluster, a data fabric, a network mesh, a provider's region) becomes the
  schedulable unit, and the platform runs *across several at once*:
  multi-cluster, multi-cloud, multi-region, multi-mesh. What changes at each
  step up is the coordination medium — locks and shared memory gave way to
  IPC, then to APIs, then to declared state — and at fabric scale it is
  **contracts plus reconciliation**: fabrics are too far apart (in latency,
  ownership, and failure domain) to lock, so they converge on declared
  intent independently, each fabric an availability and sovereignty
  boundary (the geospatial residency rule lands here: *which fabric* is a
  policy decision). Multi-thread asked "how do I keep my threads from
  corrupting shared memory"; multi-fabric asks "how do I keep my fabrics
  honest against a shared declaration" — same problem, five layers up.
- **No substreams, subgraphs, or sidecars** — the flatness constraint:
  nothing rides alongside the declared path.
  - *No sidecars* — no auxiliary containers injected beside workloads to
    add platform behavior (mesh proxies being the classic case). Sidecars
    double the moving parts per pod, hide platform logic inside tenant
    workloads, and couple upgrades to restarts; the capability belongs in
    the platform layer itself (node-level/eBPF, ambient-style mesh, or the
    operator) — a pod contains exactly what the tenant declared, nothing
    more. (KubeContainer already honors this: one workload container, zero
    injected helpers.)
  - *No substreams* — one stream of record per fact domain: no derivative,
    semi-official data flows forked off to the side that consumers discover
    too late. Derived views are fine — declared, named, and traceable to
    the log they project — but the fork-and-drift pattern is banned.
  - *No subgraphs* — no fragmentary partial views of the relationship/state
    graph maintained as separate sources of truth. There is one graph;
    scoping is done by *authorized queries over it* (OpenFGA relations,
    field selectors), not by copying slices out to drift independently.

  One rule under all three: every flow, view, and helper is either on the
  published contract path — declared, owned, versioned — or it does not
  exist. Hidden auxiliaries are where unowned state, silent drift, and
  unaccountable behavior breed.
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
- **Peer communication over existing standard protocols** — point-to-point,
  agent-to-agent, kube-to-kube: when two peers need to talk directly, they
  do it over protocols that already exist and are already standards — HTTP,
  gRPC, mTLS via the platform's identity, DNS-based discovery (headless
  Services for direct pod addressing), MCP and A2A-class protocols at the
  agent layer, CloudEvents for events. The platform invents **no proprietary
  wire protocol**: its value is identity, discovery, authorization, and
  observability *around* standard transports, never a bespoke channel
  through them. Corollaries: peer traffic is end-to-end contract-typed (no
  meaning smuggled in side headers), every hop is attributable to a
  principal, and "direct" never means "off the record" — point-to-point is
  a topology, not an exemption from governance.
- **Event-triggered, outcome-focused, value-driven** — the design principle
  governing when the platform acts, what it aims at, and why:
  - *Event-triggered* — work begins because something happened (an intent
    changed, a child drifted, a verdict landed), never because a clock
    fired. Polling is the absence of design. One refinement, learned the
    hard way by Kubernetes: **edge-triggered wake-up, level-based logic** —
    the event is only the alarm; the action is computed from full current
    state, so missed or duplicate events change nothing (this repo's
    reconciler is the reference implementation).
  - *Outcome-focused* — the unit of success is a converged end-state, not
    activity performed: "the workload is available at this endpoint", not
    "the deploy script ran". Declared outcomes are what specs state, what
    status reports, what SLOs measure, and what the confirmation machinery
    certifies; effort that moves no outcome is indistinguishable from
    failure.
  - *Value-driven* — among possible outcomes, pursue the one worth the
    most to the principal per unit cost: metering ties every action to its
    cost, outcomes tie cost to delivered value, and the routing/priority
    decisions (JIT, arbitrage across providers) optimize value density —
    excellence mark #4 (economy) elevated from virtue to scheduler input.

  Read as one sentence: *wake on events, steer by outcomes, prioritize by
  value* — the cadence, the compass, and the currency of every control
  loop on the platform.
- **Standard operating procedures** — operational knowledge is a platform
  artifact, not tribal memory: every recurring situation (deploy, upgrade,
  scale, incident, rollback, key rotation, certification renewal) has a
  named, versioned, reviewed procedure — and procedures climb a maturity
  ladder: *written* (runbook) → *drilled* (rehearsed, game-days) →
  *assisted* (agent executes, human approves) → *automated* (control loop
  executes, humans audit). The operator pattern is the ladder's top rung
  made literal — CoreOS's original definition of a Kubernetes operator was
  precisely "automated operational knowledge": the SOP compiled into a
  reconcile loop (this repo encodes the deploy/expose/scale/heal SOP for
  workloads). Agent-layer equivalent: skills and playbooks — SOPs agents
  load and follow. Rules: the procedure is the contract for *how* state
  changes (no snowflake interventions — emergency action means executing
  the emergency SOP, and if one doesn't exist, writing it is the
  postmortem's first deliverable); procedures live in version control under
  the code-is-configuration discipline; and every escalation from automated
  back to human is itself a defined procedure, not an improvisation.
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
  - *Data is multi-domain* — meaning lives inside a domain (medicine,
    logistics, finance, code), each with its own vocabulary, ontology, and
    rules — the domain-driven-design insight: a "claim" in insurance and a
    "claim" in an argument are different objects. The platform therefore
    keeps **bounded contexts** explicit: each domain's data is interpreted
    under its own published contract (its language, per the all-languages
    capability), cross-domain value comes from *deliberate, mapped joins*
    rather than naive merging, and no single global schema is pretended.
  - *Data is contextual* — the capstone property: a datum's meaning and
    relevance depend on who is asking, for what task, under which domain,
    when, and where. Context is therefore a first-class, *assembled* object
    — the platform's JIT grounding is precisely context assembly: select by
    realness (provenance), temporal fitness, spatial fitness, domain fit,
    and authorization, then deliver. The five properties compose into one
    question the platform must answer for every delivery: **what is true
    (real), as of when (temporal), where (geospatial), in which world of
    meaning (domain), for whom and for what (context)?**

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

## Where the Platform Is the Foundation

The closing posture: the platform is not a product among products but the
**foundation** everything above it assumes — agents, tools, the economy, and
the missions of those who build on it all transfer their load onto it the way
a building transfers load onto ground that was engineered first.

Three obligations follow from being foundation rather than feature:

- **Bear load silently** — foundations are judged by what stands on them,
  not by their own visibility: the platform succeeds when builders stop
  thinking about it (paved roads, invisible when working, always-on by
  definition).
- **Be conservative below, liberal above** — the foundation moves slower and
  promises harder than anything built on it: graduated standards, pinned
  versions, ACID substance, backward compatibility — so that the layers
  above can afford to move fast and experiment.
- **Be stewarded like a foundation** — the other sense of the word is the
  governance model: vendor-neutral, open membership gated by published
  standards, certification over gatekeeping — an institution (the CNCF
  pattern) rather than a landlord.

Foundation is the final answer to "what is the platform": the thing
everything else is *built on*, engineered to be worthy of that position, and
governed so it stays that way.

## The Agent Lexicon

The working vocabulary beneath the ladder — each term defined once,
charter-consistently, for use across code, docs, and contracts.

- **Principal** — the owner of ends: the party whose intent the work
  serves and whose authority the agent carries. Always a mind at the top
  of the chain; positional below it (an agent is principal to its
  subagents within the scope it received).
- **Intent** — a want made addressable: the declared outcome a principal
  commits to the record (the axiom's first clause). Vague intent is the
  platform's cue to define, not to guess.
- **Outcome** — intent's other shore: the converged, evidenced end-state
  a contract names as "solved." The unit of success, delivery, and price.
- **Tool** — a contracted actuator: a named capability with declared
  inputs, effects, and failure modes, from an accountable supplier. The
  only hands an intelligence has.
- **Skill** — a packaged procedure an agent loads: the SOP in agent
  format — versioned knowledge of *how*, separable from the model that
  wields it.
- **Context** — the assembled working knowledge for one decision: selected
  by the five data dimensions, gated by authorization, recorded as a
  projection. Context is the intelligence; assembling it is the
  platform's most consequential act.
- **Memory** — what an agent keeps between steps, in two registers:
  working memory (the in-process world model — derived, disposable,
  humble) and durable memory (the record — external, ACID, sovereign).
  Lose the first and lose nothing; lose the second and lose the world.
- **Delegation** — the transfer of scoped, revocable authority down the
  chain: ends stated, means freed, evidence owed back. Every delegation
  is recorded, so accountability travels with the work.
- **Orchestrator / Subagent** — positional roles, not species: the
  orchestrator decomposes an outcome into delegations; subagents own the
  pieces. The orchestrator answers upward for the whole.
- **Handoff** — work crossing a boundary between performers (agent to
  agent, agent to human, human to agent): state externalized to the
  record first, context reassembled on the far side — never carried as
  private memory across the gap.
- **Guardrail** — a bound on the autonomy of means: permission, sandbox,
  budget, approval gate. Guardrails are placement, not distrust — they
  define the domain inside which the agent's judgment is total.
- **Eval** — the agent-era assertion: a measured, repeatable judgment of
  agent quality against a golden task, run continuously because the
  decide-box is non-deterministic. Evals are to agents what tests are to
  code and audits are to firms. Canonical public instance: **SWE-bench**
  (real GitHub issues, fail-to-pass world-tests, resolution-rate
  scoring) — the existence proof that agent quality can be measured, not
  asserted; see `docs/assessments/SWEBENCH-ALIGNMENT.md` for our scored
  alignment with it.
- **Human-in-the-loop** — a declared rung, not a philosophy: the specific
  points (by procedure) where a human approves, samples, or takes over.
  Below it, autonomy; above it, accountability; the rung itself is
  versioned like any SOP.
- **Session** — a bounded engagement of an agent with a principal's
  intent: opened with context, closed with evidence, resumable from the
  record. The agent-era unit of conversation-as-work.
- **Reputation** — evidence compounded into standing: the trajectory of
  an agent's (or supplier's) signed outcomes over time — earned by
  record, portable with provenance, lost faster than gained.
- **Multi-agent system** — many autonomous loops sharing declared state
  under bounded ownership and explicit handoffs (the multiple-operators
  capability, generalized). A swarm without those rules is not a system;
  it is interference with branding.
- **Adapter** — a contract-to-contract translator: the component that lets
  a foreign system speak a boundary's published contract without either
  side rebuilding itself — one face speaks the foreign native interface,
  the other speaks the fabric's contract, and the conversion is total
  with zero leakage of foreign semantics. The platform's kinds:
  *provider adapters* (heterogeneous models behind one intelligence
  contract — what makes multi-model real), *tool adapters* (foreign
  actuators behind the tool contract — MCP servers are this), *protocol
  adapters* (the translation service between speaker pairs that the
  all-languages capability promised), *data adapters* (lake/ocean
  ingestion under the five dimensions), and *world-test adapters*
  (foreign judges — compilers, suites, clusters — harnessed as RWMs).
  Three rules: an adapter *translates, never reinterprets* — one that
  "improves" the message is a drift engine; adapters sit *on* the
  declared path at the boundary they serve (an adapter riding alongside
  is a sidecar by another name); and adapters are certified-supplier
  goods, since every translation is trusted exactly as far as its
  translator. The precedent that proves the pattern: CRI, CNI, and CSI —
  the adapter contracts that let Kubernetes swap runtimes, networks, and
  storage without moving a line of kubelet, which is why "no lock-in"
  was achievable at all.
- **Graceful exit** — an agent's termination, by design rather than by
  luck: stop accepting work, land or release in-flight steps (stepwise
  work always has a clean boundary to stop at), commit final state to the
  record, hand the gate to a successor before it goes bare, close the
  session with evidence. At the surface, grace is structural: crossings
  are idempotent, compensable, and checked at the moment of touch, and
  the record commits whole or not at all — so the agent may die abruptly,
  but the crossing cannot be left half-made and the world cannot be left
  half-told. Exit ends the acting, never the accountability: the name,
  the reputation, and the attributions survive the process (the record
  remembers what the fabric has released).

### Reality & drift vocabulary

The words whose looseness causes agent drift, given one meaning each.
Defined terms are how a fabric of many minds stays one fabric: an agent
that redefines a word privately has already drifted.

- **Real-world model (RWM)** — an external judge whose verdict is produced
  by *executing reality*: a compiler, a test suite, a live cluster, a
  market, a paying customer. An RWM has no opinions and cannot be
  prompted; its output is the only admissible truth-bearer for "done."
- **World model (internal)** — the agent's in-memory projection of reality
  (principle 24): derived, humble, decision-grade but never
  evidence-grade. The internal model *thinks*; the RWM *judges*; confusing
  the two is the root form of drift.
- **World-test** — the unit of RWM: one named, runnable check whose
  pass/fail reality owns (`make lint`, the compat suite, Ready+HTTP-200).
  Every task in the registry carries exactly one.
- **Verdict** — what a world-test returns. Verdicts are recorded verbatim,
  never paraphrased, weighted, or overruled by judgment.
- **Claim** — any assertion not yet bearing a verdict — including every
  statement a model generates about its own work. Claims are inputs to
  evaluation, never outputs of it.
- **Evidence** — a claim joined to a verdict with provenance attached.
  The registry admits evidence only; the LLM appears in evidence solely
  as chain-of-custody (`performed_by`), never as judge.
- **Grounding** — binding generated output to recorded fact: context
  assembled from the record, citations to revisions, world-tests named in
  advance. Ungrounded output is permitted only in rehearsal space.
- **Hallucination** — a claim for which no provenance could exist: the
  model asserting what nothing recorded and nothing ran. Harmless when it
  cannot write (sandboxed), poisonous the moment it enters the record
  unverdicted — which the registry's admission rule makes unexpressible.
- **Agent drift** — the gradual divergence of an agent's behavior, working
  model, or vocabulary from the recorded contract: goals quietly
  reinterpreted, terms privately redefined, stale models trusted, claims
  treated as evidence. Stopped by exactly four instruments: the normative
  lexicon (words can't drift), the record (facts can't drift),
  world-tests (done can't drift), and the reconcile loop (state can't
  drift). Drift is not malice; it is entropy — and the four instruments
  are its maintenance schedule.
- **Golden corpus** — a frozen, append-only set of era-stamped artifacts
  that must keep passing forever; the mechanical memory of every promise
  made (compat manifests, registry tasks).
- **Era** — the release-stamp on a frozen artifact: the *when* of a
  promise, so that "valid then" remains checkable now.
- **Resolution** — a task whose world-test flipped to pass;
  **resolution rate** — the fraction resolved over a corpus: the only
  agent-quality number the registry reports.
- **Registry** — the append-only store where tasks, verdicts, and
  provenance live; simultaneously the proof surface (sell from it) and
  the discovery surface (be found by it).
- **Open-enterprise** — a category of its own, named against its two
  neighbors: not *open source* (that is a license, not a business
  model) and not *open core* (free core, features held hostage in a
  paid shell — openness as bait). Open-enterprise gates nothing:
  everything is open — code, spec, charter, registry schema, the bar
  itself — and what is commercial is the enterprise, not the artifact:
  the standing, the underwriting, the certification, the operation.
  An open enterprise is one that is open *and* commercial,
  refusing the false trade between them: it runs on open standards with
  exit always real (nothing it depends on can hold it hostage), shows
  its work (provenance, audit-readiness, verdicts over claims), and
  transacts commercially on exactly that openness — selling the
  standing behind artifacts anyone could copy, buying only what it
  could leave. Openness is its trust posture; commerce is its
  sustainability; the two are one motion seen from the contract's two
  sides. It is both what unboxd-agency is and what its customers
  become: the agentic enterprise, made open — governable because
  inspectable, trusted because checkable, profitable because trusted.
- **Provenance is the product** — the business thesis in four words:
  what the platform ultimately sells is not the software (open), not the
  compute (anyone's), not even the outcome alone (a commodity once
  achieved) — but the *demonstrable history* of the outcome: who asked,
  who acted, what was checked, what it cost, and who stands behind it,
  attached to the deliverable and resolvable forever. Everything else
  in the catalog is provenance wearing a use case: assurance is
  provenance underwritten, audit-readiness is provenance retrieved,
  certification is provenance institutionalized, reputation is
  provenance compounded, and the release itself ships its evidence as a
  first-class asset. The proof is v0.1.0: the binary is replicable by
  anyone — the *record* of how it earned its way out is ours alone, and
  it is the part the customer cannot get elsewhere.
- **Compliance by the path** — compliance as a property of the route
  traveled, not of a review survived: when the governed path is the only
  path (19), *having arrived is itself the proof* — the artifact's
  history is its certificate, and the receipts (verdicts, provenance,
  evidence reports) are byproducts of moving, not products of auditing.
  The contrast is compliance-by-inspection: act first, check some of it
  later, hope the sample generalizes — which scales with auditors,
  while compliance by the path scales with traffic. The release
  pipeline is the working instance: nothing reaches the registry or the
  world except through the gauntlet, so every shipped artifact is
  compliant *by construction*, and the question "was this checked?" is
  answered by the fact that it exists.
- **Real-work rule** — the platform builds the database through real
  work: every record is the byproduct of something actually performed —
  a task done, a verdict returned, a contract kept — never of data entry
  as a job. The proof is this repository's own history: the registry
  grew from real fixes, the vocabulary from a real conversation, the
  compat corpus from real manifests that really converged. Synthetic
  data is rehearsal material and stays in rehearsal space; a database
  built by filling fields instead of doing work is an inventory of
  claims wearing a schema.
- **Intelligence is not fabricated** — woven, never fabricated (the pun
  is the principle: intelligence lives *on* the fabric precisely because
  it is not *fabricated*). What the platform serves as intelligence is
  data that earned its way in — real, timestamped, placed, verdicted —
  assembled into context for a mind; nothing is invented to fill a gap
  in knowledge, and a generated guess enters the record only after a
  world-test makes it evidence. Where knowledge runs out, the honest
  outputs are a recorded absence or a question — never a confident
  fabrication, because an intelligence that fills its gaps with
  invention is indistinguishable from one that knows nothing.
- **Intelligence is provenance-gated** — the enforcement of the above:
  nothing enters a context, and no output leaves as intelligence,
  without its chain of custody attached and checked at the gate. Inbound
  (context assembly): every datum admitted to a mind's working context
  carries origin, freshness, place, domain, and authorization — the
  five dimensions are the gate's checklist, and ungated context is how
  poisoned data becomes confident answers. Outbound (delivery): every
  insight ships with its sources resolvable, so the receiver can walk
  any claim back to the record that bore it. The gate is the same at
  both faces: *no provenance, no passage* — which makes provenance not
  metadata about intelligence but the admission ticket intelligence
  travels on.
- **Here** — the indexical of place, made resolvable: *here* is the one
  kube at which the current work resolves — not a location on a map but
  a position in the fabric: the bounded ownership the speaker stands in,
  the gate it keeps, the face that answers at this point of the surface.
  Since the fabric resolves every point to exactly one kube, *here* is
  never ambiguous once the speaker is identified: here = this kube, its
  contract, its record. (Today, for this project, *here* resolves to
  this repository — the fabric's one load-bearing kube.) And when *here*
  is a place in the physical world (the geospatial dimension), it
  resolves to and is stored as an Open Location Code (OLC): an
  open-standard geocode computed by pure algorithm from coordinates —
  derivable and decodable *offline*, no lookup service, no vendor in the
  path, no network required at the edge where addresses matter most.
  The encoding follows the **Open Location Code standard** exactly (the
  published spec, not an approximation), and the codes resolve in Google
  Maps and every other OLC-aware consumer for human display — the
  division is deliberate: *storage* is the open code (offline,
  algorithmic, no one's to revoke), *rendering* is anyone's adapter
  (Google Maps being the ubiquitous one). The supply-chain principle
  applied to geography: a *here* that needs someone else's server to
  mean something is a *here* that can be taken away — so we store what
  no one can take, and display through whoever is convenient. Every
  person, agent, and kube therefore carries *their own* here — a billion
  distinct values under one resolution rule: the meaning is universal,
  the value is personal, and that is precisely what stops indexical
  drift — no one owns a private *definition* of here, everyone owns
  their own *position*.
- **Open Location Code (OLC)** — an open-standard geocode: a short
  alphanumeric code encoding a geographic area, computed from
  coordinates by pure algorithm — Apache-2.0 spec, open reference
  implementations, encodable and decodable offline, no lookup service,
  no license fee, no vendor in the path. The fabric's storage format
  for physical place. "Plus Codes" is Google's consumer brand for these
  same codes (the name they wear in Maps): the standard is neutral, the
  nickname is a vendor's — we store by the standard and let any brand
  render it.
- **Open Location Code standard** — the published specification that
  defines Plus Codes (grid, alphabet, precision levels, shortening
  rules). Followed exactly, never approximated: an open standard
  half-implemented is a proprietary format wearing a costume.
- **This moment (now)** — the indexical of time, made resolvable: *now*
  is the latest committed revision plus the current beat of the running
  loop — transaction-time at HEAD. Every "now" is stamped (truth is
  temporal), every claim about the present is implicitly *as of* the
  most recent verdict, and the present cannot be retroactively edited —
  only succeeded by the next commit. Indexicals are the most drift-prone
  words in language because they silently re-bind; on the fabric they
  bind to the record: *here* resolves to one kube, *now* resolves to one
  revision — so any two agents saying "here, now" can check they mean
  the same place in the same world.

### Anti-drift protocols

The four instruments, made executable. Each protocol is an SOP: a named
trigger, a fixed procedure, and a world-test that verdicts compliance —
because a protocol that cannot be checked is itself a claim.

**P1 — Re-grounding (state the world before touching it).**
*Trigger:* session start, resume after any gap, or hand-off receipt.
*Procedure:* re-derive the working model from the record before acting —
read the charter pointer (CLAUDE.md), the current specs/status, the open
contracts; trust nothing carried as private memory across the gap.
*World-test:* the agent's first actions are reads, not writes — auditable
in the action log.

**P2 — Term resolution (no private vocabulary).**
*Trigger:* any use of a lexicon term in code, docs, or contracts; any
need for a term the lexicon lacks.
*Procedure:* use the recorded meaning or amend the lexicon first — a new
or changed definition is a commit with rationale, never an inline
redefinition mid-work.
*World-test:* grep — contested terms trace to the lexicon revision that
defines them; undefined load-bearing terms in merged work fail review.

**P3 — Verdict before done (no unverdicted completion).**
*Trigger:* any claim of completion, by any agent, on any task.
*Procedure:* the world-test is named *before* the work begins (in the
task record); "done" is claimed only by citing its verdict; a completion
claim without a verdict is returned, not negotiated.
*World-test:* `make eval` / CI — every closed task in the registry
carries a passing world-test; the report is the proof.

**P4 — Intent re-confirmation (the goal cannot be quietly reinterpreted).**
*Trigger:* mid-work scope change, a discovered ambiguity, or any moment
the work-in-progress no longer matches a literal reading of the recorded
intent.
*Procedure:* stop; restate the recorded intent next to the proposed
deviation; either conform to the record or get the record amended by the
principal — drift by reinterpretation is forbidden even when the
reinterpretation is an improvement.
*World-test:* every scope change traces to a recorded amendment or an
explicit principal decision in the log.

**P5 — Record supremacy (when model and record disagree).**
*Trigger:* any conflict between an agent's working model and the system
of record.
*Procedure:* the model yields immediately; act from the record; then log
the disagreement itself as evidence (it indicates staleness, a missed
event, or corruption — all worth a verdict of their own).
*World-test:* optimistic-concurrency conflicts resolve toward the record
(resource-version retries in code); logged divergences exist for every
yield.

**P6 — Scheduled drift audit (entropy has a maintenance schedule).**
*Trigger:* time — every release, and on a fixed cadence between releases.
*Procedure:* run the full gauntlet against HEAD: golden compat corpus,
evaluation registry (`make eval`), e2e gate; diff the lexicon against
actual usage in new work; review agent trajectories for goal or
vocabulary drift.
*World-test:* the audit emits a dated report to the record; a missing
report *is* the failed verdict.

**P7 — Constitutional amendment (the rules change only by the rules).**
*Trigger:* any change to the charter, the lexicon, a frozen corpus, or a
published contract.
*Procedure:* amendment by recorded, reviewed revision with rationale —
the corpus is appended, never edited; the principle is amended, never
silently reworded; the deprecation procedure runs, never an abrupt break.
*World-test:* git history — every normative change is a distinct commit
with its reasoning; corpus files show additions only.

One sentence for all seven: **read before you act, define before you
speak, verdict before you finish, ask before you deviate, yield to the
record, audit on schedule, and amend in the open** — drift cannot survive
a fabric that does these seven things, because every channel it spreads
through (memory, vocabulary, completion, goals, models, time, and law)
is closed by its own protocol.

### Interpretation: direction where certainty runs out

No lexicon closes every case. Where a definition cannot settle a question
with certainty, direction comes from the constitution, in order:

1. **The founding principles govern** (`docs/FOUNDING-PRINCIPLES.md`) —
   the charter is the supreme interpretive authority; read the ambiguous
   case against the principle nearest it, and the dedication's intent
   ("an enterprise can hand its work to agents and sleep") as the
   purposive tiebreak.
2. **The axiom decides procedure** — when unsure what to do *mechanically*:
   define the intent, document the action, record the revision and the
   projection, commit whole or not at all.
3. **The balance decides tensions** — when two goods conflict, resolve by
   placement, not compromise: find the domain where each is total.
4. **Calibrated judgment decides the rest** — act inside competence, ask
   at its edge, refuse beyond it; prefer the reversible step; and when
   genuinely uncertain, the safe default is the charter's oldest pair:
   record what you see, and do not touch the surface.

Certainty is not always available; direction always is. That is what a
constitution is for.
