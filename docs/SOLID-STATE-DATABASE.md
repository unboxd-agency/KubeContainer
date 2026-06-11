# SolidStateDatabase — Product Brief

**Working names:** SolidStateDatabase (SSDB), **SolidBaseData**,
**StructuredSchemas**, or **Fabric of Graphs** — the first emphasizes the
phase (data in its solid state, the SSD echo intended); the second the
role (the solid *base* every other layer builds on — the foundation
principle as a product name); the third the contract (schema-first truth:
structure declared before data crystallizes, the code-is-configuration
principle applied to the record itself — an immutable fact is only as
trustworthy as the schema that shaped it); the fourth the shape (what the
store actually holds: every fact is an edge — entity, field, value, time,
provenance all relations — so each sovereign's record is a graph, and
federation weaves the graphs into fabric without merging them: one weave,
many graphs, queryable along the threads — the "no subgraphs" rule made
architecture, since scoping is authorized traversal of the one weave,
never a copied fragment); or **Kubes of Kubes** — the recursion (the
charter's own geometry as the name: each fact a kube — whole,
indivisible, canonically countable; facts packed into entities, entities
into domains, domains into sovereign stores, stores into the federation —
kubes of kubes all the way up, the same shape at every scale, so the
store is not a container *for* the record but the record built from the
unit the whole platform already speaks). Naming decision deferred to the
trademark/brand pass; this brief uses SolidStateDatabase throughout.

**Status:** concept approved for definition; first product named by the
charter. Per the theory-to-delivery corollary
([FOUNDING-PRINCIPLES.md](FOUNDING-PRINCIPLES.md)): the theory is documented
and protected — what follows is the operating model taking shape.

## One-liner

The system of record the charter keeps invoking, built as a product: a
database where committed facts are **solid** — immutable, bitemporal,
provenance-native, ACID — completing the data phase diagram: lakes hold
liquid data (raw, still, sovereign), oceans move it (shared, by treaty),
and the SolidStateDatabase is where data **crystallizes into truth**.

## The phase model

| Phase | Where | Properties |
|---|---|---|
| Liquid | Lakes & oceans | Raw, schemaless until read, flowing by contract |
| Solid | SolidStateDatabase | Committed, immutable, versioned, attributable |
| Vapor | Compute | Ephemeral working state; never authoritative |

Data freezes (commits) into the solid state through one gate — the
transaction — and never melts: corrections are new crystals referencing the
old, not edits to history.

## Requirements (derived directly from the charter's axiom)

1. **Intent is defined** → every write carries the declared intent and
   named principal that caused it; anonymous writes are unrepresentable.
2. **Action is documented** → the write *is* the audit record: one
   operation produces fact + actor + intent reference + timestamp — there
   is no separate audit log to drift.
3. **All revisions recorded** → append-only, immutable revisions;
   bitemporal axes (valid time + transaction time); point-in-time queries
   ("what was true at T, as known at T'") are first-class, not forensics.
4. **All projections recorded** → reads that feed decisions are themselves
   recordable events: which revisions, what query, for whom, when — the
   context of a decision can be replayed exactly.
5. **All transactions ACID** → serializable commits; a half-written fact is
   unrepresentable; eventual consistency permitted only *between* federated
   stores, never within one.

Plus the charter's standing constraints:

- **Provenance-native** — every fact carries its chain of custody (origin,
  author, authorizing contract); provenance is a query dimension, not
  metadata garnish.
- **Five-dimension addressing** — real (provenance), temporal, geospatial
  (residency as a write-time constraint), domain, and context tags on
  every fact.
- **One field, one writer** — field-level ownership enforced by the store
  (the multi-operator coordination rule, made schema).
- **Sovereign & federated** — per-member stores; cross-store sharing by
  declared contract only; exit takes your data *with provenance attached*.
- **Serverless, not stateless** — runs on object storage (external storage
  is in cloud: S3-compatible, region-pinned), compute layer fully
  disposable and rebuildable from the log.
- **Open standards only** — SQL + a documented change-feed protocol
  (CloudEvents-class) at the boundary; no proprietary wire (the platform
  invents no protocol).

