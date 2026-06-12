# Leap Micro — the host the kube stands on

Status: decision record + walk. Founder's direction, 2026-06-12:
"lets build the kube with this" — openSUSE Leap Micro
(get.opensuse.org/leapmicro). This page records the verdict and the
exact walk; `deploy/deploy.sh` already knows this host.

## The verdict

The kube host is **openSUSE Leap Micro** (currently 6.2), not MicroOS.

- Both are transactional, immutable container hosts: read-only root,
  atomic updates with automatic rollback, SELinux enforcing, built
  for exactly one job — running a container runtime at the edge.
  That is the native fit the geography demands: core to runtime,
  headless edge to edge, operator in env.
- The difference: MicroOS is rolling (Tumbleweed base — the substrate
  moves daily); Leap Micro is fixed releases. The pinning policy that
  governs every tool and image in this stack governs the host too:
  the layer under the kube must be the most boring one. Fixed release
  wins.
- MicroOS remains the rolling option for throwaway dev boxes only,
  never for a host that carries a verdict.

## The walk (three commands, one reboot)

Leap Micro's root filesystem is read-only; packages are layered with
`transactional-update` and arrive on the next boot. So the SELinux
policy for k3s goes in first, then the one-command deploy runs
unchanged:

```sh
sudo transactional-update pkg install k3s-selinux
sudo reboot
curl -sfL https://raw.githubusercontent.com/unboxd-cloud/KubeContainer/main/deploy/deploy.sh | sudo sh
```

`deploy.sh` detects a transactional host itself: if `k3s-selinux` is
not yet layered it refuses with these exact instructions instead of
half-installing; once the host is prepared it proceeds — k3s binary
to `/usr/local` (a writable subvolume), operator from the latest
release, arithmetic kube applied, verdict printed. The same one URL
deploys on Ubuntu and on Leap Micro; the script, not the reader,
carries the difference.
