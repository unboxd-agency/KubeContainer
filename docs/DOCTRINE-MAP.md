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

## Tier

- **Tier** — a level of binding force: how strongly a passage holds
  the reader, and what it costs to cross it. A tier is not a topic, a
  folder, or a quality grade — it is the *weight* a passage carries:
  whether departing from it is a breach (to be answered), a debate
  (to be argued), a failed gate (to be fixed), a judgment call (to be
  defended), or a refused build (already stopped). Every passage in
  the record sits at exactly one tier per clause, and the tier — not
  the document it lives in — decides how the reader must treat it.
  Tiers are ordered by force: Constitution binds hardest, Harness
  acts hardest, and the rest range between; where a clause could be
  read at two tiers, the stricter governs. The tier is the answer to
  one question asked of every sentence in the corpus: *what happens
  if I do not honor this?*

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
| `docs/DOCTRINE-MAP.md` (this file) | Constitution + Harness | Constitution: it governs how the whole record is read, amendable only by process. Harness: it is the classifier itself — the index that catches orphans, including, now, its own. |
| `docs/TOOLS.md` | Constitution + Constraints | The tool law, casing rule, command words, and reference doctrine bind; the registry rows are constraints (built/declared, owned or out). |
| `docs/PERSONAL.md` | Theory + Constraints | The personal doctrine is theory; the agent's contract terms 1–10 and the flow's gate order are constraints. |
| `deploy/` (STACK, VPS, REFERENCES, the declaration) | Constraints + Best practice | The pins, seams, and decision record are constraints; the runbooks are practice. |
| `registry/` (SKELETON, blueprints, agents) | Harness + Constraints | The skeleton is the admission constraint; the registry is the machine that holds it; append-only, duplicates refused. |
| `cmd/codecompiler`, `cmd/structuredinstructions`, `hack/deployrehearsal` | Harness | Tools and rehearsal: they do not argue, they check. |
| `site/` (the front door, the arithmetic function) | Best practice | The face; it presents the record and computes its own experience — nothing here binds. |
| `docs/CONTROL-PANEL.md` | Theory + Best practice | The panel design and its references; theory until built. |
| `deploy/AGENT-STACK.md`, `deploy/TOOLCHAIN.md`, `deploy/GITLAB.md`, `deploy/ORG-NAME.md`, `deploy/KUBECONTAINER-HOME.md` | Constraints + Best practice | The stack picks bind; the walks advise. |
| `cmd/homesetup`, `hack/operatorpoc`, `hack/deployrehearsal` | Harness | Tools and POCs: they check and demonstrate, they do not argue. |
| `docs/manual/CONTRACTS.md, LICENSES.md, FAQ.md, HOW-TO-USE.md` | Constraints + Best practice | The instruction manual (KUBE-SPEC §10): the terms bind as written in their sources; the manual presents them and names the contacts. |
| `deploy/LEAPMICRO.md` | Constraints + Best practice | The host decision binds (Leap Micro, fixed release, never rolling for a verdict host); the walk advises. |
| `cmd/recordgraph`, `eval/graph.txt`, `eval/graph.jsonld`, `eval/graph-planned.txt` | Harness | The record as a graph, extracted never hand-written; a broken reference fails the gate; the planned file is shrink-only. |
| `cmd/schemakeeper` | Harness | The schema kept as a tool: the graph's JSON-LD validated against the pinned schema.org vocabulary; a dangling node fails the gate. |

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

## The question rule

The founder's law for the whole corpus: if anyone is asking a
question, we have lacked detailing it clearly — a question asked is
an indication that the record is not clear enough. Questions are not
interruptions; they are defect reports against the documentation,
and they are handled the way every defect is handled here: the asker
gets the answer, and the record gets the amendment, at the exact
place the question arose — so the same question can never need
asking twice. The proof this law already runs: "what is critical
path," "why Cinder and not Ceph," "does minikube fit," "where should
I put Keystone" — each was answered in conversation and landed in
the corpus the same hour, which is the loop this rule makes
mandatory: question in, answer out, record amended, clarity
converged. A corpus that keeps receiving the same question has a
failing world-test it refuses to write down; a corpus that treats
every question as a patch grows clearer with every reader — and the
destination is the record this house has aimed at from the first
line: one where the questions run out not because askers gave up,
but because the answers were already standing where the questions
would have stood.

And the bar the question rule aims at, set by the founder: make it
so obvious that the gatekeeper reads, agrees, and opens the gate.
That is the clarity standard in one scene — not "defensible under
questioning" but *never questioned*: the reviewer who needs no
clarifying round, the auditor who finds the evidence already laid in
their path, the compiler that admits on first pass, the buyer who
sees the label and the seal and signs. A gate that opens after
debate was opened by persuasion; a gate that opens on reading was
opened by clarity — and only the second scales, because persuasion
spends a person and clarity spends nothing. Every artifact of this
house is written for that scene: the gatekeeper reads, agrees, and
opens — and if any gatekeeper anywhere hesitates, the question rule
already says what that hesitation is: a defect, ours, to fix where
it stood.
