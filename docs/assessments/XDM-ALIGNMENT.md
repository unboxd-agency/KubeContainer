# Alignment Assessment: Platform Data Model vs. Adobe XDM

**Subject:** the platform data model defined in `docs/AGENT-PLATFORM.md`,
`docs/FOUNDING-PRINCIPLES.md`, and `docs/SOLID-STATE-DATABASE.md` ("our
system"), assessed against the **Adobe Experience Data Model (XDM)** — the
schema framework behind Adobe Experience Platform.

**Method & honesty note:** Adobe's documentation site refuses automated
retrieval (HTTP 403), so this assessment is grounded in XDM's publicly
documented structure (schema = class + field groups; record vs. time-series
behaviors; Schema Registry with non-destructive, additive-only evolution;
namespaced identity via `identityMap`; field-level data-usage labels), not a
live crawl. Scores are a structured expert assessment against a published
rubric — points are argued, not measured by machine. Re-score when a live
XDM schema export is available.

## Rubric

| Metric | Question | Scale |
|---|---|---|
| **Conformance** | If our facts were mapped into XDM's structural rules today, how much already complies? | 0–100, higher better |
| **Convergence** | Are the two models evolving toward the same concepts? | 0–100, higher better |
| **Drift** | Expected rate of divergence over time without active alignment | 0–100, *lower* better |
| **Distance** | Current semantic + structural gap | 0.0 (identical) – 1.0 (unrelated) |

## Scores

| Area | Conformance | Notes |
|---|---|---|
| Schema-first composition (class + field groups ≈ kind + field ownership) | 18/20 | Both compose schemas from a base type plus declared field groups; our CRD/OpenAPI + CEL is structurally equivalent composition |
| Behaviors: record vs. time-series | 16/20 | Direct map: solid facts/entities = record; the event log = time-series (ExperienceEvent analog, timestamps mandatory in both) |
| Non-destructive evolution | 19/20 | XDM's additive-only rule is our backward-compatibility principle verbatim — and ours is now enforced by a golden-corpus test, which XDM itself doesn't require |
| Identity | 13/20 | Both demand unique, namespaced identity; XDM's `identityMap` + identity-graph stitching has no implemented counterpart here yet (designed: independently-resolvable IDs, delegation chains) |
| Governance labels (DULE ≈ policy/authorization) | 12/20 | XDM labels fields for usage policy; we govern via OPA/OpenFGA at admission/access — same intent, different attachment point (field-level labeling not yet in our schema) |
| **Conformance total** | **78/100** | |

| Metric | Score | Reading |
|---|---|---|
| **Convergence** | **86/100** | Strong and accelerating: schema-first, additive-only evolution, mandatory timestamps, identity-centricity, and policy-labeled data are *both* systems' direction of travel; our five data dimensions (real/temporal/geospatial/domain/context) are a superset of XDM's record/time-series + identity worldview |
| **Drift** | **22/100** (low) | Structural drift risk is low (both anchor to JSON-Schema-class contracts); *semantic* drift is the real exposure — XDM's ontology is customer-experience (profiles, segments, journeys), ours is work (intents, outcomes, agents); vocabularies will diverge even as structures rhyme |
| **Distance** | **0.31** | Close in skeleton, apart in flesh: an XDM ExperienceEvent and our recorded action are the same shape (timestamped, identified, schema-validated, immutable); the domain vocabularies and the registry/union-schema machinery account for most of the remaining gap |

## Composite

> **Alignment score: 78 conformant / 86 converging / 22 drifting / 0.31 distant**
> — structurally near-conformant, directionally aligned, semantically distinct
> by domain (by design, per the multi-domain principle: deliberate mapped
> joins, never naive merging).

## What would close the gap (if XDM interop became a goal)

1. **Field-level usage labels** in the CRD/SSDB schema (DULE-equivalent) —
   our biggest conformance hole; OPA policies already consume such labels
   naturally.
2. **identityMap-style namespaced identity** on facts — our identity
   principle specified it; nothing implements it yet.
3. **A schema registry surface** (publish, version, union-view) — the SSDB
   brief's schema-first naming candidate ("StructuredSchemas") is this
   component.
4. **A declared XDM bridge** — per the charter: a mapped join between the
   work domain and the experience domain, as a contract, never a merge.
