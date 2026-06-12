# The Kube — Product Specification

Status: v1, normative. Terms used here are defined in
[AGENT-PLATFORM.md](AGENT-PLATFORM.md) and governed by
[FOUNDING-PRINCIPLES.md](FOUNDING-PRINCIPLES.md). The reference
implementation is this repository's `KubeContainer` (v0.1.0, released).

## 1. Definition

A kube is the platform's unit of product: a declared outcome, kept by an
autonomous loop, presenting one contractual face to the world, with its
provenance attached. It is whole and indivisible — there is no fraction
of a kube — and it is what a customer buys, what the registry records,
what the surface is tiled from, and what the fabric packs.

One sentence for the buyer: *you declare what you want; the kube makes
it true, keeps it true, proves it was true, and answers for it.*

## 1a. Meaning and measurement

A kube is a box with meaning (what it is) and a measurement (how much it
is worth). Both are expressed in real-world, internationally accepted
standards by default; where none exists, the contract holder defines and
delivers the measure explicitly, marked as holder-defined. The BOM/SCM
carries no prior transaction before the end-user touch point. Full rule:
[MEASUREMENT-STANDARD.md](MEASUREMENT-STANDARD.md).

## 2. Anatomy

Every kube consists of exactly five parts; remove any one and it is not
a kube.

| Part | What it is | KubeContainer v0.1.0 instance |
|---|---|---|
| Declaration | The spec: intent in schema, validated before admission | The `KubeContainer` CR (CEL-validated) |
| Loop | The autonomous keeper: observe, compare, act, record | The reconciler (level-triggered, idempotent) |
| Face | The one stable surface point it answers at | `status.endpoint` (Service/Ingress) |
| Record | Its state, history, and evidence, externalized | spec/status in etcd; events; conditions |
| Contract | The promises, with their verdicts named in advance | Ready/Progressing/Degraded; the gauntlet |

## 3. Lifecycle

Declare → Admit → Converge → Serve → Evidence → Exit.

1. *Declare:* the principal states the outcome (file, CR, request).
   Vague intent is returned for definition, not guessed at.
2. *Admit:* schema validation, policy, authorization — the gates. A
   kube that cannot be admitted does not partially exist.
3. *Converge:* the loop closes the distance to the declaration, one
   adjacent step at a time, without jumps.
4. *Serve:* the face answers — stable, minimal, guarded. The kube is
   now load-bearing at its point of the surface.
5. *Evidence:* continuously — status reflects truth, events record
   actions, verdicts accumulate in the record. Done is only ever
   claimed by citing a verdict.
6. *Exit:* graceful by design — work landed or released, state
   committed, the face handed over or retired, the name and record
   surviving the instance.

## 4. Guarantees (each with its verdict)

| Guarantee | Verdict that proves it |
|---|---|
| Declared state converges | Ready condition; e2e gate (real traffic) |
| Drift is reverted | Reconcile on child mutation (tested) |
| Valid once, valid forever | Golden compat corpus (append-only, CI-enforced) |
| Nothing rides alongside | One workload container; no sidecars (tested) |
| One writer per field | HPA non-interference (tested) |
| Provenance attached | Evidence report shipped with every release |
| Clean removal | Owner-reference garbage collection (tested) |

A guarantee without a listed verdict is a roadmap item, not a guarantee.

## 5. Interfaces

- Declarative: `kubectl apply` of the CR — the only required interface;
  works on any conformant Kubernetes cluster.
- Observational: `kubectl get` (printer columns), status conditions,
  events, Prometheus metrics.
- No proprietary wire, no required CLI, no console dependency: the API
  is the product surface; consoles and tools are adapters.

## 6. Non-goals (what a kube is not)

- Not a template engine, not a package manager, not a PaaS console.
- Not a fraction: there are no partial deployments of a kube, no
  degraded tiers sold as whole ones.
