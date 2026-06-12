# The Control Panel — Podman as Design Reference

The founder's pick: Podman as the design reference for the panel.
Podman is chosen for what it proves, not merely what it runs —
daemonless, rootless, fork/exec, open — and each property is a panel
law already in this house's doctrine:

| Podman property | Why it is the reference | The panel law it grounds |
|---|---|---|
| Daemonless | no always-on root daemon mediating everything (the Docker model, struck) | the panel is not a privileged server that owns the box; it issues declarations and exits — command as code, no resident master |
| Rootless | containers run as the invoking user, not root | the user owns all the faces: the panel acts as the logged-in principal, least privilege, never a god-daemon |
| fork/exec, no broker | each container a child process, directly owned | one custodian per resource: the panel launches and the OS owns; no broker holding everyone's processes hostage |
| Pods (the namesake) | groups of containers sharing a context | the kube: a declared group with one record, one face |
| Docker-compatible CLI/API | speaks the standard, invents no lock-in | the panel speaks the conformance contract (kube CRD, OCI, k0s) — adapters, not proprietary wires |
| systemd/Quadlet units | containers declared as units, not imperatively run | Config As Code: every panel action is a declaration on the record, idempotent, diffable |

## What the panel is (v0, the edge-as-UI-ref)

One screen, served surface-native (site/panel.html is the seed):

