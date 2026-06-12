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
- **Block storage (Cinder) — layer 1's storage arm**, surfacing
  upward only through the conformant seam: Cinder serves volumes at
  the substrate; Kubernetes consumes them at layer 3 through CSI;
  the kube above never names Cinder, only its declared claim — so
  storage stays swappable (any CSI ground) and the substrate stays
  exitable.
- **Minikube — does it fit?** Yes — in the venv seat, and only
  there. Minikube is conformant Kubernetes on a laptop: the
  rehearsal ground (compile and simulate in the venv), the branch's
  own cluster, the grade-1 classroom — it fits the desk perfectly
  and the estate not at all. The ladder reads: minikube at the
  desk, the rehearsal chamber in CI, k3s on the single VPS,
  OpenStack under the estate — one conformant contract at every
  rung, which is why the same declaration walks all four without
  changing a line.