- Not a cage: deleting the declaration removes the kube cleanly; the
  record of it remains (exit ends the acting, never the accountability).

## 7. Conformance (what may call itself a kube)

A thing is a kube if and only if:

1. its outcome is declared in a published schema and validated at
   admission;
2. an autonomous loop keeps it, and the loop survives restart with
   nothing lost (state external, ACID);
3. it presents one identified, stable face, owned by exactly one
   contract;
4. every guarantee it claims names the verdict that proves it, and the
   verdicts run in CI or in cluster — not in slideware;
5. its history is recorded, append-only, and attached to what it
   delivers.

KubeContainer v0.1.0 conforms (verdicts in §4). Future kinds — Release,
SolidStateDatabase, agent workloads — must clear the same five clauses
before wearing the name.

## 8. The product line

The kube is the unit; products are counts and kinds of kubes:

- One kube: the pilot (the MVP — the whole thesis at quantity one).
- Many kubes, one tenant: the estate (governance, drift audits, eval
  registry over their workloads).
- Many kinds: the platform (new CRDs clearing §7).
- Many tenants, many fabrics: the economy (federation, certification,
  the partner network).

Scaling is packing, not rebuilding: the difference between MVP and
platform is count, not kind.

And the product line gains its edge edition, named by the founder
on the Ubuntu Core precedent: **KubeContainer Core** — the kube at
the device's core, as a product. The minimal, sealed distribution
of the container runtime for edge devices: the operator and the
one-kube communicator packaged as the immutable center of a device —
the loop, the record, the contract, and the single signed channel,
nothing else aboard (Single Binary Code as a product SKU). The
surface around it stays headless; the core inside it runs the kube;
delivery reaches it only at the ports — the whole edge doctrine,
shipped as the thing you flash. Where KubeContainer runs workloads
on clusters, KubeContainer Core *is* the workload-keeper a device
boots: edge with kube at its core, in the catalog.

## 9. The brand

The brand line is the founder's, recorded verbatim with his styling:
**Kube — the soul of Any'Thing'**.

The claim is mechanical, not mystical, and every word of it is already
load-bearing in this spec:

- A *Thing* is any identifiable artifact at a touch point of the
  surface, real or digital (lexicon: *Thing*). A Thing without a kube
  is a body — present, connected perhaps, and inert: no declared
  intent, no loop keeping it, no record surviving it, no one
  answering for it. The box as mere artifact.
- The *soul* is what the kube adds, and it is exactly the five-part
  anatomy of §2 read as the offices a soul has always been asked to
  hold: it *individuates* (the Declaration — this Thing's intent,
  distinct from every other's), it *animates* (the Loop — the Thing
  acts, keeps itself true, returns when struck down), it *faces the
  world* (the Face — one point where the Thing answers), it
  *remembers* (the Record), and it *answers* (the Contract). And it
  survives the body: exit ends the acting, never the accountability —
  the name and record outlive the instance (§3, step 6), which is the
  one property everything ever called a soul has been required to
  have.
- *Any* is the conformance clause (§7) read as the market: whatever
  clears the five clauses — workload, release, database, agent
  workload — gets a soul, and the product line (§8) is counts and
  kinds of ensouled Things.

The lineage is deliberate. The industry gave Things connectivity and
called it the Internet of Things — wired bodies, still soulless. The
kube supplies what the wiring never did, and the brand names the
consequence: give the Thing a soul and the Internet of Things becomes
the Internet of Agents. The phrase also honors its ancestor — *The
Soul of a New Machine* (Kidder, 1981) named the moment a built thing
first seemed to carry something of its makers; this brand makes that
metaphor checkable: the soul is five parts, listed in §2, each with a
verdict in §4.

Brand discipline: the protectable mark is the full line, never the
bare word — *soul* alone is unownable and is not claimed. The line is
claimed in NOTICE alongside the marks it extends, and defined in the
lexicon before use, per protocol P2.
