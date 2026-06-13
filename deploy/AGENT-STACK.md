# The Agent-Engineering Stack — Complete, but Lightweight

The founder's ruling: a complete stack for agent engineering, kept
lightweight. Complete = every concern of agent engineering (the six
of AGENT-PLATFORM.md §3) has a named seat; lightweight = each seat is
the smallest leader that fills it, nothing runs that isn't earning,
the whole thing standing on one VPS. Weight is the enemy of the edge;
completeness is the enemy of the toy. Both held — balance, not
maximization.

| Agent-engineering concern | The seat | The lightweight pick | Weight |
|---|---|---|---|
| Goal & contract design | the CRD + CEL; StructuredInstructions | in-repo, already built | ~0 (a binary, a schema) |
| Tool/actuator design | the operator framework (Kubebuilder) | the kube itself | the operator pod |
| Loop & orchestration | the reconciler; the platform-bound orchestrator | KubeContainer | one controller |
| Memory & context | the record — no additional DB; FabricDB when built | k0s etcd now (the kube carries its state) | none added |
| Guardrails & permissions | k0s RBAC; the agent contract terms 1–10 | k0s built-in | none added |
| Evaluation & observability | the eval registry; CodeCompiler; the gauntlet | in-repo scripts + binaries | run on demand |
| Runtime | k0s (single binary, own containerd) | k0s | ~1 binary |
| Registry | zot | zot (CNCF, minimal) | one pod |
| Source + CI | GitLab CE + runner | GitLab (the one heavy guest) | the one big tenant |
| Ground | OpenTofu-declared VPS | one VPS | the metal |
| Build | Cloud Native Buildpacks | pack | on demand |
| Surface | the browser; the site kube | nginx kube | one pod |

The lightweight discipline, stated so it cannot drift into bloat:

- **No additional database** — the record is the store; k0s etcd holds
  cluster state, the kube carries its own; FabricDB is the only store
  ever added, and only when built.
- **One custodian per function** — one runtime (k0s containerd, Docker
  struck), one registry (zot), one source host (GitLab), one image
  builder (pack). No second of anything.
- **On-demand, not always-on, for the desk tools** — CodeCompiler,
  the eval harness, pack, the gauntlet run when invoked; they are not
  resident daemons eating the box while idle.
- **GitLab is the one heavy resident, and it earns it** — source,
  CI, runner in one. If even that is too heavy for the chosen metal,
  the lawful lighter swap is Forgejo (the candidate already named) —
  recorded so the choice is the founder's, not a silent downgrade.

Completeness check: every concern has a seat, no seat is empty but
FabricDB (declared, awaiting the build order), and nothing on the
list is there for show. Lightweight check: one VPS carries it; the
only always-on residents are k0s, the operator, zot, the site, and
GitLab — five, each earning its memory. Complete and light is a
balance, and this is where it resolves.

## Consider the edge — the stack folds smaller

The founder's law: the same stack must fold to the edge, where weight
is not a preference but a wall (a sensor, a card, a browser tab, a
microcontroller cannot carry GitLab, cannot carry a VPS's worth of
residents). So the stack has two readings, one declaration:

- **The home (estate/VPS)** — the full table above: source, CI,
  registry, runtime, site, all resident, GitLab the heavy guest.
- **The edge** — only the kube at the core: The Metal Kube
  (KubeContainer Core) — loop, record, contract, one signed channel,
  the surface-native runtime, and nothing else aboard. No GitLab, no
  zot, no full control plane on the device: the edge pulls signed
  declarations from the home through the one channel, runs the kube,
  and serves its face. Source, CI, registry, build all live at the
  home; the edge receives the product, never the toolchain.

The seam between them is the headless-delivery doctrine, exact: the
home is where it is built, certified, registered (the cloud side, the
capable guard); the edge is where it is delivered and run (the ports,
the surface). What folds away at the edge is everything that builds;
what remains is everything that keeps. The stack is complete at the
home, minimal at the edge, and one declaration spans both — because
the difference between home and edge is count and resident, not kind:
the kube is the same kube; only its company changes.

The lightweight rule, edge-stated: at the edge, the only resident is
the kube. Everything else is the home's, reached over the one
channel, never installed on the device whose surface must stay
headless. Consider the edge, and the stack does not grow a second
architecture — it sheds, down to the one resident the edge was always
for.


## The AGenNextHub estate

