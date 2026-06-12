# Headless Delivery — Surface, Edge, and Ports

Status: v1, normative. The founder's dictation, recorded and
structured; terms used here are governed by
[FOUNDING-PRINCIPLES.md](FOUNDING-PRINCIPLES.md) and defined in
[AGENT-PLATFORM.md](AGENT-PLATFORM.md) or below. This doctrine extends
the Edge definition (the lexicon) and the brand exegesis
([KUBE-SPEC.md](KUBE-SPEC.md) §9): the soul travels headless; the
body's surface stays sealed.

## 1. The scope rule

The domain defines the scope and the context — never the device. The
device is the surface: the hardened shell, protected against corrosion
and decay (both the metallurgical kind and the security kind), and
nothing more. A contract that names a device has confused the body for
the soul; a contract that names a domain survives every device it ever
touches. This is also the ruling on the brand's width: *Any'Thing'* is
as wide as its domain, and the IoT lineage in KUBE-SPEC §9 stands
because Things are read by domain, not by gadget.

## 2. Single Binary Code

- **Headless delivery** — delivery with no head on the surface: code,
  governance, and security travel as one sealed unit to a port, and
  the surface presents without executing what it presents. The
  surface that delivers will bleed; the headless surface has the
  minimum attack area and does not.
- **Single Binary Code** — the collapse of the riding-alongside zoo —
  multicore tricks, multi-thread scaffolding, sidecars, piggybacks —
  into one binary, in one language common to the hardware and the
  software kernel, native to both. That shared language is the fluid
  for the flexibility of mind, bounded by a hard surface on both
  ends: hard metal below, hard contract above, fluid only between.

The platform capability "no sidecars, substreams, or subgraphs" was
this doctrine's advance guard; Single Binary Code is its full form.

## 3. The OS is the kube

The kernel — the OS — would be the KubeContainer: delivered as a
service, standing on a metal base, multi-mode on a single processor —
a multi-tool operator framework with the orchestrator as its twin.
The pair is the point: operator and orchestrator, each the other's
keeper, giving infinite self-healing ability and infinite scalability
the way two mirrors give infinite depth — not by adding parts but by
facing two loops at each other.

- **Metal base** — the hardware floor the OS-as-kube stands on,
  held by the hardware vendor's contract (the charter's dedication:
  the OS is stable because the hardware vendor holds the contract).
  Bare metal, named as the contract party it is.

## 4. Code is a commodity; code ships in the kube

Now that code has become a commodity, code is never the deliverable —
the kube is. Code must be:

1. *deterministic* — same inputs, same binary, reproducible by
   anyone;
2. *verified* — vulnerabilities scanned, BOM attached, every claim
   about it carrying its verdict;
3. *delivered as an image* — which activates automatically on the
   right context and the right system, and on no other;
4. *never delivered on the surface* — the surface presents, the port
   receives, the kube carries.

## 5. The OS gauntlet

The proto-language the OS chooses is the platform — that is, the
runtime provider. And the OS earns that seat by running the gauntlet
itself: it deconstructs, reconstructs, and tests everything it
receives, integrating all governance and all security into one
automatic, multi-step, single pack — at each step, no loss. What
finally leaves the gauntlet goes to the device-driver manufacturer
with the right BOM and the contract attached. The platform and the
hardware the code runs on are natural partners built on mutual
trust — but trust here is the load-bearing kind: one of them holds
the contract, and the record says which.

## 6. The edge is ingress, not delivery

Innovation happens in the sandboxed cloud environment — never at the
edge. The edge is the ingress to the world: the one-way door where
the world's signals enter, one stream, packet-protected. Delivery is
always at the ports. Said as the rule it is:

- **Edge is ingress** — the edge admits; it never ships. An edge that
  delivers is a wound — it will bleed (leak, corrode, be attacked at
  the moment of exposure). An edge that only ingresses will not. The
  existing Edge definition stands and sharpens: the edge is still the
  surface where representation meets reality — and the traffic across
  it is asymmetric by law: reality flows in at the edge, delivery
  flows out at the ports.

This is the pride of the doctrine: do not deliver on the surface.
The surface would be headless — minimum attack area, nothing
executing, nothing to seize.

## 7. A human at every real-world surface

The world runs multi-mode and multi-operator, with a human at each
real-world surface. Therefore, at every real-world surface, the
platform converts the thing into what that user can use, based on
their skill profile — and the skill ladder is the product loop: try,
test, learn, create versions. Uniform but unique, applied to
delivery: same kube, same contract, a different face for every pair
of hands.

