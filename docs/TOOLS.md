# The Tools

Status: v1, normative for tool naming and the tool registry. A tool,
by the primitive, is a thing used to act that does nothing on its
own; this document is the law of how the house's tools are named,
shaped, and registered — and the registry itself.

## The tool law

- **Two-word composite, one binary** — every tool of the house is
  named as a composite of exactly two words fused into one name, and
  ships as exactly one binary (Single Binary Code applied to the
  toolbox: no helper processes, no sidecar scripts, the name tells
  you what it does and the binary is all there is). The precedent is
  the founder's: *Unicode* — two words, one standard, so fused the
  seam is invisible — the proof that a two-word composite can carry
  a whole world's agreement under one name.
- **The tool equation** — Kube+Code = Thing: give code a kube (the
  declaration, the loop, the face, the record, the contract) and the
  result is a Thing — bounded, named, answerable. Tools are how the
  equation is worked in practice, and every tool in this registry
  serves some step of it.
- **The Any' family** — tools and the equation hold on AnyCompiler,
  in AnyLanguage or AnyScript, over AnyProtocol, at AnySurface, on
  AnyMetal, against AnyDB, in AnyCloud, on AnyStack: the founder's
  Any' styling extended to the axes of portability — compiler,
  compiled language, scripted language, protocol, surface, hardware,
  store, cloud, and stack. The family is open the way the fabric is
  open: a new axis enters when work cannot be said portable without
  it. A tool that holds on every axis is a tool of this house; a
  tool that holds on one vendor's axis is a guest of that vendor.

## The registry

One entry per tool; built tools name their command, declared tools
name their intent; nothing removed, nothing duplicated.

- **CodeCompiler** — built. One binary answering the counsel's
  standing question: does the code compile, and does it conform? It
  runs the admission gates in order — compile, vet, vocabulary — and
  returns one verdict; it changes nothing and acts only when
  invoked. Source: `cmd/codecompiler/main.go`; build:
  `go build ./cmd/codecompiler`; run from the repo root:
  `./codecompiler`.
- **StructuredInstructions** — declared, not yet built. One binary
  for the proto-first law of the wire: declare a skeleton, validate
  a block against its skeleton, and refuse the undeclared shape —
  the edge's admission gate as a tool in the hand.
- **SourceGround** — working, as scripts to be fused into one
  binary. The tool of the founder's law that you define a term
  before using it and bind the meaning from the source, with source
  verification: Every'Word' in the corpus is source-grounded —
  extracted from its defining line, indexed with its exact location
  (`docs/VOCABULARY.md` maps every term to the file and line that
  defines it), and every bold use anywhere is checked against that
  index, so no word is ever used whose meaning cannot be walked back
  to its source. Today it lives as `hack/build-vocabulary.sh` and
  `hack/check-vocabulary.sh` (run as `make vocab` and
  `make vocab-check`); its registry seat is held for the one-binary
  form. The reason the tool must exist is the founder's: because
  everything is a reference — every word points at a meaning, every
  meaning at a source, every source at a record — and a chain of
  references resolves only if somewhere it lands: you need an
  anchor, the one defining line a chain bottoms out at. Or you bind
  the open end as a variable, declared rather than dangling — the
  domain as a variable (*maturity* resolves per industry), the
  context as a variable (*now* and *here* resolve per reading), the
  memory as a variable (the record consulted resolves per world the
  reader stands in) — so the reference still resolves, just
  parameterized by where, when, and against which record it is
  read. And all variables are configurable at all levels: set at
  the fabric, overridden at the platform, the world, the device,
  the field, the reading — the nearest declaration wins, every
  level's setting recorded, no level able to make a variable
  vanish. Anchored, or bound to a declared variable — the two
  lawful endings of a reference chain; a word with neither is a
  dangling pointer in prose, and SourceGround exists to refuse it.
  And the variables are not a concession but the openness itself:
  the more variables you use, the more open your code is — every
  hardcoded value is a decision taken away from whoever runs the
  thing, every declared variable is that decision handed back — and
  more config means more control to the end user, at all levels:
  the same person the box answers to, the same user who owns all
  the faces, now owning the settings too, from the fabric's
  defaults down to their own device's last override. And the
  config is held to the code's own law — Config As Code: every
  setting declared, versioned, diffed, and gated like any other
  source — config that compiles and conforms, recorded in the same
  history, walked through the same gauntlet, because a setting
  that escapes the record is a decision nobody can audit, and the
  charter said it first: code is configuration, and configuration
  is code. Config As Code is a guiding principle of this house —
  not one tool's habit but the rule every tool, platform, and kube
  is built under.

## The rule of the registry

A tool enters by being built or declared here, with its two-word
name and its one-line purpose; a declared tool that is built moves to
built in the same line; and a tool that no longer has an owner
leaves the registry the day its seat empties, because a tool with no
one answering is not a tool but a hazard with a handle.
