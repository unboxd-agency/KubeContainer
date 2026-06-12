# Self-Hosting GitLab — the Individual's Host, Seated Honestly

The founder's order: let's self-host GitLab. The house's own law
answers first, and it is obeyed, not waived:

## The seating (the law applied to ourselves)

The record holds: *open source, self-hosted, is a vulnerability — of
the seating, not the code: the layer runs with its contract seat
vacant... open-source hosting needs an agency.* Self-hosting GitLab
is lawful here for exactly one reason: the agency exists and takes
the seat — unboxd holds the hosting contract for this layer, in its
own name, with the answering priced and the exit real. The
vulnerability the law names is not waived but staffed: breaches at
this layer escalate to a named party (the agency, today resolving to
the founder — the contacts table governs), and the layer is kept
like every kept thing: pinned, patched on cadence, drift-audited.

One more honesty, recorded because honest words demand it: GitLab is
the textbook open-core company — the model this house's own
licensing decision critiques. The pick survives the critique because
of what we take from it: the CE core (MIT) serving git — and git's
own law is the exit: every clone whole, the canon never hostage to
its host. We run GitLab as a serving layer, never a residence; the
open-core risk is contained to a layer we can leave in an afternoon.

## The picks, pinned

- **GitLab CE**, self-hosted, pinned to an exact release at install
  (point of reference, official: https://docs.gitlab.com/install/ —
  collated, not restyled; record the version you stand on the day
  you stand on it).
- **Ground:** the declared stack — the VPS rung (deploy/VPS.md) or
  the estate (deploy/STACK.md); GitLab Omnibus on Ubuntu LTS for the
  single machine, the official Helm chart on the cluster.
- **The registry stays zot** — one custodian per function: GitLab's
  bundled container registry is disabled; zot (already picked, CNCF)
  holds the images. GitLab hosts the record; zot holds the artifacts.
- **No additional DB — the DB stays in the kube.** The founder's
  law, applied here exactly: the stack gains no new database seat.
  GitLab requires PostgreSQL and Redis to exist — and they exist as
  internal organs of GitLab's own kube (the Omnibus bundle keeps
  them inside the box: not provisioned, not exposed, not on the
  stack's deconstruction, no separate layer with a separate
  contract). The deconstruction's storage seats do not grow by one;
  a database that never leaves its kube is not a stack component
  but an implementation detail behind a sealed face — keep the DB
  in the kube, and the kube keeps the DB invisible.
- **The fabric is the DB — the destination, declared.** For the
  fabric's own products the rule is absolute from day one: no
  additional database, ever; the record is the store, the kube
  carries its own state (the way the cluster's record is etcd and
  the box's record is its manifest), and FabricDB — the declared
  seat, durable, indexed, replicated, meta-modeled — is the one
  store the fabric builds when the founder gives the build order.
  GitLab is a guest with its organs sealed inside it; the fabric's
  own organs are the fabric's own, and there is exactly one of
  them.
- **CI:** .gitlab-ci.yml at the repo root — the gauntlet ported:
  same make targets, same verdicts, test/security/eval reports as
  artifacts, the promise check before anything ships. The portability
  claim, no longer a claim.

## The walk (the founder's hand, paste-ready)

    # on the declared Ubuntu ground, per the official docs:
    sudo apt update && sudo apt install -y curl ca-certificates
    curl -s https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.deb.sh | sudo bash
    sudo EXTERNAL_URL="https://gitlab.kubecontainer.xyz" apt install gitlab-ce=<PINNED-VERSION>
    # register a runner for the gauntlet (official: https://docs.gitlab.com/runner/install/)
    # mirror the canon: GitLab is the home, GitHub remains a mirror —
    # push both remotes; every clone is whole either way.

## The exit (always)

`gitlab-ctl backup` + the git remotes themselves: the canon lives in
every clone, the issues and CI history are the host's context (the
NOTICE discipline: decision-bearing context is absorbed into the
canon as it happens, so losing a host loses conversation, never
canon). Uninstall is one package removal; the record walks out the
door whole, as designed from the first principle.