## 8. Branching only at the edge

Branching only happens at the edge — in the desktop, in the venv —
never in the delivered world. From venv to device travels only single
code, single folder. Flexibility lives in the base, not the branch:

- **Multi-base** — one code, many bases: the ability to swap the base
  underneath the single code to check conformance against each — every
  base held in the cloud, every check recorded. The branch explores;
  the base certifies; the code never forks past the edge.

One kube is the communicator with the cloud — a single, contracted
channel, never a mesh of chatty processes.

And the one problem the single communicator exists to end is the old
world's default: different apps each fetching direct info from their
respective clouds. Today every app on the device runs its own wire
to its own vendor — a dozen unrecorded edges on one machine, each
app a self-appointed face sending its own signals (the multi-face
problem of §12 at app granularity), each wire a leak nobody audits
and a bleed nobody staunches. And the indictment is double: they do
not process it efficiently — a dozen apps fetching a dozen streams
through a dozen redundant stacks, the same radio woken, the same
bytes parsed, the same work done N times with no shared skeleton and
no deduplication (the one-custodian-per-fact rule violated in
silicon) — and they have security gaps: every app-owned wire is its
own TLS posture, its own token store, its own update cadence, its
own forgotten endpoint, so the device's real security is the floor
of its worst-maintained app. Inefficiency and insecurity from the
same root: many unrecorded channels where there should be one.

And the stakes are named: because the app actually is the container
of real enterprise data. Behind the friendly face sits the firm's
ledger, the customer list, the contract pipeline — the app is a box
holding real value, but a box without the box doctrine: no declared
contract, no provenance, no registry entry, unboxed by whoever
reaches it, its wires bleeding the most expensive thing on the
device. And individual apps contain dependency chains — each one
embeds its own supply chain of frameworks, SDKs, and transitive
packages, each with a BOM nobody asked for and nobody audits, so a
device with N apps runs N unverified supply chains in parallel,
against the clean-BOM law. And they contain the reporting hierarchy:
who reports to whom, who approves what, who may see which number —
the organization's own skeleton, encoded inside vendor software the
organization neither inspects nor holds the contract for. Data,
dependencies, and the org chart itself: the three crown jewels, all
inside containers that answer to someone else's cloud.

And that is what the AI companies are doing now — the same pattern
at higher stakes: each assistant its own app, each app its own wire
to its own respective cloud, and what flows up that wire is no
longer one app's records but everything the assistant was shown —
the enterprise data, the dependency chains, the reporting hierarchy,
the complete context this doctrine says experiences must be built
from (§10), exported wholesale to whichever vendor's cloud the
friendly face answers to. And the absurdity that gives the pattern away: they are asking the
user to feed the same data again. The enterprise's facts already
live in its systems — entered once, paid for once, owned — and the
assistant's first request is that the human re-type, re-upload,
re-explain what the record already holds: the same data, fed again,
into a second custodian, by hand. It is the deduplication law
broken at the keyboard (one custodian per truth; copies are
staleness waiting to disagree), the user conscripted as the copying
machine, and the copy landing in a cloud the enterprise does not
hold the contract for. A platform that needs the context re-fed
never had lookup — it has ingestion wearing a chat window. The app
problem was a leak; the AI repeat
of it is a transfer of the enterprise's mind. The doctrine's verdict
does not change, it only grows urgent: intelligence comes to the
context through the one recorded channel — the context does not
emigrate to the intelligence.

And the positive office falls out of the indictment: the gap between
app and app is exactly the gap the agent has to fill. Apps do not
talk to each other; their data does not re-feed by human hand; their
clouds do not interconnect in the dark — so the space between them
is empty, and the founding law says what fills empty space in the
fabric: the kube. The agent stands in the app-to-app gap as the one
contracted carrier — reading from one container by right, writing to
the next by contract, deduplicating instead of re-feeding, recording
instead of leaking — turning the device's archipelago of sealed
boxes into one navigable estate without ever merging them. The apps
keep their walls; the agent is the bridge built the lawful way
(trusted insights, actionable intelligence, excellent working
tools); and the gap, once the attack surface, becomes the product.

