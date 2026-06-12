# The Licenses Manual

The second half of the instruction manual (KUBE-SPEC §10): what you
may do with what you hold, under whose grant, with every term
anchored to its source.

## The license of record

The Software is licensed under the **KubeContainer Research and
Community Source License, Version 1.0 (June 2026)** — the LICENSE
file at the repository root is the authoritative text; this manual
summarizes and the text governs.

- **Source-available, not OSI open source** — by its own honest
  NOTICE: commercial rights are reserved exclusively to unboxd
  agency. The open≠free law, written into the grant itself: seeing
  is the code's nature; using commercially is a grant from its
  contract.
- **The non-commercial grant** — research, study, education,
  evaluation, interoperability testing, security review, community
  discussion, and contribution: granted worldwide, royalty-free,
  revocable, for Non-Commercial Purposes (the license's §1
  definitions govern the boundary).
- **Community sharing** — copies and modifications may be shared
  for Non-Commercial Purposes with the license intact, notices
  retained, modifications identified, and no representation as an
  official release.
- **Commercial use** — reserved to the Licensor: production
  operations for a business, paid services, hosting, managed
  services, resale, inclusion in commercial products. The
  commercial channel is the agency's (everything commercial, and
  an open market: the answering is what is sold).
- **The marks** — unboxd, KubeContainer, the Fabric, Kube, and the
  brand line are claimed in NOTICE and are not licensed by any
  grant here.

## One inconsistency, flagged for the founder

NOTICE (Constitution tier — amendable only by recorded process)
still carries two references to the Apache License, Version 2.0,
written before the license of record changed: its header line and
the triple-test clause ("Implementable — Apache 2.0"). The LICENSE
file governs; the NOTICE references are stale and await the
founder's amendment — flagged here per the question rule rather
than silently edited, because constitutional text is the
principal's to amend.

## Dependency licenses (what this product stands on)

The Go module graph is the dependency BOM (`go.mod` pins it exact).
The load-bearing dependencies and their licenses, each at its
source: Kubernetes libraries and controller-runtime — Apache-2.0
(https://github.com/kubernetes/kubernetes/blob/master/LICENSE,
https://github.com/kubernetes-sigs/controller-runtime/blob/main/LICENSE);
the Go standard library — BSD-3-Clause
(https://go.dev/LICENSE). Their notices travel with their code;
nothing in this product's license alters their terms (their
contracts are their keepers').

## The decision's record

Why this shape and not another is a recorded decision with its
reasons: docs/LICENSING-DECISION.md (Apache permanently excluded;
the candidates weighed; everything-SaaS; the Red Hat model adapted) —
read it the way every decision here is read: the evidence attached,
the conclusion re-runnable, idempotent until the evidence moves.
