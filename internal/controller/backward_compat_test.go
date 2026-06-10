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
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/yaml"

	kubecontainerv1alpha1 "github.com/unboxd-cloud/kubecontainer/api/v1alpha1"
)

// The manifests under testdata/compat are a frozen, append-only corpus:
// every file is a declaration that was valid in a released era and must
// remain valid and convergent forever. A failure here means a breaking
// change to the published contract — treat it as such, never by editing
// the corpus.
var _ = Describe("Backward compatibility golden corpus", func() {
	ctx := context.Background()

	files, err := filepath.Glob(filepath.Join("testdata", "compat", "*.yaml"))
	Expect(err).NotTo(HaveOccurred())
	Expect(files).NotTo(BeEmpty(), "compat corpus must not be empty")

	for _, file := range files {
		It("still accepts and converges "+filepath.Base(file), func() {
			data, err := os.ReadFile(file)
			Expect(err).NotTo(HaveOccurred())

			kc := &kubecontainerv1alpha1.KubeContainer{}
			Expect(yaml.UnmarshalStrict(data, kc)).To(Succeed(),
				"golden manifest must decode without unknown or dropped fields")
			kc.Namespace = "default"

			Expect(k8sClient.Create(ctx, kc)).To(Succeed(),
				"golden manifest must remain valid against the current CRD schema")
			name := types.NamespacedName{Name: kc.Name, Namespace: kc.Namespace}
			DeferCleanup(func() {
				_ = k8sClient.Delete(ctx, kc)
				for _, obj := range []client.Object{
					&appsv1.Deployment{}, &corev1.Service{},
					&networkingv1.Ingress{}, &autoscalingv2.HorizontalPodAutoscaler{},
				} {
					obj.SetName(kc.Name)
					obj.SetNamespace(kc.Namespace)
					if err := k8sClient.Delete(ctx, obj); err != nil {
						Expect(apierrors.IsNotFound(err)).To(BeTrue())
					}
				}
			})

			reconciler := &KubeContainerReconciler{
				Client:   k8sClient,
				Scheme:   k8sClient.Scheme(),
				Recorder: events.NewFakeRecorder(32),
			}
			_, err = reconciler.Reconcile(ctx, reconcile.Request{NamespacedName: name})
			Expect(err).NotTo(HaveOccurred(),
				"golden manifest must still reconcile without error")

			deploy := &appsv1.Deployment{}
			Expect(k8sClient.Get(ctx, name, deploy)).To(Succeed(),
				"golden manifest must still produce its Deployment")
			Expect(deploy.Spec.Template.Spec.Containers[0].Image).
				To(Equal(kc.Spec.Image))

			if kc.Spec.Expose.Type == kubecontainerv1alpha1.ExposeIngress {
				Expect(k8sClient.Get(ctx, name, &networkingv1.Ingress{})).To(Succeed())
			}
			if kc.Spec.Scaling.Autoscale != nil {
				Expect(k8sClient.Get(ctx, name, &autoscalingv2.HorizontalPodAutoscaler{})).To(Succeed())
			}
		})
	}
})