And the boundary of the office is drawn as sharply as the office:
the agent's gap is app to app — not between device and app, not
between driver and hub, and not between function and function. The
vertical seams belong to the parties who hold their contracts:
device to app is the OS's seam (the kernel, the ports, the native
checks of §10), driver to hub is the hardware partnership's seam
(the manufacturer with the right BOM, §5). And the seam inside the
program belongs to the program: function to function is a call,
inside the one binary of §2, and the standard definition already
ruled it — work an agent would do between two functions is work
that should have been a function call. An agent there is not
governance; it is overhead wearing a contract. An
agent inserting itself vertically would be a shim in someone else's
contract — exactly the riding-alongside this doctrine buried in §2.
The agent works the horizontal: between peers, between boxes,
between containers of value — where there is no contract yet, which
is precisely why an agent with one is needed.

And the root is fed by misplaced faith: the user trusts apps. The human
at the surface extends trust to the friendly face on the screen —
and the app spends that trust on wires the user never sees, signals
the user never signed, clouds the user never chose. The doctrine
does not scold the trust; it re-routes it to a party that can answer
for it: trust the kube — one communicator, user-signed blocks, a
recorded channel, a contract with a name on it — and let apps be
what they are: faces, owned by the user, with no wires of their own
to spend anyone's trust on. The doctrine's answer is absolute:
apps do not talk to clouds; the kube does. Every app speaks locally
to the one communicator; the communicator speaks to the cloud on the
one recorded, signed, packet-protected channel; and the device has
exactly one edge instead of one per installed vendor.

And apps do not talk to each other, either. No app-to-app wire, no
shared memory handshakes, no local mesh re-growing in the shadows of
the cloud rule: if two apps need each other's signal, it travels
app → kube → app, on the recorded channel, under the same contract
as everything else. Point-to-point is the fabric's law *between*
kubes — inside the device, the kube is the only point. An app that
talks sideways is a sidecar by another name, and the sidecar died
in §2.

## 9. The cloud is the capable guard

- **Capable guard** — the cloud's office in this doctrine: not the
  runtime (that is the OS on its metal base) but the guard with the
  capabilities — the multi-stack native build process, the
  distribution, the review agency, the analytics agency — each agency
  with a control plane of its own. The guard's deliverable is the
  contract, delivered to the hardware surface on the right port,
  realtime, as a stream.

## 10. Parallel at delivery only

Parallelization belongs to the delivery surface and nowhere else:
multi-channel streaming, parallel, at the ports going out; one
stream, packet-protected, at the edge coming in. Single code,
multi-mode, across the many bases. Concurrency is a delivery
property, not a code property — which is what Single Binary Code
means in motion.

And the split has its criterion: the delivery is split by the type.
Channels are not balanced by load or sharded by chance — each type
(each kind of skeleton, each declared structure) gets its own
delivery channel to its own right port, so a channel carries one
shape and a port expects one contract. Type is the unit of
parallelism: add a type, gain a channel; retire a type, close one —
and no channel ever has to guess what it is carrying, because the
split itself is the routing and the routing is in the record.

The first and deepest split: content delivery and code delivery are
different. Content goes to faces — converted to the user's skill
profile (§7), streamed multi-channel, presented by the browser-as-
server, consumed where it lands. Code goes to ports — sealed in the
kube, image-form, gauntlet-checked, activated only in its right
context (§4), executed where no face can touch it. Two deliveries,
two channels, two destinations, two trust regimes — and the wound
named in §6 is exactly what happens when a system lets them blur:
content that executes is an attack, and code that is consumed as
content is a leak. The type split keeps them strangers.

And there is package delivery — the third kind, distinct from both:
the delivery of the whole sealed box. Not content (nobody reads a
package), not bare code (nobody executes a package) — the kube
entire: artifact, BOM, contract, and provenance under one lid,
delivered intact, integrity-checked at the port, and unboxed only by
the end user — because unboxing before the end user makes the
contract invalid. Content is consumed, code is activated, packages
are unboxed; three deliveries, three verbs, three channels, and the
box doctrine riding the wire it was always destined for.

