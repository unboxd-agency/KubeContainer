# The Agent-Engineering Stack — Complete, but Lightweight

The founder's ruling: a complete stack for agent engineering, kept
lightweight. Complete = every concern of agent engineering (the six
of AGENT-PLATFORM.md §3) has a named seat; lightweight = each seat is
the smallest leader that fills it, nothing runs that isn't earning,
the whole thing standing on one VPS. Weight is the enemy of the edge;
completeness is the enemy of the toy. Both held — balance, not
maximization.

| Agent-engineering concern | The seat | The lightweight pick | Weight |
|---|---|---|---|
| Goal & contract design | the CRD + CEL; StructuredInstructions | in-repo, already built | ~0 (a binary, a schema) |
| Tool/actuator design | the operator framework (Kubebuilder) | the kube itself | the operator pod |
| Loop & orchestration | the reconciler; the platform-bound orchestrator | KubeContainer | one controller |
| Memory & context | the record — no additional DB; FabricDB when built | k0s etcd now (the kube carries its state) | none added |
| Guardrails & permissions | k0s RBAC; the agent contract terms 1–10 | k0s built-in | none added |
| Evaluation & observability | the eval registry; CodeCompiler; the gauntlet | in-repo scripts + binaries | run on demand |
| Runtime | k0s (single binary, own containerd) | k0s | ~1 binary |
| Registry | zot | zot (CNCF, minimal) | one pod |
| Source + CI | GitLab CE + runner | GitLab (the one heavy guest) | the one big tenant |
| Ground | OpenTofu-declared VPS | one VPS | the metal |
| Build | Cloud Native Buildpacks | pack | on demand |
| Surface | the browser; the site kube | nginx kube | one pod |

The lightweight discipline, stated so it cannot drift into bloat:

- **No additional database** — the record is the store; k0s etcd holds
  cluster state, the kube carries its own; FabricDB is the only store
  ever added, and only when built.
- **One custodian per function** — one runtime (k0s containerd, Docker
  struck), one registry (zot), one source host (GitLab), one image
  builder (pack). No second of anything.
- **On-demand, not always-on, for the desk tools** — CodeCompiler,
  the eval harness, pack, the gauntlet run when invoked; they are not
  resident daemons eating the box while idle.
- **GitLab is the one heavy resident, and it earns it** — source,
  CI, runner in one. If even that is too heavy for the chosen metal,
  the lawful lighter swap is Forgejo (the candidate already named) —
  recorded so the choice is the founder's, not a silent downgrade.

Completeness check: every concern has a seat, no seat is empty but
FabricDB (declared, awaiting the build order), and nothing on the
list is there for show. Lightweight check: one VPS carries it; the
only always-on residents are k0s, the operator, zot, the site, and
GitLab — five, each earning its memory. Complete and light is a
balance, and this is where it resolves.
