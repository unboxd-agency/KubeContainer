/*
Copyright 2026.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ExposeType selects how the workload is exposed.
// +kubebuilder:validation:Enum=ClusterIP;LoadBalancer;Ingress
type ExposeType string

const (
	ExposeClusterIP    ExposeType = "ClusterIP"
	ExposeLoadBalancer ExposeType = "LoadBalancer"
	ExposeIngress      ExposeType = "Ingress"
)

// Condition types reported in KubeContainerStatus.
const (
	ConditionReady       = "Ready"
	ConditionProgressing = "Progressing"
	ConditionDegraded    = "Degraded"
)

// Autoscale configures a HorizontalPodAutoscaler for the workload.
type Autoscale struct {
	// minReplicas is the lower bound for the autoscaler.
	// +kubebuilder:validation:Minimum=1
	MinReplicas int32 `json:"minReplicas"`

	// maxReplicas is the upper bound for the autoscaler.
	// +kubebuilder:validation:Minimum=1
	MaxReplicas int32 `json:"maxReplicas"`

	// targetCPUUtilization is the average CPU utilization (percent of
	// requests) the autoscaler aims for.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:default=80
	// +optional
	TargetCPUUtilization int32 `json:"targetCPUUtilization,omitempty"`
}

// Scaling selects either a fixed replica count or autoscaling.
// +kubebuilder:validation:XValidation:rule="!(has(self.replicas) && has(self.autoscale))",message="replicas and autoscale are mutually exclusive"
type Scaling struct {
	// replicas is a fixed replica count. Mutually exclusive with autoscale.
	// +kubebuilder:validation:Minimum=0
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// autoscale configures an HPA. Mutually exclusive with replicas.
	// +optional
	Autoscale *Autoscale `json:"autoscale,omitempty"`
}

// Expose configures how the workload is reachable.
// +kubebuilder:validation:XValidation:rule="self.type != 'Ingress' || has(self.host)",message="host is required when type is Ingress"
type Expose struct {
	// type selects the exposure mechanism.
	// +kubebuilder:default=ClusterIP
	// +optional
	Type ExposeType `json:"type,omitempty"`

	// host is the DNS name routed to the workload. Required when type is Ingress.
	// +optional
	Host string `json:"host,omitempty"`
}

// HealthCheck wires HTTP liveness and readiness probes to the container.
type HealthCheck struct {
	// path is the HTTP path probed on the container port.
	// +kubebuilder:validation:Pattern=`^/.*`
	Path string `json:"path"`
}

// Storage declares what the workload keeps: a PersistentVolumeClaim
// owned by the KubeContainer, mounted at path. Omitted, the kube is
// stateless by contract — everything written dies with the pod.
type Storage struct {
	// size is the requested capacity, a Kubernetes quantity (e.g. "1Gi").
	// +kubebuilder:validation:Pattern=`^[0-9]+(\.[0-9]+)?(Ki|Mi|Gi|Ti|Pi|Ei|m|k|M|G|T|P|E)?$`
	Size string `json:"size"`

	// path is where the volume is mounted in the container.
	// +kubebuilder:validation:Pattern=`^/.+`
	Path string `json:"path"`
}

// KubeContainerSpec defines the desired state of KubeContainer
type KubeContainerSpec struct {
	// image is the container image reference to run.
	// +kubebuilder:validation:MinLength=1
	Image string `json:"image"`

	// port is the container port; it is also the Service target port.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`

	// env is the container's environment.
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// resources are the container's compute resource requirements.
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// scaling selects a fixed replica count or autoscaling. Defaults to 1 replica.
	// +optional
	Scaling Scaling `json:"scaling,omitempty"`

	// expose configures how the workload is reachable. Defaults to a ClusterIP Service.
	// +optional
	Expose Expose `json:"expose,omitempty"`

	// healthCheck wires HTTP liveness and readiness probes to the container.
	// +optional
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

	// storage declares what the workload keeps. Omitted, the kube is stateless.
	// +optional
	Storage *Storage `json:"storage,omitempty"`
}

// KubeContainerStatus defines the observed state of KubeContainer.
type KubeContainerStatus struct {
	// observedGeneration is the spec generation most recently acted on.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// availableReplicas mirrors the owned Deployment's available replicas.
	// +optional
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`

	// endpoint is the address at which the workload is reachable.
	// +optional
	Endpoint string `json:"endpoint,omitempty"`

	// conditions represent the current state of the KubeContainer resource:
	// Ready, Progressing, and Degraded.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Image",type=string,JSONPath=`.spec.image`
// +kubebuilder:printcolumn:name="Available",type=integer,JSONPath=`.status.availableReplicas`
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.conditions[?(@.type=="Ready")].status`
// +kubebuilder:printcolumn:name="Endpoint",type=string,JSONPath=`.status.endpoint`

// KubeContainer is the Schema for the kubecontainers API
type KubeContainer struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of KubeContainer
	// +required
	Spec KubeContainerSpec `json:"spec"`

	// status defines the observed state of KubeContainer
	// +optional
	Status KubeContainerStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// KubeContainerList contains a list of KubeContainer
type KubeContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []KubeContainer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeContainer{}, &KubeContainerList{})
}
