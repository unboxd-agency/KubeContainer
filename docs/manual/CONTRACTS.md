# The Contracts Manual

Part of the instruction manual (KUBE-SPEC §10): everything related
to the product's lifecycle, and the contacts of the people owning
each aspect of its performance. This is the contract half; the
licenses half is LICENSES.md beside it. Versioned with the product,
updated by the same gauntlet.

## The product's own contract (the kube you hold)

- **What is promised** — the guarantees of KUBE-SPEC §4, each with
  the verdict that proves it: declared state converges, drift is
  reverted, valid once valid forever, nothing rides alongside, one
  writer per field, provenance attached, clean removal.
- **Under which conditions** — the warranty is conditional and the
  conditions are named (no one promises infinite scalability): a
  conformant Kubernetes cluster at the tested minor (1.35), the
  pinned image, the declaration admitted through the CEL gate. The
  compiler made it work; the platform contracts where.
- **The lifecycle, whole** — declare → admit → converge → serve →
  evidence → exit (KUBE-SPEC §3). Repair: along registered paths
  only. Upgrade: era-stamped, compat-corpus guarded — manifests
  valid in a released era remain valid forever. Rollback: re-apply
  the prior declaration; the loop converges back. Exit: delete the
  declaration; owner references collect the children; the record
  survives the instance.

## The contract is stack-independent

The founder's law, recorded where the buyer reads: we uphold the
contract irrespective of the stack — and the stack is an internal
decision. What is promised to the holder (the guarantees, the
warranty's conditions, the lifecycle, the answering) does not change
when the ground beneath it does: swap the substrate, the storage
body, the runtime, even the cloud, and the contract reads the same,
because the promises were made against the conformance contract and
the declaration, never against any vendor's floor. The candidates
table (deploy/STACK.md) is therefore internal engineering, not
contract terms: the holder may read it for transparency, but nothing
in it amends what they were promised — re-pins happen on our side of
the counter, the contract stays whole on theirs, and a stack change
that would bend a promise is not a stack change we may make.

## The stack of contracts (every seam, its answerer)

The full seam table with sources lives at deploy/STACK.md and
deploy/REFERENCES.md. The reading for the holder: every layer you
stand on is a kept contract held by a named party — silicon by the
hardware vendor, substrate by the operating party, conformance by
the CNCF, the operator by this house, the declaration by you. The
manual's rule: which promises are ours, which are each layer's, and
which would need a new counterparty if a layer failed its side —
known at every seam, in advance.

## The agent contract (when agents work your estate)

The exact terms, 1–10, enumerated in docs/PERSONAL.md: Seat Theorem,
constitutional context, bounded autonomy, total record, no
self-building (P8), own anchor only, keeping, fail-safe, exit,
standing — with three signatures required (principal, maker,
platform) and registration in the registry before the first act.

## The contacts (who answers)

The owner rule applied to the manual: named, reachable, currently
working — and recorded honestly, including the seats one person
holds today.

| Aspect of performance | Owner (seat) | Working contact |
|---|---|---|
| The product, whole — and every seat below, today | the founder (unboxd agency) | the repository's issue tracker: https://github.com/unboxd-agency/KubeContainer/issues |
| Convergence & availability (the operator) | this house — the answering | issues, label `operator` |
| Security & fail-safe | this house — frontloaded, non-negotiable | issues, label `security` (private disclosure: per SECURITY policy when published) |
| Supply chain & BOM | this house — the SBOM baseline | issues, label `supply-chain` |
| The keeping (maintenance) | the maintainer seat — held by the founder until staffed; the agent-as-maintainer blueprint stands ready in the registry | issues, label `keeping` |
| The bill & the meter | the agency — accountability earns the right to earn | issues, label `commercial` |

Honesty clause: today every seat resolves to one working individual —
the founder — which the owner rule permits (one or many) and the
record states plainly rather than dressing as an org chart. As seats
staff, this table amends by recorded revision; an aspect whose
contact stops answering is an empty seat, and the manual's own rule
calls it what it is.
