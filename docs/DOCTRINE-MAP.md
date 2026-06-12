# The Doctrine Map — Tiers of the Record

Status: v1, normative index. This document does not contain doctrine;
it *classifies* it. Every body of work in this repository is named
here and placed in its tier, so that a reader knows at a glance
whether a passage is law to obey, a theory to weigh, a constraint to
respect, a practice to follow, or a machine that checks. The material
itself stays exactly where it lives today.

## First principle: backward compatibility

This map is built to protect the current ecosystem, and its first
principle is backward compatibility. Therefore:

- **Nothing moves.** No document is relocated, renamed, or split to
  build this map. Every existing path, cross-reference, `§`-pointer,
  and link keeps resolving exactly as before — the classification is
  layered *on top of* the corpus, never carved *into* it.
- **The map is additive and append-only.** New material is classified
  by adding a line here; existing classifications are amended only by
  recorded, reasoned revision, never by quiet edit — the same regime
  the founding principles and the golden compat corpus already keep.
- **A label is not a move.** Naming a passage *Theory* or *Constraint*
  does not change its location or its governance; it only tells the
  reader how to read it. The stricter tier always binds where one
  body carries more than one.

The reason is the doctrine's own: relocating the record to tidy it
would break the references that make it one fabric — and a tidy that
breaks the weave has failed the first principle to serve the second.

## The five tiers

- **Constitution** — binding law: definitions that fix meaning,
  prohibitions that forbid acts, and the safety standards that may
  not be waived. Amendable only by recorded process; the stricter
  rule binds; a breach is named, recorded, and answered. To violate
  the Constitution is not a mistake but a breach.
- **Theory** — a claimed model of how the world works, offered to be
  weighed and tested, not obeyed. A theory earns assent by argument
  and evidence; disagreeing with a theory is debate, not breach. Much
  of the doctrine is theory that the Constitution then makes binding
  only at named points.
- **Constraints** — the bounds the work must stay inside: the
  invariants, the gates, the non-negotiables, the conditions a thing
  must meet to be admitted. A constraint is narrower than a law and
  sharper than a practice — a line the build either clears or does
  not, often enforced by the harness.
- **Best practice** — guidance for how the work is well done: the
  creed, the method, the disposition. Recommended, not commanded;
  departing from a best practice is a judgment call to be defended,
  never a breach. The wise follow it; the record does not punish its
  absence, only its absence's results.
- **Harness** — the machinery that enforces the tiers above: the
  scripts, the gates, the registry, the audits, the CI. The harness
  is code, not prose — it does not argue, it checks; what it refuses
  does not ship. Where a constraint is mechanical, the harness is
  where it actually lives.

## The classification of the current record

Each body is named with its tier(s); the document stays where it is.

| Body / document | Tier(s) | Note |
|---|---|---|
| `docs/FOUNDING-PRINCIPLES.md` | Constitution | The charter; twenty-four principles, normative. |
| `docs/PRIMITIVES.md` — the entries | Constitution | Definitions fix meaning; a primitive means exactly this, everywhere. |
| `docs/PRIMITIVES.md` — "The counsel" | Best practice | The working creed; followed, not enforced. |
| `docs/AGENT-PLATFORM.md` — lexicon & ladder | Constitution + Theory | Definitions bind; the ladder and capabilities are model. |
| `docs/AGENT-PLATFORM.md` — protocols P1–P8 | Constitution | Binding; P8 absolute. |
| `docs/HEADLESS-DELIVERY.md` | Theory + Constraints | A theory of delivery that imposes named constraints (safety standards, the gates, the no-bleed surface). |
| `docs/MEASUREMENT-STANDARD.md` | Constitution + Theory | The standards rule binds; the worked models are theory. |
| `docs/KUBE-SPEC.md` | Constraints + Theory | §7 conformance clauses are constraints; the rest specifies and explains. |
| `docs/BENCHMARKS.md` | Constraints + Best practice | The law of the bench is a constraint; the desired-state posture is practice. |
| `docs/VENDOR-ELIGIBILITY.md` | Constraints | E1–E10: the bar a vendor must clear. |
| `docs/LICENSING-DECISION.md` | Constitution | Apache permanently excluded; the decision binds. |
| `NOTICE` | Constitution | Marks, copyright, the canonical identity. |
| `Makefile`, `hack/*.sh`, `.github/workflows/*` | Harness | The gauntlet, vocab check, eval registry, drift audit, CI. |
| `internal/controller/testdata/compat/` | Harness + Constraints | The frozen golden corpus; the backward-compatibility constraint, mechanically enforced. |
| `eval/corpus/`, `eval/harness.sh` | Harness | The world-tests and the runner. |

Where a row carries two tiers, the stricter binds the passage in
question: a Theory paragraph inside a Constitution document is still
weighed as theory, but a Constitution clause inside a Theory document
still binds as law. The reader resolves by clause, not by cover.

## How to use this map

- Writing new doctrine? Name its tier here in the same change that
  introduces it — an unclassified body is an orphan, and orphans are
  rejected.
- Citing a passage? Cite its tier when the weight matters: "a
  constraint (KUBE-SPEC §7)" reads differently from "a theory
  (HEADLESS-DELIVERY)."
- Amending a classification? By recorded revision with its reason,
  never by quiet relabel — this map is held to the regime it
  describes.

The map protects the ecosystem by labeling it where it stands; the
first principle is kept because nothing the map does could break a
single reference in the record it indexes.