The founder's repositories, each holding one seat; this house records
the map and builds only the non-agent seats (protocol P8: agent
machinery is the principal's act alone):

| Seat | Repository | State on the record |
|---|---|---|
| Law | AGenNextHub/Agent-Constitution | binding; loaded by every session here |
| Boundary contract | AGenNextHub/Agent-Space | spec; implemented as `registry/SKELETON-SPACE.json` |
| Files (Agent-Drive) | this house: FileFabric | built; image ships per release (the space drive, agennext.space) |
| The room | AGenNextHub/Agent-IDE | active code (Theia workbench, orchestrator, governance); no releases yet; agent machinery — the principal's |
| Research office | AGenNextHub/Agent-Converter | active code (deep-agent research, confidence-tagged); pre-release; agent machinery — the principal's |
| Canonical primitives | AGenNextHub/Agent-Fabric | spec draft v1.0 (Entity, Relationship, Assertion, Observation, Projection, Drift, Reconciliation) — the assertion-native model; KubeContainer is a running reconciliation of exactly this shape |
| Capability catalog | AGenNextHub/Agent-Skills | spec/early code: skill contracts, permissions, risk class; the hierarchy Skills -> Blueprint -> Team -> Runtime |
| Channels (the chat) | AGenNextHub/Agent-Chat | working Go, v0.1.x loop spine (zero-dep, tested, headless daemon, content-addressed boxes); the native candidate for the space's channels seat, displacing the external Matrix pick when its operator ships; Apache-2.0 - a license divergence from this house's record, named for the principal |
| Trust contracts | AGenNextHub/Agent-Trust | spec (provenance, evidence, audit: "no trust without evidence"); this house's promise machinery is that creed already running |
| The agent internet | AGenNextHub/Internet-of-Agents | named placeholder; AgentRegistry (this house) is its first running piece |
| The knowledge base | AGenNextHub/Agents-Wiki | named placeholder; the record's graph and vocabulary machinery is its natural feed |
| Edge AI compiler | Eclipse Aidge (pinned) | the ModelCompiler seat: model in, surface-native artifact out; first walk owed on the founder's metal |
| Platform backend | the backend contract (registry/SKELETON-BACKEND.json) | founder's ruling: the language is Java - the Apache ecosystem is the toolbox; liferay is the first Java blueprint, pocketbase the lightweight alternative on the same contract |
| The Attic rule | https://attic.apache.org/ | retired Apache projects are a legitimate mine - explicitly forkable; one enters only pinned at its last release and with a keeper named (the ownership doctrine: an empty seat is taken, never squatted) |
| The company | github.com/openautonomyx - OpenAutonomyx (OPC) Private Limited, "The Canonical System for Federated Work" | the platform's legal and product home: dxp (the digital experience platform), trust-center, canonical, search, findr, insights, Publications, campaigns, partners; the decision-intelligence-platform repo is its homepage seat (private or pending at review time) |
| DB candidate | https://surrealdb.com/ - SurrealDB | founder-supplied: multi-model, single binary, embedded mode (the one-engine law holds); Agent-Skills already speaks SurrealQL - the natural backend-contract DB blueprint, pinned when it enters; its home description supplied the foundations line (infinite scale) |
| Positioning reference | https://www.zluri.com/ - Zluri | founder-supplied: the SaaS-management category, the enterprise-work positioning the platform speaks to ("Next Generation Of Enterprise Work") |
| Business-model reference | https://www.redhat.com/en - Red Hat | founder-supplied: the proof that open source plus enterprise trust is the product - subscription on free code, certification as the gate, a verified partner network; the model the platform's trust-center and partner network mirror |
| Category reference | https://www.gartner.com/en - Gartner | founder-supplied: the analyst frame - decision intelligence is a Gartner-named category, and the platform's positioning picks it up as insight-driven everything |
| Peer-trust reference | https://www.gartner.com/peer-insights/home - Gartner Peer Insights | founder-supplied: verified peer reviews as trust evidence - the no-trust-without-evidence law applied socially; the pattern for the trust-center's outward face (reviews as assertions, with provenance) |
| Design platform | https://penpot.app/ - Penpot + Storybook | founder-supplied: Penpot the open design canvas (the Fabric.js doctrine's sibling at the design layer), Storybook the component-truth platform - design and code from one source; the new-work seats (browser server, 3D, documents that carry data) |
| Runtime | this house: KubeContainer on k3s | live on the founder's metal |

The division is the prohibition applied as architecture: this house
supplies the ground every agent stands on — runtime, storage, space
contract, record — and never the agents themselves.
