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

package controller

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	kubecontainerv1alpha1 "github.com/unboxd-cloud/kubecontainer/api/v1alpha1"
)

var _ = Describe("KubeContainer Controller", func() {
	const resourceName = "test-resource"

	ctx := context.Background()

	typeNamespacedName := types.NamespacedName{
		Name:      resourceName,
		Namespace: "default",
	}

	var reconciler *KubeContainerReconciler

	doReconcile := func() {
		GinkgoHelper()
		_, err := reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: typeNamespacedName})
		Expect(err).NotTo(HaveOccurred())
	}

	BeforeEach(func() {
		reconciler = &KubeContainerReconciler{
			Client:   k8sClient,
			Scheme:   k8sClient.Scheme(),
			Recorder: record.NewFakeRecorder(32),
		}

		resource := &kubecontainerv1alpha1.KubeContainer{
			ObjectMeta: metav1.ObjectMeta{
				Name:      resourceName,
				Namespace: "default",
			},
			Spec: kubecontainerv1alpha1.KubeContainerSpec{
				Image: "ghcr.io/acme/my-app:1.0.0",
				Port:  8080,
				Scaling: kubecontainerv1alpha1.Scaling{
					Replicas: ptr.To(int32(2)),
				},
				HealthCheck: &kubecontainerv1alpha1.HealthCheck{Path: "/healthz"},
			},
		}
		Expect(k8sClient.Create(ctx, resource)).To(Succeed())
	})

	AfterEach(func() {
		resource := &kubecontainerv1alpha1.KubeContainer{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, resource)).To(Succeed())
		Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		// envtest runs no garbage collector, so remove children explicitly.
		for _, obj := range []client.Object{
			&appsv1.Deployment{}, &corev1.Service{}, &networkingv1.Ingress{}, &autoscalingv2.HorizontalPodAutoscaler{},
		} {
			obj.SetName(resourceName)
			obj.SetNamespace("default")
			if err := k8sClient.Delete(ctx, obj); err != nil {
				Expect(apierrors.IsNotFound(err)).To(BeTrue())
			}
		}
	})

	It("creates a Deployment and ClusterIP Service from the spec", func() {
		doReconcile()

		deploy := &appsv1.Deployment{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, deploy)).To(Succeed())
		Expect(deploy.Spec.Replicas).To(HaveValue(BeEquivalentTo(2)))
		Expect(deploy.OwnerReferences).To(HaveLen(1))
		container := deploy.Spec.Template.Spec.Containers[0]
		Expect(container.Image).To(Equal("ghcr.io/acme/my-app:1.0.0"))
		Expect(container.Ports[0].ContainerPort).To(BeEquivalentTo(8080))
		Expect(container.LivenessProbe.HTTPGet.Path).To(Equal("/healthz"))
		Expect(container.ReadinessProbe.HTTPGet.Path).To(Equal("/healthz"))

		svc := &corev1.Service{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, svc)).To(Succeed())
		Expect(svc.Spec.Type).To(Equal(corev1.ServiceTypeClusterIP))
		Expect(svc.Spec.Ports[0].Port).To(BeEquivalentTo(8080))
	})

	It("reports status conditions and the endpoint", func() {
		doReconcile()

		kc := &kubecontainerv1alpha1.KubeContainer{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, kc)).To(Succeed())
		Expect(kc.Status.ObservedGeneration).To(Equal(kc.Generation))
		Expect(kc.Status.Endpoint).To(Equal(fmt.Sprintf("%s.default.svc.cluster.local:8080", resourceName)))
		// envtest runs no deployment controller, so no replicas become available.
		Expect(meta.IsStatusConditionTrue(kc.Status.Conditions, kubecontainerv1alpha1.ConditionProgressing)).To(BeTrue())
		Expect(meta.IsStatusConditionFalse(kc.Status.Conditions, kubecontainerv1alpha1.ConditionReady)).To(BeTrue())
		Expect(meta.IsStatusConditionFalse(kc.Status.Conditions, kubecontainerv1alpha1.ConditionDegraded)).To(BeTrue())
	})

	It("creates an Ingress when exposed via Ingress and removes it when switched back", func() {
		kc := &kubecontainerv1alpha1.KubeContainer{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, kc)).To(Succeed())
		kc.Spec.Expose = kubecontainerv1alpha1.Expose{
			Type: kubecontainerv1alpha1.ExposeIngress,
			Host: "my-app.example.com",
		}
		Expect(k8sClient.Update(ctx, kc)).To(Succeed())
		doReconcile()

		ing := &networkingv1.Ingress{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, ing)).To(Succeed())
		Expect(ing.Spec.Rules[0].Host).To(Equal("my-app.example.com"))
		backend := ing.Spec.Rules[0].HTTP.Paths[0].Backend.Service
		Expect(backend.Name).To(Equal(resourceName))
		Expect(backend.Port.Number).To(BeEquivalentTo(8080))

		Expect(k8sClient.Get(ctx, typeNamespacedName, kc)).To(Succeed())
		Expect(kc.Status.Endpoint).To(Equal("my-app.example.com"))

		kc.Spec.Expose = kubecontainerv1alpha1.Expose{Type: kubecontainerv1alpha1.ExposeClusterIP}
		Expect(k8sClient.Update(ctx, kc)).To(Succeed())
		doReconcile()

		err := k8sClient.Get(ctx, typeNamespacedName, &networkingv1.Ingress{})
		Expect(apierrors.IsNotFound(err)).To(BeTrue())
	})

	It("creates an HPA when autoscaling and leaves the replica count to it", func() {
		kc := &kubecontainerv1alpha1.KubeContainer{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, kc)).To(Succeed())
		kc.Spec.Scaling = kubecontainerv1alpha1.Scaling{
			Autoscale: &kubecontainerv1alpha1.Autoscale{
				MinReplicas:          2,
				MaxReplicas:          10,
				TargetCPUUtilization: 75,
			},
		}
		Expect(k8sClient.Update(ctx, kc)).To(Succeed())
		doReconcile()

		hpa := &autoscalingv2.HorizontalPodAutoscaler{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, hpa)).To(Succeed())
		Expect(hpa.Spec.MinReplicas).To(HaveValue(BeEquivalentTo(2)))
		Expect(hpa.Spec.MaxReplicas).To(BeEquivalentTo(10))
		Expect(hpa.Spec.ScaleTargetRef.Name).To(Equal(resourceName))
		Expect(hpa.Spec.Metrics[0].Resource.Target.AverageUtilization).To(HaveValue(BeEquivalentTo(75)))

		// Simulate the HPA scaling the deployment, then verify a reconcile
		// does not fight it.
		deploy := &appsv1.Deployment{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, deploy)).To(Succeed())
		deploy.Spec.Replicas = ptr.To(int32(7))
		Expect(k8sClient.Update(ctx, deploy)).To(Succeed())
		doReconcile()
		Expect(k8sClient.Get(ctx, typeNamespacedName, deploy)).To(Succeed())
		Expect(deploy.Spec.Replicas).To(HaveValue(BeEquivalentTo(7)))

		kc = &kubecontainerv1alpha1.KubeContainer{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, kc)).To(Succeed())
		kc.Spec.Scaling = kubecontainerv1alpha1.Scaling{Replicas: ptr.To(int32(3))}
		Expect(k8sClient.Update(ctx, kc)).To(Succeed())
		doReconcile()

		err := k8sClient.Get(ctx, typeNamespacedName, &autoscalingv2.HorizontalPodAutoscaler{})
		Expect(apierrors.IsNotFound(err)).To(BeTrue())
		Expect(k8sClient.Get(ctx, typeNamespacedName, deploy)).To(Succeed())
		Expect(deploy.Spec.Replicas).To(HaveValue(BeEquivalentTo(3)))
	})

	It("rejects a spec with both replicas and autoscale", func() {
		kc := &kubecontainerv1alpha1.KubeContainer{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, kc)).To(Succeed())
		kc.Spec.Scaling.Autoscale = &kubecontainerv1alpha1.Autoscale{MinReplicas: 1, MaxReplicas: 5}
		err := k8sClient.Update(ctx, kc)
		Expect(err).To(MatchError(ContainSubstring("mutually exclusive")))
	})

	It("rejects an Ingress expose without a host", func() {
		kc := &kubecontainerv1alpha1.KubeContainer{}
		Expect(k8sClient.Get(ctx, typeNamespacedName, kc)).To(Succeed())
		kc.Spec.Expose = kubecontainerv1alpha1.Expose{Type: kubecontainerv1alpha1.ExposeIngress}
		err := k8sClient.Update(ctx, kc)
		Expect(err).To(MatchError(ContainSubstring("host is required")))
	})
})
