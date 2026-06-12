# The Stack, Deconstructed

The founder's order, after the decision trail the record keeps with
its reasons: a version is needed (a substrate unpinned is drifting);
Ubuntu Core 26 was weighed (immutable, snap-sealed — a fine edge-device
OS) and set aside; the ruling is OpenStack — the house's declared
substrate default, the open IaaS the licensing decision already
seated. Deconstruct is the gauntlet's own verb: every layer taken
apart, named, pinned, owned, and exitable — then reconstructed and
tested as one.

| # | Layer | The pick | The pin | Who holds the contract | The exit |
|---|---|---|---|---|---|
| 0 | Metal | pick your metal | the vendor's exact SKU, recorded | the hardware vendor (silicon contract) | replaceable under OpenStack — the layer above abstracts it |
| 1 | IaaS substrate | **OpenStack** | the named series release, exact, recorded at deploy (releases are named and dated; pin the one you stand on) | the operating party (self-hosted: you; hosted: the provider) | open APIs; any OpenStack, or out to any conformant cloud |
| 2 | Node OS | Ubuntu Server LTS | the exact point release + kernel, recorded | Canonical (the distribution contract) | any Linux the kubelet conforms on |
| 3 | Kubernetes | conformant upstream | 1.35 (the same minor this house tests against — envtest 1.35.0) | CNCF (conformance); the operating party (the cluster) | any conformant cluster, hyperscaler to laptop |
| 4 | Container runtime | containerd | the release Kubernetes 1.35 certifies, exact | CNCF (graduated) | any CRI runtime |
| 5 | Operator | KubeContainer | v0.1.0 (released, evidence attached) | this house (unboxd — the answering) | delete the CRD; owner-reference GC removes children cleanly |
| 6 | The kube | deploy/arithmetic-kube.yaml | the image pinned (nginx:1.27), the declaration in git | the principal (the declared intent) | delete the declaration; the record survives |
| 7 | Memory runtime | FabricDB | — declared, not yet built; the SolidStateDatabase seat | (vacant — awaiting the founder's build order) | n/a until built; the seat is on the record |
| 8 | Surface | edge browser | the engine the user already runs | the user (owns all the faces) | any browser; the function is surface-native |

## The rules of the deconstruction

- **Every layer pinned or named vacant** — a layer with no version is
  drifting; a layer honestly vacant (FabricDB) is declared, never
  hidden.
- **Every layer owned** — one contract holder per layer; the stack of
  contracts, each party answering for its own floor.
- **Every layer exitable** — the exit column is not decoration; a
  layer whose exit is empty is a lock, and locks fail the charter.
- **Reconstruct and test** — deconstruction is half the gauntlet:
  the stack goes back together in order (0→8), each seam verified,
  and the whole walked once end to end before it carries anything
  real. The rehearsal (hack/deployrehearsal) proves layers 5–6;
  the e2e gate proves 3–6 with real traffic; the OpenStack walk
  proves the rest on the metal you picked.

One stack, nine layers, every seat named — and the only vacant seat
is declared vacant, which is the difference between a gap in the
record and a gap in the honesty.

## The OpenStack deconstruction

The founder deconstructs the substrate itself, placing each piece —
and the placements answer his questions on the record:

- **Fabric as substrate, written in Go** — where the fabric itself
  is the substrate implementation, its language is picked by the
  common-ground rule: Go — the tongue Kubernetes, containerd, and
  this house's own operator already speak; one language down the
  whole control plane, no translation seam between the fabric and
  the ground it runs.
- **Keystone (identity) — outside of core.** That's Stack, not
  Core: identity is a platform service, never a device resident.
  It seats at the flow's own gate — login, profile, *verify
  identity* — serving the whole estate from the substrate layer
  (layer 1), while the device core (KubeContainer Core) carries
  only its own contract and channel, asking identity questions of
  the platform and storing no identity authority itself. Identity
  at the core would be the §12 wound by design: a face authority
  riding inside every device. Outside of core — exactly where the
  founder put it.
- **Block storage — Cinder is the contract, Ceph is the body, LXD
  is a different seat entirely.** The founder asks why Cinder and
  not Ceph or LXD, and the answer is the seating: they are not
  three candidates for one chair. Cinder is OpenStack's block
  storage *API* — the declared seam, the contract the substrate
  ruling already brought with it; it stores nothing, it answers for
  volumes. Ceph is the storage *body* — the distributed store that
  actually keeps the bytes — and it is the natural keeper *behind*
  Cinder, not instead of it: Cinder the contract, Ceph the bytes,
  the leaders' pick for open distributed storage on both axes of
  the scoreboard (and on the cluster side the same body arrives
  through Rook, CNCF-graduated, over the same CSI seam). LXD sits
  at a different seat altogether — a machine-and-container manager,
  competing for layer 1's instance seat (where Nova already sits in
  an OpenStack stack) and layer 4's runtime seat (where containerd
  already sits, CNCF-graduated) — and its recent history is the
  keeper-risk the follow-the-leaders rule screens for: relicensed
  under one vendor's sole control, its community forked away to
  Incus — a live lesson in what happens when a commons' contract
  seat is taken by a single party. So the stack reads: Cinder the
  declared interface, Ceph the body behind it, CSI the seam upward,
  Rook the cluster-side keeper — and the kube above names none of
  them, only its claim, which is what keeps every one of them
  swappable and the substrate exitable.