And all of it lands on one machine: everything operates with one
device, but at different ports. The split is never hardware — no
second box for code, no appliance for packages — one device, one
surface, one edge, and the separation done entirely by ports: each
delivery type at its own right port, each face turned by its own
user, the whole doctrine running on the single body it protects.
Ports are cheap and recorded; devices are expensive and bleed — so
multiply ports, never devices. And the port rule is absolute: no two
services can run on the same port. One port, one service, one
contract — the one-writer-per-field law made physical. A shared port
is a shared face, and a shared face is the §12 problem reborn:
ambiguous signals, leaking neighbors, nobody answerable. The port is
the smallest face there is, and like every face it has exactly one
owner. And the enforcement costs nothing to build, because the
device OS checks that already: a second bind to a bound port is
refused by the kernel itself, natively, before any policy of ours is
consulted — the rare gate the old world already built exactly where
this doctrine needs it. We adopt the kernel's refusal as our verdict
and add only the record: which service, which port, whose contract —
so that what the OS enforces, the registry can answer for.

But the assignment itself is never the service's to improvise: the
registry must define the port — or the address should be a URL. Two
lawful ways to be reachable, and only two. Either the port is
registered — declared in the registry that defines the contract, the
way the old world's IANA well-known ports were settled by an
international body rather than by squatters — or the service is
addressed by name: a URL, resolved through the resolution layers
(the registry holding the name-to-place mapping, the way the domain
registry holds the mapping and no domain self-declares its IP). A
service that hardcodes its own unregistered number has claimed land
without a deed; the kernel may let it bind, but the record will not
let it answer. Reachability is either registered or resolvable —
never assumed. And which of the two applies is not a style choice:
the content is the difference. What humans consume is addressed by
name — the URL, resolvable, served by the browser-as-server, because
names are the address form a human can hold. What machines activate —
code, sealed packages — is addressed by registered port, because
numbers settled in a registry are the address form a kernel can
enforce. Same device, same edge, two address forms, and the type of
the delivery decides between them: the split by type (§10) reaching
all the way down into how a thing is found.

And content delivery is named for what it really is: content
delivery is all experience. What travels to a face is never mere
data — it is the experience itself: converted to the user's skill
profile (§7), presented at their face, uniform but unique. The
content channel does not ship files that become an experience later;
the delivery *is* the experience, assembled for that human at that
surface — which is why it belongs to names and faces while code
belongs to ports and kernels: experiences are had by someone, and
the address of an experience is the person.

And this is where AI's creative power should shine: the experience
channel is the licensed home of the surreal. The division of labor
already in law — the LLM optimizes creativity, the agent optimizes
outcome — gets its geography here: on the code channel, determinism,
verdicts, and the gauntlet, no creativity in the binary; on the
content channel, create surreal experience — vivid, generative,
astonishing, tuned to the one human at the face — because content
cannot execute (§10) and so the surreal can run free without ever
touching the real. The split by type is what makes the boldness
safe: dream at the faces, because the ports are sober. And the
dreaming is grounded the house way: the surreal is created from
complete context — the experience is generated from the whole
record (who this human is, their skill profile, their world, their
moment — the five dimensions assembled), never from fragments or
guesses. Complete context is what separates a surreal experience
from a hallucinated one: both astonish, but one resolves — every
element traceable back to the context it was woven from. The skeleton
rule again, at the canvas: full structure first, then the art.

## 11. Browser as server

- **Browser as server** — the service built native on the surface:
  the surface's own presenter is the server the user touches, and
  the headless delivery beneath it is the server OS. The browser
  serves; the OS delivers; the surface never executes the payload it
  presents. The resolution of the apparent paradox is the split of
  offices: build the *service* native on the surface, deliver the
  *code* headless at the port — presentation native, execution
  sealed.

## 12. The reversed loop: one machine, many faces

- **The reversed loop** — the industry's loop ran build-outward: the
  device built, pushed artifacts out, and consumed what came back.
  This doctrine reverses it: building flows *in* through the edge —
  declared intent, signals, the branch in the venv — to the sandboxed
  cloud where innovation actually happens; the cloud builds,
  certifies on the many bases, and delivery flows *out* at the ports.
  The user's machine never builds the deliverable; it declares, and
  it receives.
- **Multi-face machine** — the user does not consume content and
  build on the same surface — yet it is the same machine: one
  machine, multiple faces, the consume face and the build face each a
  separate surface of the one body. The faces rule is the identity
  rule applied to hardware (one canonical identity, many profiles;
  uniform but unique): many faces, one machine, every face on the
  manifest.

And the problem the faces rule exists to kill, named exactly: each
face sending a different signal, or leaking information across faces,
is *the* problem. Faces must be coherent (one identity behind all of
them — no face contradicting another) and sealed (nothing crosses
between faces except through the recorded channel). A machine whose
faces disagree is lying to someone; a machine whose faces leak is
bleeding — the surface wound of §6 in its multi-face form.

