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
