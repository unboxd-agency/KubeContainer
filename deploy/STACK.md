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
