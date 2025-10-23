<!-- BEGIN MUNGE: GENERATED_TOC -->

- [v1.35.0-alpha.2](#v1350-alpha2)
  - [Downloads for v1.35.0-alpha.2](#downloads-for-v1350-alpha2)
    - [Source Code](#source-code)
    - [Client Binaries](#client-binaries)
    - [Server Binaries](#server-binaries)
    - [Node Binaries](#node-binaries)
    - [Container Images](#container-images)
  - [Changelog since v1.35.0-alpha.1](#changelog-since-v1350-alpha1)
  - [Changes by Kind](#changes-by-kind)
    - [Deprecation](#deprecation)
    - [API Change](#api-change)
    - [Feature](#feature)
    - [Documentation](#documentation)
    - [Bug or Regression](#bug-or-regression)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake)
  - [Dependencies](#dependencies)
    - [Added](#added)
    - [Changed](#changed)
    - [Removed](#removed)
- [v1.35.0-alpha.1](#v1350-alpha1)
  - [Downloads for v1.35.0-alpha.1](#downloads-for-v1350-alpha1)
    - [Source Code](#source-code-1)
    - [Client Binaries](#client-binaries-1)
    - [Server Binaries](#server-binaries-1)
    - [Node Binaries](#node-binaries-1)
    - [Container Images](#container-images-1)
  - [Changelog since v1.34.0](#changelog-since-v1340)
  - [Changes by Kind](#changes-by-kind-1)
    - [API Change](#api-change-1)
    - [Feature](#feature-1)
    - [Bug or Regression](#bug-or-regression-1)
    - [Other (Cleanup or Flake)](#other-cleanup-or-flake-1)
  - [Dependencies](#dependencies-1)
    - [Added](#added-1)
    - [Changed](#changed-1)
    - [Removed](#removed-1)

<!-- END MUNGE: GENERATED_TOC -->

# v1.35.0-alpha.2


## Downloads for v1.35.0-alpha.2



### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes.tar.gz) | acba342356249738a81bf6bc6de95e4a30097fdd0ebe956b8cd8a2b0715e3161930f7408bd3b1ca1e05c07de4359485cf887b278987366efef3caf9024e80c6d
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-src.tar.gz) | 6e9f58180f53e57ae6b462d4ab3a13f7cafc9bb9802f8af3254e9f3c78b9883103972dced5dd0796c9c8e4176fd8557754981a63fc4b5eb4fb0d07838027ac70

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-darwin-amd64.tar.gz) | cb54b9aa876b327915048fa3d9a152abcde442d60cee750566339335b19c668f1d440f1dd79409137e7ee5d7e32e2d3c6e8b3fcaf7f4932b19508b483e3d4172
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-darwin-arm64.tar.gz) | 600f2922a818c9c750269695b9158892fcfdd1dd1311701033f93b396689c7d4625c24880598ea36ca3d1ff76be53dcdff911a96d8f337ec93847e340639a92b
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-linux-386.tar.gz) | bcc1d2c3b5577b22636b7c9aa515fb9944e586d5ae657e066e204388992bb1e9c94dd54ecc7feaaafe46c89943e5500366a26dac11ee2eb32ea3106daf1da51b
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-linux-amd64.tar.gz) | 4250b2063c70cd69b49a50a4a416a9bd5a4e7734ed8b9ccc1081ed12e23c30018c2be9dc377100eb14823bab26aa33670e92d7ba38588a2a0ca011c3d63ecbf5
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-linux-arm.tar.gz) | 203825398afd6c697ac2fc13b126d7419b1c108362e6bb8a27eddef57e2845dd02735e4c48a5c2aa813f9e0ce24ee97ae94360cf50a9197fba53ba3ac736a50e
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-linux-arm64.tar.gz) | b6669df8c4e096d7ca435bfa481823b74e131907433fc7b7dbf6e6a699f2905a60c98e3c23c9321462ae3afdd707ffea3acf473a13905e63203cebefa80028c2
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-linux-ppc64le.tar.gz) | 37b0ce0d3dfa8dcd2222c63b6572e32ad1a7f07d4164de886b3eca04d4c655a3cc07786090eb24cc20f0bf641cae2efba7ab3c3cd2da5536575571db31aa89da
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-linux-s390x.tar.gz) | 3d7054c4b8d18501b535b0cd070bab316b7393bcb575fc869e2fde7190044b15a42e32dbea6aee64aef933ec1d8c7c11581c61bcb4829e710b26971b133180c3
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-windows-386.tar.gz) | ae011d1aa7b41160d50b9cd9bc4fe2890bbc2ce2f2b6c63695ae20f36e93cbf189c32deafc0d99c46532917ba291f40965cd4038edcb5bb3a27cd66974dba539
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-windows-amd64.tar.gz) | bcff5d410cea98ab7e83a66c545b4322d17055ed0b3c7acb110a757e6f0ee55aadfc0174c8c641511ac832024af5b2660f4e2be5c3076a12e0b862aa55a1d02f
[kubernetes-client-windows-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-client-windows-arm64.tar.gz) | 4fb60b2747b500f1139f590e436318fabdd692fc7d2de27be9667c1e5f9af3a6a67796fcd3c69b92e225a04ae92292d715c4c5e1a1437f1723e5bd16d30e5c59

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-server-linux-amd64.tar.gz) | b29aaf01ad35edf7d24ac2a1d493c28a65941fd9f490bbcaeecfc418b1e26060f90e1677353ace6229ae1b8416f5080e116fcfb90732a7aab094761d9f1dadbe
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-server-linux-arm64.tar.gz) | fc8aeaec77c22d1cb777d9d626f6cbdc0bd178a29d1c125305592d4b40680c51d30d6ffeeb5754abd029c56ed2f49462a85799939f7ffc343f4816da3b9a2d20
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-server-linux-ppc64le.tar.gz) | 61a70fa842e8afff5bdf3ab45a85b0bae183eb0e3910c440ca21520d3f03e0ee66ffbdd8b335b0d9ccfd2c77fb1a82e0f1480a267da6e6c10255c87465b12965
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-server-linux-s390x.tar.gz) | 9275068613ddfaa163bbfaed5ae0c69dc2ca2031b3f42f990ed42995b14e7d5ed1bd5d49d3c3b7e95f7024a4434cb3518e05240b82d508f2be2c6d4971d3ab43

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-node-linux-amd64.tar.gz) | 2200656cfe27817ec8bfc67564fba75c0afb582c75b2ce37734dff1c2757d142d45a24695c7e898b4663362f4058ca0ae8399ee485883833498cac9867caccdc
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-node-linux-arm64.tar.gz) | fe087958b7b7b7197132473508deffa90a740fffc2bf7a06c9a7c7df029394fd27a307efc6bb8003c6f95d9013f57ed577ec4a777881c44acba26a1ebc918ae5
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-node-linux-ppc64le.tar.gz) | 621a7d6f7f3fcc382922f0912a5dd3f9587ec15992c65be806a18e4b3254895d42dc78ac2b1aab10a16dee1227ca315c1b5b35c27b29946c1337d548b799ddc7
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-node-linux-s390x.tar.gz) | 16a9074fba6db7ef45b38ed5ea05ae9cd47a6388b01cfa551377de5bc1b720df3507b8bda974c1220d8a82394902192a4013754026f2d71732bb480743862c05
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.2/kubernetes-node-windows-amd64.tar.gz) | 6a893a6a4ad7f664fea7536f22bd27d4f874e7568658f5863e156c290b4afee2ef566797c6e3dae86b4d706f219b8169da5e03ef61e565b7a4ba2123a6b43c5c

### Container Images

All container images are available as manifest lists and support the described
architectures. It is also possible to pull a specific architecture directly by
adding the "-$ARCH" suffix  to the container image name.

name | architectures
---- | -------------
[registry.k8s.io/conformance:v1.35.0-alpha.2](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-s390x)
[registry.k8s.io/kube-apiserver:v1.35.0-alpha.2](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-s390x)
[registry.k8s.io/kube-controller-manager:v1.35.0-alpha.2](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-s390x)
[registry.k8s.io/kube-proxy:v1.35.0-alpha.2](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-s390x)
[registry.k8s.io/kube-scheduler:v1.35.0-alpha.2](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-s390x)
[registry.k8s.io/kubectl:v1.35.0-alpha.2](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-s390x)

## Changelog since v1.35.0-alpha.1

## Changes by Kind

### Deprecation

- FailCgroupV1 will be set to true from 1.35. 
  This means that nodes will not start on a cgroup v1 in our default behavior. 
  This is putting cgroup v1 into a deprecated state. ([#134298](https://github.com/kubernetes/kubernetes/pull/134298), [@kannon92](https://github.com/kannon92)) [SIG Node]
- Mark ipvs mode in kube-proxy as deprecated. ipvs mode in kube-proxy is deprecated and will be removed in a future version of Kubernetes. Users are encouraged to move to nftables. ([#134539](https://github.com/kubernetes/kubernetes/pull/134539), [@adrianmoisey](https://github.com/adrianmoisey)) [SIG Network]

### API Change

- Kube-apiserver: fix a possible panic validating a custom resource whose CustomResourceDefinition indicates a status subresource exists, but which does not define a `status` property in the `openAPIV3Schema` ([#133721](https://github.com/kubernetes/kubernetes/pull/133721), [@fusida](https://github.com/fusida)) [SIG API Machinery, Apps, Architecture, Auth, Autoscaling, CLI, Cloud Provider, Cluster Lifecycle, Etcd, Instrumentation, Network, Node, Release, Scheduling, Storage and Testing]
- Kubernetes API Go types removed runtime use of the github.com/gogo/protobuf library, and are no longer registered into the global gogo type registry. Kubernetes API Go types were not suitable for use with the google.golang.org/protobuf library, and no longer implement `ProtoMessage()` by default to avoid accidental incompatible use. If removal of these marker methods impacts your use, it can be re-enabled for one more release with a `kubernetes_protomessage_one_more_release` build tag, but will be removed in 1.36. ([#134256](https://github.com/kubernetes/kubernetes/pull/134256), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Apps, Architecture, Auth, CLI, Cluster Lifecycle, Instrumentation, Network, Node, Scheduling and Storage]
- Promoted HPA configurable tolerance to beta. The `HPAConfigurableTolerance` feature gate is now enabled by default. ([#133128](https://github.com/kubernetes/kubernetes/pull/133128), [@jm-franc](https://github.com/jm-franc)) [SIG API Machinery and Autoscaling]
- The MaxUnavailableStatefulSet feature is now beta and enabled by default. ([#133153](https://github.com/kubernetes/kubernetes/pull/133153), [@helayoty](https://github.com/helayoty)) [SIG API Machinery and Apps]

### Feature

- Enable the feature gate `ContainerRestartRules` by default. The ContainerRestartRules feature is promoted to beta. Fixing a bug in this feature that caused probes continue to run even if the container has terminated and is not restartable. ([#134631](https://github.com/kubernetes/kubernetes/pull/134631), [@yuanwang04](https://github.com/yuanwang04)) [SIG Node]
- Kube-apiserver: the subresources `pods/exec`, `pods/attach`, and `pods/portforward` now require `create` permission for both SPDY and Websocket API requests. Previously, SPDY requests required `create` permission, but Websocket requests only required `get` permission. This change is gated by the `AuthorizePodWebsocketUpgradeCreatePermission` feature-gate, which is enabled by default.
  
  Before upgrading to 1.35, ensure any custom ClusterRoles and Roles intended to grant `pods/exec`, `pods/attach`, or `pods/portforward` permission include the `create` verb. ([#134577](https://github.com/kubernetes/kubernetes/pull/134577), [@seans3](https://github.com/seans3)) [SIG API Machinery, Auth, Node and Testing]
- Kubeadm: print the errors during retires related to the WaitForAllControlPlaneComponents functionality at verbosity level 5. ([#134433](https://github.com/kubernetes/kubernetes/pull/134433), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubernetes is now built using Go 1.25.3 ([#134611](https://github.com/kubernetes/kubernetes/pull/134611), [@cpanato](https://github.com/cpanato)) [SIG Architecture, Cloud Provider, Etcd, Release, Storage and Testing]
- Locked the (generally available) feature gate `ExecProbeTimeout` to true. ([#134635](https://github.com/kubernetes/kubernetes/pull/134635), [@vivzbansal](https://github.com/vivzbansal)) [SIG Node and Testing]
- Promoted the `HostnameOverride` feature gate to beta and is enabled by default. ([#134729](https://github.com/kubernetes/kubernetes/pull/134729), [@HirazawaUi](https://github.com/HirazawaUi)) [SIG Network and Node]

### Documentation

- Kubectl describe, get, drain and events have the ability to set chunk-size using --chunk-size flag, which is now officially stable. ([#134481](https://github.com/kubernetes/kubernetes/pull/134481), [@soltysh](https://github.com/soltysh)) [SIG CLI]

### Bug or Regression

- DRA API: the "tolerations" field in exact and sub requests now gets dropped properly when the DRADeviceTaints API is disabled. ([#132927](https://github.com/kubernetes/kubernetes/pull/132927), [@pohly](https://github.com/pohly))
- DRA Device Taints: tolerating a NoExecute did not work because the scheduler did not inform the eviction controller about the toleration, so the scheduled pod got evicted almost immediately. ([#134479](https://github.com/kubernetes/kubernetes/pull/134479), [@pohly](https://github.com/pohly)) [SIG Apps, Node, Scheduling and Testing]
- Endpoints/endpointslice controllers perform much better when there are a large number of services in a single namespace ([#134739](https://github.com/kubernetes/kubernetes/pull/134739), [@shyamjvs](https://github.com/shyamjvs)) [SIG Apps and Network]
- Fixed a bug that prevents schedule next pod when using DRAConsumableCapacity feature. (#133705, @sunya-ch) ([#133706](https://github.com/kubernetes/kubernetes/pull/133706), [@sunya-ch](https://github.com/sunya-ch)) [SIG Node]
- Fixed a bug where 64 bit IPv6 ServiceCIDRs allocated addresses outside the subnet range. ([#134193](https://github.com/kubernetes/kubernetes/pull/134193), [@hoskeri](https://github.com/hoskeri)) [SIG Network]
- Fixed a startup probe race condition that caused main containers to remain stuck in "Initializing" state when sidecar containers with startup probes failed initially but succeeded on restart in pods with restartPolicy=Never. ([#133072](https://github.com/kubernetes/kubernetes/pull/133072), [@AadiDev005](https://github.com/AadiDev005)) [SIG Node and Testing]
- Kube-apiserver: when --requestheader-client-ca-file and --client-ca-file contain overlapping certificates, --requestheader-allowed-names must be specified to ensure regular client certificates cannot set authenticating proxy headers for arbitrary users ([#131411](https://github.com/kubernetes/kubernetes/pull/131411), [@ballista01](https://github.com/ballista01)) [SIG API Machinery, Auth and Security]
- Kube-controller-manager: Resolves potential issues handling pods with incorrect uids in their ownerReference ([#134654](https://github.com/kubernetes/kubernetes/pull/134654), [@liggitt](https://github.com/liggitt)) [SIG Apps]
- Kubeadm: avoid panicing if the user has malformed the kubeconfig in the cluster-info config map to not include a valid current context. Include proper validation at the appropriate locations and throw errors instead. ([#134715](https://github.com/kubernetes/kubernetes/pull/134715), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: fixes a preflight check that can fail hostname construction in IPV6 setups ([#134588](https://github.com/kubernetes/kubernetes/pull/134588), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Auth, Cloud Provider, Cluster Lifecycle and Testing]
- Legacy watch calls (RV = 0 or unset) that generate init-events weigh higher in APF seat usage now. Properly accounting for their cost protects the API server from CPU overload. Users might see increased throttling of such calls as a result. ([#134601](https://github.com/kubernetes/kubernetes/pull/134601), [@shyamjvs](https://github.com/shyamjvs)) [SIG API Machinery]
- Prevent a segfault occurring when updating deeply nested JSON fields ([#134381](https://github.com/kubernetes/kubernetes/pull/134381), [@kon-angelo](https://github.com/kon-angelo)) [SIG API Machinery and CLI]
- The kubelet now honors the configuration userNamespaces.idsPerPod. Before it was ignored. ([#133373](https://github.com/kubernetes/kubernetes/pull/133373), [@AkihiroSuda](https://github.com/AkihiroSuda)) [SIG Node and Testing]

### Other (Cleanup or Flake)

- Building Kubernetes is now implemented by running a pre-built container image directly, without running rsyncd, and is substantially simplified. ([#134510](https://github.com/kubernetes/kubernetes/pull/134510), [@BenTheElder](https://github.com/BenTheElder)) [SIG Release and Testing]
- CPU Manager static policy option `strict-cpu-reservation` moved to the GA version ([#134388](https://github.com/kubernetes/kubernetes/pull/134388), [@psasnal](https://github.com/psasnal)) [SIG Node]
- Dropped support for policy/v1beta1 PodDisruptionBudget in kubectl ([#134685](https://github.com/kubernetes/kubernetes/pull/134685), [@scaliby](https://github.com/scaliby)) [SIG CLI]
- Kubeadm: stoped applying the --pod-infra-container-image flag for the kubelet. The flag has been deprecated and no longer served a purpose in the kubelet as the logic was migrated to CRI. During upgrade, kubeadm will attempt to remove the flag from the file /var/lib/kubelet/kubeadm-flags.env. ([#133778](https://github.com/kubernetes/kubernetes/pull/133778), [@carlory](https://github.com/carlory)) [SIG Cloud Provider and Cluster Lifecycle]
- Kubeadm: updated the supported etcd version to v3.5.23 for supported control plane versions v1.31, v1.32, and v1.33. ([#134692](https://github.com/kubernetes/kubernetes/pull/134692), [@joshjms](https://github.com/joshjms)) [SIG Cluster Lifecycle and Etcd]
- Kubeadm: updated the supported etcd version to v3.5.24 for supported control plane versions v1.32, v1.33, and v1.34. ([#134779](https://github.com/kubernetes/kubernetes/pull/134779), [@joshjms](https://github.com/joshjms)) [SIG API Machinery, Cloud Provider, Cluster Lifecycle, Etcd and Testing]
- Kubernetes is now built with go 1.25.3 ([#134598](https://github.com/kubernetes/kubernetes/pull/134598), [@BenTheElder](https://github.com/BenTheElder)) [SIG Release]
- Promote the Topology Manager policy option max-allowable-numa-nodes to GA ([#134614](https://github.com/kubernetes/kubernetes/pull/134614), [@ffromani](https://github.com/ffromani)) [SIG Node]
- Rsync is no longer required to build kubernetes. ([#134656](https://github.com/kubernetes/kubernetes/pull/134656), [@BenTheElder](https://github.com/BenTheElder)) [SIG Release and Testing]
- The storage.k8s.io/v1alpha1 VolumeAttributesClass API is no longer served in 1.35 ([#134625](https://github.com/kubernetes/kubernetes/pull/134625), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Etcd, Storage and Testing]

## Dependencies

### Added
_Nothing has changed._

### Changed
_Nothing has changed._

### Removed
_Nothing has changed._



# v1.35.0-alpha.1


## Downloads for v1.35.0-alpha.1



### Source Code

filename | sha512 hash
-------- | -----------
[kubernetes.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes.tar.gz) | 1d6fb6a4c7f82fe04e56757b733c3fc4aac652f8c2113e79ddce83b6cbe0179404147b35ddbc18e1b60eb802acb3f6d884599fd573f3d16f0558ef7ddfb8aae2
[kubernetes-src.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-src.tar.gz) | 364788bac4d405ac6180fe3cb7e3d847e7960fcb0532146b105270aeac2624ade2ff87370c5aa8f768eda07fd28e5e75f73afbdf9cc1b786827a0e123bdea561

### Client Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-client-darwin-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-darwin-amd64.tar.gz) | 1ba40849b104851d922bce32dc9306004e9b95cfadeff9ecfb65f779892009f9a70878b8efe96159088b1ad8c700bf19e58d68416dfdff7853660e6074dd3752
[kubernetes-client-darwin-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-darwin-arm64.tar.gz) | 0c939895ad2d53f57e9137774eed99cbfbfa5f15d4276f5f55c4ea40b922a5f37ab375a065fdd330f5a1ddf452896f2a13621075b049f382ab42a65ea1085dac
[kubernetes-client-linux-386.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-linux-386.tar.gz) | e98a9b2d5f1c8bec552be6353b623a8f12078befd968a662a933907d0ff72b0164fa8b38e4cbd4aae6191e28aedc19d996ec99298053d7bafc458f91580a7cfa
[kubernetes-client-linux-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-linux-amd64.tar.gz) | 34ab1e9edf70c84fe58a223e91b0bf679e5d1273a2b6503a18a61a4bea79231948efe098f84e39f83dbfa2c5271aad8e9819aac104d6c3360c97e5c348b15be7
[kubernetes-client-linux-arm.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-linux-arm.tar.gz) | 554f8240597e7eb8470c8e2b4bca33c06a0a91746831ef93f76b344f5c3d6226d4ef26cb59f127d1436ac091b7b79395320fe8a9b2acc512afd601989c138d4a
[kubernetes-client-linux-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-linux-arm64.tar.gz) | 62c750654e898622aa87b2d81d4b0cbaf36614899f37181c2e3a6aa645d2270c4dedb7d6e7b974059a716ad144a5615407e6a9a4c03f761a512d37fdda796e50
[kubernetes-client-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-linux-ppc64le.tar.gz) | bbe22979c4e300675dfa955ce9855b7b33b29119a9f78e58bc1b088dba2b8dedbf0b068092d42cd98dbc10cda1da317eae6414b91587593678ec76780ff575df
[kubernetes-client-linux-s390x.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-linux-s390x.tar.gz) | b8360d0bf930149d360da2a95549b35cb7e14932ae8507d99e34d93729ae645bab203dfba325c74db13204e09e5ee032f887cdd67badfbf3bc8a08d71ccf9c3d
[kubernetes-client-windows-386.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-windows-386.tar.gz) | cc935a74f30dcd1eaaeadc8f2353a9742ebc4a36b133342c6402b065750f4028a1a392bd5f7ea51533c2d799ff2bdb3d0f21493e7fabacb27289e58f011bb229
[kubernetes-client-windows-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-windows-amd64.tar.gz) | 119742103985be0cd4296b85aa2c713cdc510b9a9412706fdf88ca1c703f69338146efc5cf37168ab56e74576ad561ce37c3f500d29b63002139d11544b1b7cf
[kubernetes-client-windows-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-client-windows-arm64.tar.gz) | 67f5fe6b14aa4c49acb6f706ec4a9e43b87f1e19555579895452183e0b2d2be2202f8d48622208ac5ef6e0fb9050d99bb7e1ed9e4e31e8fcac7c0b5e44787c39

### Server Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-server-linux-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-server-linux-amd64.tar.gz) | a1c935db625766b02113087068fa1087a6b74e9f57ad72cc1d5d85e830c0569b9257746013053ee8dc89404940458c3bba00064666978ddff4df9f3cae0ae066
[kubernetes-server-linux-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-server-linux-arm64.tar.gz) | 5f87f8f46719af413fa864b7d91d5b95fa86adf27df63cf904470b15e844ebda4c802d7ec7cb4006b4e5a3780903d0436eb57a7e2f2b79b74cf5e7e8f65496b1
[kubernetes-server-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-server-linux-ppc64le.tar.gz) | 9f2e9476ae3b95919c991dd0438b31eeded7c7f686948ef2d6227311dabc952e74f61d3638f224eb18e63a4fe4955b55a9e358032477b989800541617a8b5f6a
[kubernetes-server-linux-s390x.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-server-linux-s390x.tar.gz) | 043311de9dad81d3774decc3aeeac48833d2d234a15ce4a2062fd9af778879afde5e2eb9fdec5f3641725f630e7c3bd845a348ad713b54275e77293941d4c8d0

### Node Binaries

filename | sha512 hash
-------- | -----------
[kubernetes-node-linux-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-node-linux-amd64.tar.gz) | 865b4ee818cd53bc91001f658243e6b6fd9464f17ef8dc0cc739586689f39998e5c47630df439ad43553d6830ae3a7375cc780c3a5e49f1422d35d77194efc35
[kubernetes-node-linux-arm64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-node-linux-arm64.tar.gz) | fc7df94bc328817d20c59e1ab1371634bf3849141ad982c9d403b136d009c3ed9ee3f7a20659a0f93af7179e13e2c0835b80fefc677b8840f7bcbde0dabb4483
[kubernetes-node-linux-ppc64le.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-node-linux-ppc64le.tar.gz) | 8ec718436680766d9026b56ece5bff7a6bac9f63da0edf33843a7a6c255e1a1d22aecf260c1d5fc394c1e2f581931c5ae0acad80a0141fc8d4d7730bf04566b5
[kubernetes-node-linux-s390x.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-node-linux-s390x.tar.gz) | fff0562e3a89b4f9444ce83bb8cfc860bcd5178a2c1ba0f1404d87556f483d48241aa3f927f2534d811b1c554958c167f1e5199fa916ebfd9a754cef2f761139
[kubernetes-node-windows-amd64.tar.gz](https://dl.k8s.io/v1.35.0-alpha.1/kubernetes-node-windows-amd64.tar.gz) | 7774299a2b581a4ab2514d8dc95c1d0dff0652aa0518efbffa9ac39f5c510cb83d8a41a70e9b4e78f0b179e5a806402310dac161ff3a2398b97755938c225586

### Container Images

All container images are available as manifest lists and support the described
architectures. It is also possible to pull a specific architecture directly by
adding the "-$ARCH" suffix  to the container image name.

name | architectures
---- | -------------
[registry.k8s.io/conformance:v1.35.0-alpha.1](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/conformance-s390x)
[registry.k8s.io/kube-apiserver:v1.35.0-alpha.1](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-apiserver-s390x)
[registry.k8s.io/kube-controller-manager:v1.35.0-alpha.1](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-controller-manager-s390x)
[registry.k8s.io/kube-proxy:v1.35.0-alpha.1](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-proxy-s390x)
[registry.k8s.io/kube-scheduler:v1.35.0-alpha.1](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kube-scheduler-s390x)
[registry.k8s.io/kubectl:v1.35.0-alpha.1](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl) | [amd64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-amd64), [arm64](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-arm64), [ppc64le](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-ppc64le), [s390x](https://console.cloud.google.com/artifacts/docker/k8s-artifacts-prod/southamerica-east1/images/kubectl-s390x)

## Changelog since v1.34.0

## Changes by Kind

### API Change

- Added WithOrigin within apis/core/validation with adjusted tests ([#132825](https://github.com/kubernetes/kubernetes/pull/132825), [@PatrickLaabs](https://github.com/PatrickLaabs)) [SIG Apps]
- Component-base: validate that log-flush-frequency is positive and return an error instead of panic-ing ([#133540](https://github.com/kubernetes/kubernetes/pull/133540), [@BenTheElder](https://github.com/BenTheElder)) [SIG Architecture, Instrumentation, Network and Node]
- Feature gate dependencies are now explicit, and validated at startup. A feature can no longer be enabled if it depends on a disabled feature. In particular, this means that `AllAlpha=true` will no longer work without enabling disabled-by-default beta features that are depended on (either with `AllBeta=true` or explicitly enumerating the disabled dependencies). ([#133697](https://github.com/kubernetes/kubernetes/pull/133697), [@tallclair](https://github.com/tallclair)) [SIG API Machinery, Architecture, Cluster Lifecycle and Node]
- In version 1.34, the PodObservedGenerationTracking feature has been upgraded to beta, and the description of the alpha version in the openapi has been removed. ([#133883](https://github.com/kubernetes/kubernetes/pull/133883), [@yangjunmyfm192085](https://github.com/yangjunmyfm192085)) [SIG Apps]
- Introduce a new declarative validation tag +k8s:customUnique to control listmap uniqueness ([#134279](https://github.com/kubernetes/kubernetes/pull/134279), [@yongruilin](https://github.com/yongruilin)) [SIG API Machinery and Auth]
- Kube-apiserver: Fixed a 1.34 regression in CustomResourceDefinition handling that incorrectly warned about unrecognized formats on number and integer properties ([#133896](https://github.com/kubernetes/kubernetes/pull/133896), [@yongruilin](https://github.com/yongruilin)) [SIG API Machinery, Apps, Architecture, Auth, CLI, Cloud Provider, Contributor Experience, Network, Node and Scheduling]
- OpenAPI model packages of API types are generated into `zz_generated.model_name.go` files and are accessible using the `OpenAPIModelName()` function.  This allows API authors to declare the desired OpenAPI model packages instead of using the go package path of API types. ([#131755](https://github.com/kubernetes/kubernetes/pull/131755), [@jpbetz](https://github.com/jpbetz)) [SIG API Machinery, Apps, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Network, Node, Scheduling, Storage and Testing]
- Support for `kubectl get -o kyaml` is now on by default.  To disable it, set `KUBECTL_KYAML=false`. ([#133327](https://github.com/kubernetes/kubernetes/pull/133327), [@thockin](https://github.com/thockin)) [SIG CLI]
- The storage version for MutatingAdmissionPolicy is updated to v1beta1. ([#133715](https://github.com/kubernetes/kubernetes/pull/133715), [@cici37](https://github.com/cici37)) [SIG API Machinery, Etcd and Testing]

### Feature

- Add paths section to kubelet statusz endpoint ([#133239](https://github.com/kubernetes/kubernetes/pull/133239), [@Peac36](https://github.com/Peac36)) [SIG Node]
- Add paths section to scheduler statusz endpoint ([#132606](https://github.com/kubernetes/kubernetes/pull/132606), [@Peac36](https://github.com/Peac36)) [SIG API Machinery, Architecture, Instrumentation, Network, Node, Scheduling and Testing]
- Added kubectl config set-context -n flag as a shorthand for --namespace ([#134384](https://github.com/kubernetes/kubernetes/pull/134384), [@tchap](https://github.com/tchap)) [SIG CLI and Testing]
- Added remote runtime and image `Close()` method to be able to close the connection. ([#133211](https://github.com/kubernetes/kubernetes/pull/133211), [@saschagrunert](https://github.com/saschagrunert)) [SIG Node]
- Adds metric for Maxunavailable feature ([#130951](https://github.com/kubernetes/kubernetes/pull/130951), [@Edwinhr716](https://github.com/Edwinhr716)) [SIG Apps and Instrumentation]
- Applyconfiguration-gen now generates extract functions for all subresources ([#132665](https://github.com/kubernetes/kubernetes/pull/132665), [@mrIncompetent](https://github.com/mrIncompetent)) [SIG API Machinery]
- Applyconfiguration-gen now preserves struct and field comments from source types in generated code ([#132663](https://github.com/kubernetes/kubernetes/pull/132663), [@mrIncompetent](https://github.com/mrIncompetent)) [SIG API Machinery]
- DRA: the resource.k8s.io API now uses the v1 API version (introduced in 1.34) as default storage version. Downgrading to 1.33 is not supported. ([#133876](https://github.com/kubernetes/kubernetes/pull/133876), [@kei01234kei](https://github.com/kei01234kei)) [SIG API Machinery, Etcd and Testing]
- Events:
    Type     Reason   Age                 From               Message
    ----     ------   ----                ----               -------
    Warning  Failed   7m11s (x2 over 7m33s) kubelet          spec.containers{nginx}: Failed to pull image "nginx": failed to pull and unpack image... ([#133627](https://github.com/kubernetes/kubernetes/pull/133627), [@itzPranshul](https://github.com/itzPranshul)) [SIG CLI]
- Introduces e2e tests that check component invariant metrics across the entire suite run. ([#133394](https://github.com/kubernetes/kubernetes/pull/133394), [@BenTheElder](https://github.com/BenTheElder)) [SIG Testing]
- K8s.io/apimachinery: Introduce a helper function to compare resourceVersion strings from two objects of the same resource ([#134330](https://github.com/kubernetes/kubernetes/pull/134330), [@michaelasp](https://github.com/michaelasp)) [SIG API Machinery, Apps, Auth, Instrumentation, Network, Node, Scheduling, Storage and Testing]
- Kubeadm: graduate the kubeadm specific feature gate ControlPlaneKubeletLocalMode to GA and lock it to enabled by default. To opt-out manually from this desired default behavior you must patch the "server" field in the  /etc/kubernetes/kubelet.conf file. The subphase of "kubeadm join phase control-plane-join" called "etcd" is now deprecated, hidden and replaced by the subphase with identical functionality "etcd-join". "etcd" will be removed in a follow-up release. The subphase "kubelet-wait-bootstrap" of "kubeadm join" is no longer experimental and will always run. ([#134106](https://github.com/kubernetes/kubernetes/pull/134106), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubernetes is now built using Go 1.25.1 ([#134095](https://github.com/kubernetes/kubernetes/pull/134095), [@dims](https://github.com/dims)) [SIG Release and Testing]
- Kubernetes now uses Go Language Version 1.25, including https://go.dev/blog/container-aware-gomaxprocs ([#134120](https://github.com/kubernetes/kubernetes/pull/134120), [@BenTheElder](https://github.com/BenTheElder)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Network, Node, Release, Scheduling and Storage]
- Lock down the `AllowOverwriteTerminationGracePeriodSeconds` feature gate. ([#133792](https://github.com/kubernetes/kubernetes/pull/133792), [@HirazawaUi](https://github.com/HirazawaUi)) [SIG Node]
- Metrics: exclude dryRun requests from apiserver_request_sli_duration_seconds ([#131092](https://github.com/kubernetes/kubernetes/pull/131092), [@aldudko](https://github.com/aldudko)) [SIG API Machinery and Instrumentation]
- The validation in the resouce.k8s.io has been migrated to declarative validation.
  If the `DeclarativeValidation` feature gate is enabled, mismatches with existing validation are reported via metrics.
  If the `DeclarativeValidationTakeover` feature gate is enabled, declarative validation is the primary source of errors for migrated fields. ([#134072](https://github.com/kubernetes/kubernetes/pull/134072), [@yongruilin](https://github.com/yongruilin)) [SIG API Machinery, Apps and Auth]

### Bug or Regression

- Added the correct error when eviction is blocked due to the failSafe mechanism of the DisruptionController. ([#133097](https://github.com/kubernetes/kubernetes/pull/133097), [@kei01234kei](https://github.com/kei01234kei)) [SIG Apps and Node]
- Bugfix: the default serviceCIDR controller was not logging events because the event broadcaster was shutdown during its initialization. ([#133338](https://github.com/kubernetes/kubernetes/pull/133338), [@aojea](https://github.com/aojea)) [SIG Network]
- Deprecated metrics will be hidden as per the metrics deprecation policy https://kubernetes.io/docs/reference/using-api/deprecation-policy/#deprecating-a-metric ([#133436](https://github.com/kubernetes/kubernetes/pull/133436), [@richabanker](https://github.com/richabanker)) [SIG Architecture, Instrumentation and Network]
- Fix incorrect behavior of preemptor pod when preemption of the victim takes long to complete. The preemptor pod should not be circling in scheduling cycles until preemption is finished. ([#134294](https://github.com/kubernetes/kubernetes/pull/134294), [@ania-borowiec](https://github.com/ania-borowiec)) [SIG Scheduling and Testing]
- Fix missing kubelet_volume_stats_* metrics ([#133890](https://github.com/kubernetes/kubernetes/pull/133890), [@huww98](https://github.com/huww98)) [SIG Instrumentation and Node]
- Fix occasional schedule delay when the static PV is created ([#133929](https://github.com/kubernetes/kubernetes/pull/133929), [@huww98](https://github.com/huww98)) [SIG Scheduling and Storage]
- Fix resource claims deallocation for extended resource when pod is completed ([#134312](https://github.com/kubernetes/kubernetes/pull/134312), [@alaypatel07](https://github.com/alaypatel07)) [SIG Apps, Node and Testing]
- Fixed SELinux warning controller not emitting events on some SELinux label conflicts. ([#133425](https://github.com/kubernetes/kubernetes/pull/133425), [@jsafrane](https://github.com/jsafrane)) [SIG Apps, Storage and Testing]
- Fixed a bug in kube-proxy nftables mode (GA as of 1.33) that fails to determine if traffic originates from a local source on the node. The issue was caused by using the wrong meta `iif` instead of `iifname` for name based matches. ([#134024](https://github.com/kubernetes/kubernetes/pull/134024), [@jack4it](https://github.com/jack4it)) [SIG Network]
- Fixed a bug in kube-scheduler where pending pod preemption caused preemptor pods to be retried more frequently. ([#134245](https://github.com/kubernetes/kubernetes/pull/134245), [@macsko](https://github.com/macsko)) [SIG Scheduling and Testing]
- Fixed a bug that caused apiservers to send an inappropriate Content-Type request header to authorization, token authentication, imagepolicy admission, and audit webhooks when the alpha client-go feature gate "ClientsPreferCBOR" is enabled. ([#132960](https://github.com/kubernetes/kubernetes/pull/132960), [@benluddy](https://github.com/benluddy)) [SIG API Machinery and Node]
- Fixed a bug that caused duplicate validation when updating PersistentVolumeClaims, VolumeAttachments and VolumeAttributesClasses. ([#132549](https://github.com/kubernetes/kubernetes/pull/132549), [@gavinkflam](https://github.com/gavinkflam)) [SIG Storage]
- Fixed a bug that caused duplicate validation when updating role and role binding resources. ([#132550](https://github.com/kubernetes/kubernetes/pull/132550), [@gavinkflam](https://github.com/gavinkflam)) [SIG Auth]
- Fixed a bug where high latency kube-apiserver caused scheduling throughput degradation. ([#134154](https://github.com/kubernetes/kubernetes/pull/134154), [@macsko](https://github.com/macsko)) [SIG Scheduling]
- Fixed broken shell completion for api resources. ([#133771](https://github.com/kubernetes/kubernetes/pull/133771), [@marckhouzam](https://github.com/marckhouzam)) [SIG CLI]
- Fixed validation error when ConfigFlags has CertFile and (or) KeyFile and original config also contains CertFileData and (or) KeyFileData. ([#133917](https://github.com/kubernetes/kubernetes/pull/133917), [@n2h9](https://github.com/n2h9)) [SIG API Machinery and CLI]
- Fixes a possible data race during metrics registration ([#134390](https://github.com/kubernetes/kubernetes/pull/134390), [@liggitt](https://github.com/liggitt)) [SIG Architecture and Instrumentation]
- Implicit extended resource name derived from device class (deviceclass.resource.kubernetes.io/<device-class-name>) can be used to request DRA devices matching the device class. ([#133363](https://github.com/kubernetes/kubernetes/pull/133363), [@yliaog](https://github.com/yliaog)) [SIG Node, Scheduling and Testing]
- Kube-apiserver: Fixes a 1.34 regression with spurious "Error getting keys" log messages ([#133817](https://github.com/kubernetes/kubernetes/pull/133817), [@serathius](https://github.com/serathius)) [SIG API Machinery and Etcd]
- Kube-apiserver: Fixes a possible 1.34 performance regression calculating object size statistics for resources not served from the watch cache, typically only Events ([#133873](https://github.com/kubernetes/kubernetes/pull/133873), [@serathius](https://github.com/serathius)) [SIG API Machinery and Etcd]
- Kube-apiserver: improve the validation error message shown for custom resources with CEL validation rules to include the value that failed validation ([#132798](https://github.com/kubernetes/kubernetes/pull/132798), [@cbandy](https://github.com/cbandy)) [SIG API Machinery]
- Kube-controller-manager: Fixes a possible data race in the garbage collection controller ([#134379](https://github.com/kubernetes/kubernetes/pull/134379), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Apps]
- Kubeadm: ensured waiting for apiserver uses a local client that doesn't reach to the control plane endpoint and instead reaches directly to the local API server endpoint. ([#134265](https://github.com/kubernetes/kubernetes/pull/134265), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: fix KUBEADM_UPGRADE_DRYRUN_DIR not honored in upgrade phase when writing kubelet config files ([#134007](https://github.com/kubernetes/kubernetes/pull/134007), [@carlory](https://github.com/carlory)) [SIG Cluster Lifecycle]
- Kubeadm: fixed a bug where the node registration information for a given node was not fetched correctly during "kubeadm upgrade node" and the node name can end up being incorrect in cases where the node name is not the same as the host name. ([#134319](https://github.com/kubernetes/kubernetes/pull/134319), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: fixed bug where v1beta3's ClusterConfiguration.APIServer.TimeoutForControlPlane is not respected in newer versions of kubeadm where v1beta4 is the default. ([#133513](https://github.com/kubernetes/kubernetes/pull/133513), [@tom1299](https://github.com/tom1299)) [SIG Cluster Lifecycle]
- Kubelet: the connection to a DRA driver became unusable because of an internal deadlock when a connection was idle for 30 minutes. ([#133926](https://github.com/kubernetes/kubernetes/pull/133926), [@pohly](https://github.com/pohly)) [SIG Node]
- Pod can have multiple volumes reference the same PVC ([#122140](https://github.com/kubernetes/kubernetes/pull/122140), [@huww98](https://github.com/huww98)) [SIG Node, Storage and Testing]
- Previously, `kubectl scale` returned the error message `error: no objects passed to scale <GroupResource> "<ResourceName>" not found` when the specified resource did not exist. 
  For consistency with other commands(e.g. `kubectl get`), it has been changed to just return `Error from server (NotFound): <GroupResource> "<ResourceName>" not found`. ([#134017](https://github.com/kubernetes/kubernetes/pull/134017), [@mochizuki875](https://github.com/mochizuki875)) [SIG CLI]
- Promote VAC API test to conformance ([#133615](https://github.com/kubernetes/kubernetes/pull/133615), [@carlory](https://github.com/carlory)) [SIG Architecture, Storage and Testing]
- Remove incorrectly printed warning for SessionAffinity whenever a headless service is creater or updated ([#134054](https://github.com/kubernetes/kubernetes/pull/134054), [@Peac36](https://github.com/Peac36)) [SIG Network]
- The SchedulerAsyncAPICalls feature gate has been disabled to mitigate a bug where its interaction with asynchronous preemption in could degrade kube-scheduler performance, particularly under high kube-apiserver load. ([#134400](https://github.com/kubernetes/kubernetes/pull/134400), [@macsko](https://github.com/macsko)) [SIG Scheduling]
- When image garbage collection is unable to free enough disk space, the FreeDiskSpaceFailed warning event is now more actionable. Example: `Insufficient free disk space on the node's image filesystem (95.0% of 10.0 GiB used). Failed to free sufficient space by deleting unused images. Consider resizing the disk or deleting unused files.` ([#132578](https://github.com/kubernetes/kubernetes/pull/132578), [@drigz](https://github.com/drigz)) [SIG Node]

### Other (Cleanup or Flake)

- Bump addon manager to use kubectl v1.32.2 ([#130548](https://github.com/kubernetes/kubernetes/pull/130548), [@Jefftree](https://github.com/Jefftree)) [SIG Cloud Provider, Scalability and Testing]
- Dropping the experimental prefix from kubectl wait command's short description, since kubectl wait command has been stable for a long time. ([#133907](https://github.com/kubernetes/kubernetes/pull/133907), [@ardaguclu](https://github.com/ardaguclu)) [SIG CLI]
- Fix formatting of assorted go API deprecations for godoc / pkgsite and enable a linter to help catch mis-formatted deprecations ([#133571](https://github.com/kubernetes/kubernetes/pull/133571), [@BenTheElder](https://github.com/BenTheElder)) [SIG API Machinery, Architecture, CLI, Instrumentation and Testing]
- Improved HPA performance when using container-specific resource metrics by optimizing container lookup logic to exit early once the target container is found, reducing unnecessary iterations through all containers in a pod. ([#133415](https://github.com/kubernetes/kubernetes/pull/133415), [@AadiDev005](https://github.com/AadiDev005)) [SIG Apps and Autoscaling]
- Kube-apiserver: Fixes an issue where passing invalid DeleteOptions incorrectly returned status 500 rather than 400. ([#133358](https://github.com/kubernetes/kubernetes/pull/133358), [@ostrain](https://github.com/ostrain)) [SIG API Machinery]
- Kubeadm: removed the `RootlessControlPlane` feature gate. User Namespaces will serve as its replacement. ([#134178](https://github.com/kubernetes/kubernetes/pull/134178), [@HirazawaUi](https://github.com/HirazawaUi)) [SIG Cluster Lifecycle]
- Remove container name from messages for container created and started events. ([#134043](https://github.com/kubernetes/kubernetes/pull/134043), [@HirazawaUi](https://github.com/HirazawaUi)) [SIG Node]
- Removed deprecated gogo protocol definitions from `k8s.io/kubelet/pkg/apis/dra` in favor of `google.golang.org/protobuf`. ([#133026](https://github.com/kubernetes/kubernetes/pull/133026), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Removed general available feature-gate SizeMemoryBackedVolumes ([#133720](https://github.com/kubernetes/kubernetes/pull/133720), [@carlory](https://github.com/carlory)) [SIG Node, Storage and Testing]
- Removed the `ComponentSLIs` feature gate, which had been promoted to stable as part of the Kubernetes 1.32 release. ([#133742](https://github.com/kubernetes/kubernetes/pull/133742), [@carlory](https://github.com/carlory)) [SIG Architecture and Instrumentation]
- Removing Experimental prefix from the description of kubectl wait to emphasize that it is stable. ([#133731](https://github.com/kubernetes/kubernetes/pull/133731), [@ardaguclu](https://github.com/ardaguclu)) [SIG CLI]
- Removing the KUBECTL_OPENAPIV3_PATCH environment variable entirely, since aggregated discovery has been stable from 1.30. ([#134130](https://github.com/kubernetes/kubernetes/pull/134130), [@ardaguclu](https://github.com/ardaguclu)) [SIG CLI]
- Specifies the deprecated version of apiserver_storage_objects metric in metrics docs ([#134028](https://github.com/kubernetes/kubernetes/pull/134028), [@richabanker](https://github.com/richabanker)) [SIG API Machinery, Etcd and Instrumentation]
- Tests: switch to https://go.dev/doc/go1.25#container-aware-gomaxprocs from go.uber.org/automaxprocs ([#133492](https://github.com/kubernetes/kubernetes/pull/133492), [@BenTheElder](https://github.com/BenTheElder)) [SIG Testing]
- The `/statusz` page for `kube-proxy` now includes a list of exposed endpoints, making it easier to debug and introspect. ([#133190](https://github.com/kubernetes/kubernetes/pull/133190), [@aman4433](https://github.com/aman4433)) [SIG Network and Node]
- Types in k/k/pkg/scheduler/framework:
  Handle,
  Plugin,
  PreEnqueuePlugin, QueueSortPlugin, EnqueueExtensions, PreFilterExtensions, PreFilterPlugin, FilterPlugin, PostFilterPlugin, PreScorePlugin, ScorePlugin, ReservePlugin, PreBindPlugin, PostBindPlugin, PermitPlugin, BindPlugin,
  PodActivator, PodNominator, PluginsRunner,
  LessFunc, ScoreExtensions, NodeToStatusReader, NodeScoreList, NodeScore, NodePluginScores, PluginScore, NominatingMode, NominatingInfo, WaitingPod, PreFilterResult, PostFilterResult,
  Extender,
  NodeInfoLister, StorageInfoLister, SharedLister, ResourceSliceLister, DeviceClassLister, ResourceClaimTracker, SharedDRAManager
  
  are moved to package k8s.io/kube-scheduler/framework . Users should update import paths. The interfaces don't change.
  
  Type Parallelizer in k/k/pkg/scheduler/framework/parallelism is split into interface Parallelizer (in k8s.io/kube-scheduler/framework) and struct Parallelizer (location unchanged in k/k). Plugin developers should update the import path to staging repo. ([#133172](https://github.com/kubernetes/kubernetes/pull/133172), [@ania-borowiec](https://github.com/ania-borowiec)) [SIG Node, Release, Scheduling, Storage and Testing]
- Updated CNI plugins to v1.8.0. ([#133837](https://github.com/kubernetes/kubernetes/pull/133837), [@saschagrunert](https://github.com/saschagrunert)) [SIG Cloud Provider, Node and Testing]
- Updated cri-tools to v1.34.0. ([#133636](https://github.com/kubernetes/kubernetes/pull/133636), [@saschagrunert](https://github.com/saschagrunert)) [SIG Cloud Provider]
- Updated etcd to v3.6.5. ([#134251](https://github.com/kubernetes/kubernetes/pull/134251), [@joshjms](https://github.com/joshjms)) [SIG API Machinery, Cloud Provider, Cluster Lifecycle, Etcd and Testing]
- Upgrade CoreDNS to v1.12.3 ([#132288](https://github.com/kubernetes/kubernetes/pull/132288), [@thevilledev](https://github.com/thevilledev)) [SIG Cloud Provider and Cluster Lifecycle]
- `kubectl auth reconcile` now re-attempts reconciliation if it encounters a conflict error ([#133323](https://github.com/kubernetes/kubernetes/pull/133323), [@liggitt](https://github.com/liggitt)) [SIG Auth and CLI]
- `kubectl get` and `kubectl describe` human-readable output no longer includes counts for referenced tokens and secrets ([#117160](https://github.com/kubernetes/kubernetes/pull/117160), [@liggitt](https://github.com/liggitt)) [SIG CLI and Testing]

## Dependencies

### Added
- github.com/moby/sys/atomicwriter: [v0.1.0](https://github.com/moby/sys/tree/atomicwriter/v0.1.0)
- golang.org/x/tools/go/expect: v0.1.1-deprecated
- golang.org/x/tools/go/packages/packagestest: v0.1.1-deprecated

### Changed
- cloud.google.com/go/compute/metadata: v0.6.0 → v0.7.0
- github.com/aws/aws-sdk-go-v2/config: [v1.27.24 → v1.29.14](https://github.com/aws/aws-sdk-go-v2/compare/config/v1.27.24...config/v1.29.14)
- github.com/aws/aws-sdk-go-v2/credentials: [v1.17.24 → v1.17.67](https://github.com/aws/aws-sdk-go-v2/compare/credentials/v1.17.24...credentials/v1.17.67)
- github.com/aws/aws-sdk-go-v2/feature/ec2/imds: [v1.16.9 → v1.16.30](https://github.com/aws/aws-sdk-go-v2/compare/feature/ec2/imds/v1.16.9...feature/ec2/imds/v1.16.30)
- github.com/aws/aws-sdk-go-v2/internal/configsources: [v1.3.13 → v1.3.34](https://github.com/aws/aws-sdk-go-v2/compare/internal/configsources/v1.3.13...internal/configsources/v1.3.34)
- github.com/aws/aws-sdk-go-v2/internal/endpoints/v2: [v2.6.13 → v2.6.34](https://github.com/aws/aws-sdk-go-v2/compare/internal/endpoints/v2/v2.6.13...internal/endpoints/v2/v2.6.34)
- github.com/aws/aws-sdk-go-v2/internal/ini: [v1.8.0 → v1.8.3](https://github.com/aws/aws-sdk-go-v2/compare/internal/ini/v1.8.0...internal/ini/v1.8.3)
- github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding: [v1.11.3 → v1.12.3](https://github.com/aws/aws-sdk-go-v2/compare/service/internal/accept-encoding/v1.11.3...service/internal/accept-encoding/v1.12.3)
- github.com/aws/aws-sdk-go-v2/service/internal/presigned-url: [v1.11.15 → v1.12.15](https://github.com/aws/aws-sdk-go-v2/compare/service/internal/presigned-url/v1.11.15...service/internal/presigned-url/v1.12.15)
- github.com/aws/aws-sdk-go-v2/service/sso: [v1.22.1 → v1.25.3](https://github.com/aws/aws-sdk-go-v2/compare/service/sso/v1.22.1...service/sso/v1.25.3)
- github.com/aws/aws-sdk-go-v2/service/ssooidc: [v1.26.2 → v1.30.1](https://github.com/aws/aws-sdk-go-v2/compare/service/ssooidc/v1.26.2...service/ssooidc/v1.30.1)
- github.com/aws/aws-sdk-go-v2/service/sts: [v1.30.1 → v1.33.19](https://github.com/aws/aws-sdk-go-v2/compare/service/sts/v1.30.1...service/sts/v1.33.19)
- github.com/aws/aws-sdk-go-v2: [v1.30.1 → v1.36.3](https://github.com/aws/aws-sdk-go-v2/compare/v1.30.1...v1.36.3)
- github.com/aws/smithy-go: [v1.20.3 → v1.22.3](https://github.com/aws/smithy-go/compare/v1.20.3...v1.22.3)
- github.com/containerd/containerd/api: [v1.8.0 → v1.9.0](https://github.com/containerd/containerd/compare/api/v1.8.0...api/v1.9.0)
- github.com/containerd/ttrpc: [v1.2.6 → v1.2.7](https://github.com/containerd/ttrpc/compare/v1.2.6...v1.2.7)
- github.com/containerd/typeurl/v2: [v2.2.2 → v2.2.3](https://github.com/containerd/typeurl/compare/v2.2.2...v2.2.3)
- github.com/coredns/corefile-migration: [v1.0.26 → v1.0.27](https://github.com/coredns/corefile-migration/compare/v1.0.26...v1.0.27)
- github.com/docker/docker: [v26.1.4+incompatible → v28.2.2+incompatible](https://github.com/docker/docker/compare/v26.1.4...v28.2.2)
- github.com/go-logr/logr: [v1.4.2 → v1.4.3](https://github.com/go-logr/logr/compare/v1.4.2...v1.4.3)
- github.com/google/cadvisor: [v0.52.1 → v0.53.0](https://github.com/google/cadvisor/compare/v0.52.1...v0.53.0)
- github.com/opencontainers/cgroups: [v0.0.1 → v0.0.3](https://github.com/opencontainers/cgroups/compare/v0.0.1...v0.0.3)
- github.com/opencontainers/runc: [v1.2.5 → v1.3.0](https://github.com/opencontainers/runc/compare/v1.2.5...v1.3.0)
- github.com/opencontainers/runtime-spec: [v1.2.0 → v1.2.1](https://github.com/opencontainers/runtime-spec/compare/v1.2.0...v1.2.1)
- github.com/prometheus/client_golang: [v1.22.0 → v1.23.2](https://github.com/prometheus/client_golang/compare/v1.22.0...v1.23.2)
- github.com/prometheus/client_model: [v0.6.1 → v0.6.2](https://github.com/prometheus/client_model/compare/v0.6.1...v0.6.2)
- github.com/prometheus/common: [v0.62.0 → v0.66.1](https://github.com/prometheus/common/compare/v0.62.0...v0.66.1)
- github.com/prometheus/procfs: [v0.15.1 → v0.16.1](https://github.com/prometheus/procfs/compare/v0.15.1...v0.16.1)
- github.com/spf13/cobra: [v1.9.1 → v1.10.0](https://github.com/spf13/cobra/compare/v1.9.1...v1.10.0)
- github.com/spf13/pflag: [v1.0.6 → v1.0.9](https://github.com/spf13/pflag/compare/v1.0.6...v1.0.9)
- github.com/stretchr/testify: [v1.10.0 → v1.11.1](https://github.com/stretchr/testify/compare/v1.10.0...v1.11.1)
- go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp: v0.58.0 → v0.61.0
- go.opentelemetry.io/otel/metric: v1.35.0 → v1.36.0
- go.opentelemetry.io/otel/sdk/metric: v1.34.0 → v1.36.0
- go.opentelemetry.io/otel/sdk: v1.34.0 → v1.36.0
- go.opentelemetry.io/otel/trace: v1.35.0 → v1.36.0
- go.opentelemetry.io/otel: v1.35.0 → v1.36.0
- golang.org/x/crypto: v0.36.0 → v0.41.0
- golang.org/x/mod: v0.21.0 → v0.27.0
- golang.org/x/net: v0.38.0 → v0.43.0
- golang.org/x/oauth2: v0.27.0 → v0.30.0
- golang.org/x/sync: v0.12.0 → v0.16.0
- golang.org/x/sys: v0.31.0 → v0.35.0
- golang.org/x/telemetry: bda5523 → 1a19826
- golang.org/x/term: v0.30.0 → v0.34.0
- golang.org/x/text: v0.23.0 → v0.28.0
- golang.org/x/tools: v0.26.0 → v0.36.0
- google.golang.org/genproto/googleapis/rpc: a0af3ef → 200df99
- google.golang.org/grpc: v1.72.1 → v1.72.2
- google.golang.org/protobuf: v1.36.5 → v1.36.8
- gopkg.in/evanphx/json-patch.v4: v4.12.0 → v4.13.0
- k8s.io/gengo/v2: 85fd79d → ec3ebc5
- k8s.io/kube-openapi: f3f2b99 → 589584f
- k8s.io/system-validators: v1.10.1 → v1.11.1
- sigs.k8s.io/json: cfa47c3 → 2d32026

### Removed
- gopkg.in/yaml.v2: v2.4.0