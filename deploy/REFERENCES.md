# References — Every Component and Seam, With Its Source

The founder's order: source reference URLs to review, for every
component's spec and every seam's contract. Per the law: define term
with point of reference; the point of reference must be real and
currently, actively maintained. Honesty note from the environment
that wrote this: outbound fetches are restricted here, so these are
the leaders' canonical, stable homes, cited for review and verified
on click — each one a living anchor by reputation and the scoreboard,
to be re-pinned to exact versions at deploy time.

## The layers

| # | Component | Source of record |
|---|---|---|
| 0 | Metal / ISA | the vendor's architecture manual (Intel SDM: https://www.intel.com/sdm · Arm ARM: https://developer.arm.com/documentation/ddi0487 · RISC-V: https://riscv.org/technical/specifications/) |
| 1 | Kernel / KVM | https://docs.kernel.org/virt/kvm/ (kernel docs, kernel.org) |
| 2 | OpenStack | https://docs.openstack.org/ · API refs: https://docs.openstack.org/api-ref/ · releases: https://releases.openstack.org/ |
| 2 | Keystone (identity) | https://docs.openstack.org/keystone/latest/ |
| 2 | Cinder (block API) | https://docs.openstack.org/cinder/latest/ |
| 2 | Ceph (storage body) | https://docs.ceph.com/en/latest/ |
| 2 | Rook (cluster-side keeper, CNCF) | https://rook.io/docs/rook/latest/ |
| 3 | Ubuntu Server LTS | https://ubuntu.com/about/release-cycle · https://releases.ubuntu.com/ |
| 3 | cloud-init | https://cloudinit.readthedocs.io/en/latest/ |
| 4 | Kubernetes | https://kubernetes.io/docs/reference/ · source: https://github.com/kubernetes/kubernetes · conformance: https://github.com/cncf/k8s-conformance |
| 4 | k3s (single VPS rung) | https://docs.k3s.io/ |
| 4 | minikube (venv rung) | https://minikube.sigs.k8s.io/docs/ |
| 5 | containerd | https://containerd.io/docs/ · https://github.com/containerd/containerd |
| 6 | KubeContainer / The Metal Kube | this repository: the CRD at config/crd/bases/, the spec at docs/KUBE-SPEC.md, the release at https://github.com/unboxd-agency/KubeContainer/releases/tag/v0.1.0 |
| 7 | The arithmetic kube | deploy/arithmetic-kube.yaml + site/arithmetic.html (in-repo; the formula's law at docs/MEASUREMENT-STANDARD.md) |
| 8 | FabricDB (vacant seat) | the contract declared ahead: docs/SOLID-STATE-DATABASE.md + the four properties in docs/HEADLESS-DELIVERY.md |
| 9–10 | Shared context, agents | docs/PERSONAL.md (terms + flow) · registry/SKELETON.json · docs/AGENT-PLATFORM.md (the ladder, the protocols) |
| 11 | Browser surface | WHATWG HTML: https://html.spec.whatwg.org/ · ECMAScript: https://tc39.es/ecma262/ |
| 12 | The human | no spec; the seat the rest exist to serve |

## The seams

| Seam | Contract | Source of record |
|---|---|---|
| 1↔2 | virtio | https://docs.oasis-open.org/virtio/virtio/ (OASIS) |
| 3↔4 | kubelet host contract | https://kubernetes.io/docs/reference/command-line-tools-reference/kubelet/ · cgroups: https://docs.kernel.org/admin-guide/cgroup-v2.html · systemd: https://systemd.io/ |
| 4↔5 | CRI | https://github.com/kubernetes/cri-api |
| 5↔6 | OCI image / runtime / distribution | https://github.com/opencontainers/image-spec · https://github.com/opencontainers/runtime-spec · https://github.com/opencontainers/distribution-spec |
| storage | CSI | https://github.com/container-storage-interface/spec |
| network | CNI | https://github.com/containernetworking/cni |
| 6↔7 | CRD + CEL admission | https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/ · CEL: https://github.com/google/cel-spec |
| 9↔10 | proto-first wire | https://protobuf.dev/ · gRPC: https://grpc.io/docs/ |
| events | CloudEvents (CNCF) | https://cloudevents.io/ · https://github.com/cloudevents/spec |
| identity | W3C DID | https://www.w3.org/TR/did-core/ |
| things | W3C WoT Thing Description | https://www.w3.org/TR/wot-thing-description11/ |
| ports | IANA service-name & port registry | https://www.iana.org/assignments/service-names-port-numbers/ |
| the scoreboard | LF Insights (founder-pinned) | https://insights.linuxfoundation.org/collection/details/top-open-source-projects · https://insights.linuxfoundation.org/project/k8s |

## The review rule

Review forever means the anchors are re-checked, not re-trusted:
each URL is a living point of reference — still edited, still kept,
still answering — and at deploy time every "latest" above is pinned
to the exact version stood upon, recorded beside the deployment.
A reference that stops resolving is replaced through the registry,
never left dangling — the alias law, applied to the library.