1. Login — the principal arrives; the panel acts as them, never above them (rootless).
2. Point at a server — the declared ground (the VPS, k0s on it).
3. One resident per row, one button each — GitLab, zot, the site kube, an agent: each row a declaration (the skeleton already in registry/SKELETON.json and deploy/*.yaml), the button an apply, no imperative script.
4. Live verdicts — each row shows its state from the cluster's own truth (kubectl get under the hood), the way Podman shows container status — observed, not asserted.
5. No resident daemon — the panel renders, issues the declaration, and the loop (the operator) keeps it; the panel can close and the kube stays kept. Daemonless, exactly as Podman taught.

## The boundary (honest, recorded)

The panel issues declarations and reads verdicts; it never becomes
the keeper. The operator keeps, k0s runs, the panel only declares and
observes — so closing the panel changes nothing that was converged.
That is the daemonless law: the control plane is the cluster, not the
panel; the panel is a face, and faces present, they do not rule.

## The market gap (the opening)

The founder's finding, recorded as the reason to build: a pod-native
panel is not there. The kube-native control surfaces that exist are
thin — Headlamp has very few options, the dashboards assume you
already speak kubectl, and the easy panels (Coolify and kin) are
Docker-native, not pod-native, and bring the daemon this house
struck. So the seat is open: an easy, pod-native, daemonless panel —
Podman's model at the cluster scale, the kube as the unit, the
operator as the keeper. The gap is the product; this is why we build
rather than adopt.

## The trick — one app, many envs, the operator may differ

Proven by hack/operatorpoc (rehearsal chamber): one application
declaration (arithmetic) moves through the gates into two
environments; the operator keeps it to each env's own policy (dev: 1
replica, prod: 3). Same app, two envs, two views — and the operator
need not be the same in each. That is the panel's central trick: the
application is the constant that travels; the operator is the variable
that keeps. The founder's design, recorded fragment by fragment:

- Podman Desktop to kube-cloud — the panel spans the desk to the
  cloud in one app, one browser: the local (Podman-model, the desk)
  and the cluster (the kube-cloud) both visible in one surface, the
  way Podman Desktop shows local containers — but reaching all the
  way to the cloud's pods, not stopping at the laptop.
- One nav menu — everything is one navigation: no separate
  consoles per env, per server, per registry. One menu, the whole
  estate.
- Each view tells its env clearly — every panel, row, and
  resource is labeled with its environment, unmistakably: which env,
  always on the face, so two views of one app are never confused.
- Multi-server, single operator — one operator can keep an app
  across many servers; the app can sit at any level (desk, edge,
  estate). The operator is the constant keeper; the servers and
  levels are where it keeps.
- Past history preserved — every env, every move through the
  gates, keeps its history: the record is append-only, so the panel
  shows not just the present view but the journey — what each env
  was, when it changed, who moved it (the manifest, on the face).

## The path, and the gap

The path Podman creates — Podman → zot → kind → the Kubernetes panel
(the kube dashboard): the desk container engine to the open registry
to the local cluster to the cloud panel, one continuous road from
laptop to estate. And the gap the founder names to fill gradually:
Podman has no hub like Docker Hub — no first-class, trusted,
discoverable registry-with-a-face the way Docker Hub is for Docker.
That hub is the registry strategy this house already declared (the
registry defines the contract; zot the store; the panel the face) —
the gap is the opening, and the panel is where it gets filled.

## Crossplane as reference — the gates as another view

Point of reference, founder-supplied (CNCF, incubating):
https://www.cncf.io/projects/crossplane/ — Crossplane makes the
cluster the control plane for everything (infra, services, apps) via
CRDs and compositions; the panel takes from it the principle that one
declarative control plane governs every layer, not just workloads.

And the build/deploy gates are another view in the panel: alongside
the env views (dev, prod) sits the pipeline view — the application
moving through its gates (build, test, security, promise, deploy),
each gate's verdict on the face, the same app tracked from source to
served. Two axes in one nav: across environments (where it runs) and
along the gates (how far it has passed).

The path is configurable as a flow: the gates an app moves through
are not fixed but declared — a flow the operator (or its env-specific
keeper) walks, editable in the panel, recorded like every
declaration. Different apps, different flows; the panel renders
whichever flow the app declares.

And the flow engine can be Temporal: durable, replayable workflow
execution (the gate-walk as a workflow that survives restart, retries
on failure, preserves its history) — a candidate seat for the flow
runtime, measured against the house's own loop-and-record discipline,
recorded as a candidate the way every pick is.

## The watcher and the reconciler — two seats, not one

The founder splits the flow's two offices cleanly:

- A watcher observes that the flow completes — a monitoring seat
  (Cortex named as the candidate: CNCF, the horizontally-scalable
  metrics/observability project — it watches and alerts that the
  flow reached its end, the gates all green, the deploy served). The
  watcher sees; it does not act.
- The kube is the reconciler — the keeper that closes the distance
  to the declaration, the office this whole house was built on. The
  flow declares the desired path; the kube reconciles toward it; the
  watcher confirms it arrived. Three roles, never confused: the flow
  declares, the kube reconciles, the watcher witnesses.

This is the loop's own anatomy mapped onto the panel: observe (the
watcher), compare and act (the reconciler/kube), record (the
append-only history). The panel shows all three at once — the
declared flow, the reconciling kube, the watcher's verdict that it
completed — so a human reads, in one nav, what was wanted, what is
being kept, and whether it landed.

## The flow builder — Orkes/Conductor

The flow must be built before it is walked, and the founder names the
builder's seat: Orkes (Conductor) — the workflow orchestration
platform (Netflix Conductor's lineage, open core) as the candidate
flow builder: the visual/declarative authoring of the gate-path
itself, the place a flow is composed (which gates, in what order,
with what branches) before the kube reconciles it and the watcher
confirms it. The full seating of the flow's lifecycle, candidates
named:

| Office | What it does | Candidate (recorded, measured against the house's loop-and-record law) |
|---|---|---|
| Flow builder | compose the gate-path | Orkes / Conductor |
| Flow engine | execute it durably | Temporal |
| Reconciler | keep it to declaration | the kube (this house) |
| Watcher | witness completion | Cortex |

Built (Orkes), executed (Temporal), kept (the kube), witnessed
(Cortex) — four offices, the panel showing the one flow across all of
them. Each candidate enters by conformance and is swappable; the kube
is the only seat this house fills itself, because reconciliation is
its founding craft.

## Two decisions, delegated and made

The runtime under the panel: k3s. The founder left the call ("could
be k3s — up to you") and the reasons decide it: the panel stack needs
ingress, storage, and Helm-handling out of the box, and k3s ships all
three built in (Traefik, local-path, the Helm controller) where k0s
ships bare. k0s stays the leaner candidate for headless edge nodes;
k3s takes the home — the panel's ground needs the batteries, the
edge does not. Both conformant; the declaration walks both unchanged.

Crossplane: useful — yes, at one seat, without unseating OpenTofu.
Crossplane is the kube's own model applied to infrastructure
(continuous reconciliation of external resources as CRDs — DNS,
instances, buckets kept, not just applied), which fits this house's
loop-law better than plan/apply once a long-lived cluster exists. But
it cannot bootstrap itself: you need a cluster before Crossplane can
run. So the seating is sequential, no overlap: OpenTofu for day-0
(the ground bootstrapped, plan/apply), Crossplane for steady state
(the ground reconciled from inside the cluster, drift reverted like
any kube child). Bootstrap by tofu, keep by Crossplane, both
declared, the exit real at each.

## Three more references, seated

- Ubuntu Landscape — the fleet-management reference (Canonical's
  panel for estates of Ubuntu machines: enrollment, patching,
  audit, one console for thousands of boxes). What the panel takes
  from it: machines enroll into the estate and are kept on cadence —
  the fleet view our panel needs for many Metal Kubes, with the
  difference that Landscape manages hosts imperatively while ours
  declares and reconciles them. Reference: https://ubuntu.com/landscape
- Operator Framework — the operator-pattern toolchain proper
  (Operator SDK, OLM lineage): the reference for how operators are
  built, packaged, and lifecycled at scale. This house already
  builds on Kubebuilder (the same family); the framework is the
  reference for the panel's "every resident is an operator-kept
  app" rule. Reference: https://operatorframework.io/
- Multi/hybrid-cloud frameworks — the estate spans clouds, and the
  panel must read them as one: the reference class is the
  multi-cluster/multi-cloud control planes (Crossplane already
  seated for resources; the fleet/cluster layer candidates: Open
  Cluster Management, Karmada — both CNCF) — one nav, many clouds,
  each env labeled clearly, the same trick at cloud scale: the app
  constant, the ground variable.