## What exists vs. what's new

Pieces exist in isolation: immutability and bitemporality (XTDB, Datomic),
log-structured storage on object stores, audit columns bolted onto OLTP.
The product gap this fills: **no store today makes intent, provenance, and
projection-recording primitives of the write path** — they are all
afterthoughts in application code, which is exactly why audit trails drift
and AI-era decisions can't be replayed. The SolidStateDatabase is the
record an agent economy requires, where "explainable" means *replayable*.

## MVP sketch

1. Append-only fact log on S3-compatible object storage; ACID commits via
   single-writer-per-partition consensus.
2. Facts as (entity, field, value, valid-time, tx-time, intent-ref,
   principal, provenance-ref) — field-level writer ownership enforced.
3. SQL read surface with `AS OF` (both time axes) and `WITH PROVENANCE`.
4. Recorded-projection API: a read handle that commits its own lineage.
5. Change feed as CloudEvents; per-tenant encryption and region pinning.

## Solid substance: the real twin

The charter's twin doctrine runs both directions, and this product is where
it lands: **solid substance is the real twin of solid-state data.** Every
crystallized fact in the store claims a counterpart in the world — the
delivered package, the settled payment, the running workload, the signed
agreement — and that physical/real counterpart is the *real twin* of the
recorded fact, just as the digital twin is the recorded rehearsal of the
real. The pairing is the product's integrity model:

- A solid fact **without** a real twin is an unbacked claim — detected by
  reconciliation against the surface (evidence expires unless renewed by
  touch).
- Real substance **without** a solid fact is unrecorded reality — the gap
  the axiom forbids, closed by ingestion at the edge where it happened.
- The store's deepest query is therefore twin-correspondence: *show me
  every fact whose real twin no longer answers, and every touch the record
  has not yet crystallized.*

Solid-state data and solid substance, each vouching for the other across
the surface — the record keeping reality honest, reality keeping the
record true.

## The build plan: the toolchain now exists

Every instrument the build requires is already forged, proven, and on
main — the first kube was the toolchain's commissioning run:

| SSDB need | Toolchain piece, already proven |
|---|---|
| Declared store, reconciled | The operator pattern: Kubebuilder scaffold, CRD + controller — a `SolidStateDatabase` kind clearing KUBE-SPEC §7 |
| Append-only, era-stamped facts | The golden-corpus discipline (frozen, append-only, CI-enforced) — generalized from manifests to facts |
| Verdicts on every guarantee | The eval registry: world-tests named in advance, `make eval`, evidence reports |
| ACID-or-nothing commits | The axiom as test plan: the compat suite pattern asserting whole-or-absent writes |
| Provenance-native rows | The LLM/provenance schema from eval/README — already the column set |
| Release with evidence attached | The declared-request release pipeline (release/REQUEST → gauntlet → GHCR → assets) |
| Semantics that cannot drift | The vocabulary system: terms defined before used, checked in CI |
| Eligibility of the artifact | The kube conformance clauses (§7) and the vendor bar (E1–E10) applied to our own second kube |

Build order (each step one PR, each gated, per the protocols):
1. Scaffold the `SolidStateDatabase` CRD + controller in this repo
   (same module, second kind — the kube-of-kubes packing rule).
2. Facts engine v0: append-only log on a PVC, bitemporal keys,
   single-writer-per-partition; world-tests before code (P5 of the
   spec: a guarantee without a verdict is a roadmap item).
3. The `AS OF` read surface; recorded projections.
4. Era-stamp v0 schema into the golden corpus; wire `make eval` tasks.
5. Release as v0.2.0 by declared request — evidence attached, like
   everything else.

The naming decision (five candidates) gates the API group's kind name
only at step 1; the founder rules, the record holds the runner-ups.

## Relationship to KubeContainer

Same charter, different layer: KubeContainer crystallizes *workload intent*
into running reality; the SolidStateDatabase crystallizes *facts* into
durable truth. A future `SolidStateDatabase` CRD — declared stores,
operator-reconciled — is the natural meeting point: the record, itself
declared as a kube.