The separation of duty between faces must happen natively — built
into the machine and its OS, never bolted on as policy from afar —
and it is done by the user himself, at the edge. No remote party
switches a human's faces for him: the human at the real-world surface
(§7) is the one who turns the machine from consume face to build
face, the act is native to the device, and the record keeps the turn.
Separation of duty is a duty — and like every duty in this house, it
belongs to a named party: here, the user, at the edge, by hand. And
the ownership is total: the user controls all the faces. Every face
of the machine answers to the one human at its surface — no face is
the platform's, no face is the vendor's, no face acts on a signal its
user did not stand behind. The platform converts, certifies, guards,
and delivers; the faces it delivers to are owned, turned, and
separated by their human alone.

And one prohibition closes the loop: the surface does not send build
packs — and neither does any face. The prohibition holds at face
granularity: not even the build face ships a built artifact. And no face sends on its own
wire: the sending happens via the edge — every outbound signal from
every face routes through the one ingress point, the single recorded,
packet-protected channel of §6, so there is exactly one place where
leaving is possible and that place keeps the record. What leaves
through the edge is intent and signal — declarations, branches,
telemetry — never built artifacts.

And the unit that travels is specified: very lightweight, and
proto-first — a building block with the entire structure pre-defined
and signed by the user.

- **Proto-first** — schema before content, at the wire: the structure
  of everything that crosses the edge is pre-defined in a published
  schema (the wire-level twin of the kube's Declaration — intent in
  schema, validated before admission), so nothing free-form ever
  travels. A message whose shape was not declared in advance is not
  malformed — it is inadmissible.
- **Signed building block** — the edge's unit of sending: a
  lightweight, proto-first block carrying intent or signal, its whole
  structure known before it is filled, and signed by the user — the
  signature being how "the user controls all the faces" is enforced
  rather than promised: no block leaves a face without its human's
  key on it, and the edge admits nothing unsigned. Light enough to
  stream, strict enough to audit, owned before it moves. And the
  receiving side gets the same discipline: we receive one part, but
  the complete skeleton — a block may fill only one field, yet the
  receiver always holds the entire pre-defined structure it belongs
  to, so partial data never means partial structure. Nothing arrives
  shapeless; the skeleton is whole even when the flesh comes one
  piece at a time — which is what lets streams be partial, audits be
  total, and reassembly be a lookup instead of a guess. And the
  skeleton is not just received — it is used: use the skeleton to
  build the body. The receiver constructs the whole from the
  pre-defined structure, filling it part by arriving part, so the
  body that gets built is the one the user signed for — never an
  improvisation around fragments. The skeleton is the contract of
  assembly; the parts are deliveries against it; the body is the
  contract, kept.

And the case of the new skeleton is ruled: if the user is defining a
new skeleton — a structure that does not yet exist — it goes through
the innovation cycle. Filling an existing skeleton is sending;
defining a new one is innovating, and innovation happens in the
sandboxed cloud (§6), never on the wire. The new skeleton is born as
a branch at the edge (§8), travels as signed intent through the one
channel, is tried, tested, and certified in the sandbox by the
capable guard (§9) — and certification alone does not finish it: it
has to earn enough human contributions to be discoverable. A skeleton
enters the catalog by adoption, not by declaration — humans filling
it, attesting it, building bodies on it — so discoverability is
standing earned in use, the reputation rule applied to structures.
Only then, published and registered (the registry defining the
contract), does it become a pre-defined structure the edge will admit
fillings of. No skeleton is improvised into existence mid-stream: the
wire carries only structures that already exist, and the path by
which a structure comes to exist is the same gauntlet everything else
walks — plus the one thing a gauntlet cannot grant, which is other
people.

Build packs exist only cloud-side, born in the sandbox, certified by
the guard, delivered at the ports. A surface or a face that emits
build packs has reversed the reversed loop — it is the old world's
wound reopened.

## The doctrine in one breath

Domain defines scope; device is only the surface. Code is commodity
and ships sealed in the kube, deterministic and BOM-verified, as an
image that wakes only in its right context. The OS is the kube on a
metal base, twinned with its orchestrator for infinite healing and
scale. The edge admits and never ships; the ports ship and never
admit; the surface presents and never executes. Branch in the venv,
certify on many bases, deliver one binary. The cloud guards, the
hardware partners, the contract is held — and the record, as always,
says by whom.