- **Minikube — does it fit?** Yes — in the venv seat, and only
  there. Minikube is conformant Kubernetes on a laptop: the
  rehearsal ground (compile and simulate in the venv), the branch's
  own cluster, the grade-1 classroom — it fits the desk perfectly
  and the estate not at all. The ladder reads: minikube at the
  desk, the rehearsal chamber in CI, k3s on the single VPS,
  OpenStack under the estate — one conformant contract at every
  rung, which is why the same declaration walks all four without
  changing a line.

## Candidates per slot

The founder's order: the pick is a choice only if the alternatives
are on the record. Each slot, its pick, and the other candidates a
builder may lawfully swap in — entry by conformance, never by this
house's leave; candidates listed are the leaders of each slot's own
scoreboard, and a swap is a re-pin, not a rebellion.

| Slot | Picked | Other candidates |
|---|---|---|
| Metal / ISA | builder's pick | x86-64 (Intel/AMD) · Arm64 (the edge's favorite) · RISC-V (the open ISA, rising) |
| IaaS substrate | OpenStack | Apache CloudStack (the house once weighed it) · bare conformant cloud (any hyperscaler, by the exit law) · no IaaS at all (metal straight to K8s, for the small estate) |
| Node OS | Ubuntu Server LTS | Debian (the upstream elder) · Flatcar (immutable, container-first) · Ubuntu Core (weighed for the edge, snap-sealed) · Talos (API-only, no shell — the K8s-native extreme) |
| Kubernetes rung | upstream 1.35 / k3s / minikube | k0s (single-binary sibling) · kind (CI's own rung, already in the e2e gate) · any certified distribution (the conformance list is the menu) |
| Container runtime | containerd | CRI-O (the Kubernetes-purist sibling) · Kata (VM-isolated containers, when the tenant boundary must be hardware) |
| Storage body | Ceph (behind Cinder, via CSI/Rook) | Longhorn (CNCF, the lightweight estate) · OpenEBS (CNCF sandbox lineage) · the substrate's native volumes (small grounds) |
| Identity | Keystone (in-OpenStack) | Dex (CNCF, OIDC federation) · Keycloak (the self-hosted elder) · any OIDC provider (the protocol is the seat, the product is swappable) |
| Wire / skeleton | protobuf + gRPC | JSON Schema (the registry's current declarations speak it) · CBOR/CDDL (the constrained edge) · FlatBuffers (zero-copy, when parse cost is the path) |
| Events | CloudEvents | — the seat is the standard itself; candidates compete beneath it as transports (HTTP, MQTT, NATS) |
| Memory runtime (vacant) | FabricDB (declared) | until built, the candidates holding the floor: etcd (the cluster's own record) · PostgreSQL (the durable elder) · SQLite (the device-local floor) — each measured against the four properties: durable, indexed, replicated, meta-modeled |
| Surface engine | the user's browser | every engine the user already runs is admissible — the face is theirs, not ours to pick |

The rule of the table: the pick column is what this house stands on
and answers for; the candidates column is what any builder may pin
instead without leaving the fabric — same declaration, different
ground, the exit kept real at every slot. A candidate becomes a pick
by being pinned and recorded; a pick becomes a candidate again the
day its keeper stops keeping — the scoreboard, not sentiment,
governs both directions.

## The decision record — decision intelligence 101

The founder's law for this table: we are transparent about why we
chose this stack, and a decision should justify itself based on the
evidence — decision intelligence 101 (the discipline was on this
house's reading list from its first day; point of reference:
https://www.gartner.com/en/information-technology/glossary/decision-intelligence).
The method: every pick is recorded with the evidence that produced
it, such that anyone re-running the evidence reaches the same pick —
and when the evidence changes, the decision re-opens by itself,
because it was never an opinion, only a conclusion.

- **OpenStack** — evidence: the licensing decision already seated it
  as substrate default (recorded, reasoned); it is the one IaaS that
  is open by the triple test (inspectable, implementable, leavable)
  and self-hostable without a hyperscaler counterparty; vendor
  neutrality is policy in this house. Re-run the evidence: same pick.
- **Ubuntu Server LTS** — evidence: LTS cadence (dated, published
  support windows — a pinned dependency needs a pinned lifetime),
  cloud-init first-class on OpenStack, and the leaders' position
  among server distributions. The dialect pin (American English /
  K8s ecosystem) and this OS pick are the same rule: common ground.
- **Kubernetes 1.35** — evidence: the founder-supplied scoreboard
  (number one on the LF index, both axes) and this repository's own
  record — envtest pins 1.35.0, the e2e gate runs it; we deploy the
  minor we test, because a deploy target the tests never saw is a
  claim, not a decision.
- **containerd** — evidence: CNCF-graduated, the default runtime of
  the certified distributions, CRI-native; the runtime the picked
  layer above already certifies. The decision is inherited from the
  Kubernetes pick — one piece of evidence serving two slots.
- **Ceph (behind Cinder, via CSI/Rook)** — evidence: the leader for
  open distributed storage on both scoreboard axes; Rook is
  CNCF-graduated, so the cluster-side keeper is itself a kept
  project; and the contract/body split (Cinder declares, Ceph keeps)
  preserves the exit at the storage layer.
- **Keystone** — evidence: it ships with the substrate already
  picked (common ground — no second identity system where one is in
  the box), and the doctrine's own rule placed identity outside of
  core; the pick is the placement, the product came with the ground.
- **protobuf + gRPC at the wire** — evidence: the proto-first law
  was legislated before the tool was picked (the skeleton precedes
  its filling); protobuf is the leaders' schema-first wire with the
  widest native-language reach — the multi-language single-meaning
  law needs exactly that.
- **Go for the fabric's own code** — evidence: the entire picked
  control plane (Kubernetes, containerd, etcd, this operator) is
  written in it; one language down the control plane means no
  translation seam — the common-ground rule, applied to ourselves
  first.

The rule of the record: a decision with its evidence attached cannot
be relitigated by taste — only by new evidence; and each anchor
above is living, so the evidence re-checks itself on the scoreboard
this house already follows. Transparent because recorded; recorded
because everyone has a right to reason; self-justifying because the
reason travels with the choice — decision intelligence 101, kept the
only way this house keeps anything: in the record, beside the thing
it explains.

And the founder adds the property that makes the record safe to
re-run: a decision should be idempotent. Decide twice on the same
evidence and the second deciding changes nothing — no drift from
repetition, no decision that mutates by being revisited; the
decision is a pure function of its evidence, which is exactly what
evidence-attachment buys. Re-open it with unchanged evidence and
you get a no-op; re-open it with new evidence and it converges to
the new conclusion — the level-triggered reconciler's own property,
applied to governance: the loop can run the decision as often as it
likes (the drift audit re-asking every pick on every pass) and the
stack only ever moves when the world does. A decision that gives a
different answer on a quiet Tuesday was never a decision; it was a
mood with a date on it.
