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

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/ptr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	kubecontainerv1alpha1 "github.com/unboxd-cloud/kubecontainer/api/v1alpha1"
)

// KubeContainerReconciler reconciles a KubeContainer object
type KubeContainerReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=kubecontainer.unboxd.cloud,resources=kubecontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=kubecontainer.unboxd.cloud,resources=kubecontainers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=kubecontainer.unboxd.cloud,resources=kubecontainers/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=autoscaling,resources=horizontalpodautoscalers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch

// Reconcile converges the cluster state for a single KubeContainer: it
// creates or updates the owned Deployment, Service, Ingress, and HPA to match
// spec, removes children that the spec no longer calls for, and reports the
// result in status conditions.
func (r *KubeContainerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	kc := &kubecontainerv1alpha1.KubeContainer{}
	if err := r.Get(ctx, req.NamespacedName, kc); err != nil {
		// Children are garbage-collected via owner references; nothing to do.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if err := r.reconcileChildren(ctx, kc); err != nil {
		log.Error(err, "failed to reconcile children")
		r.Recorder.Event(kc, corev1.EventTypeWarning, "ReconcileError", err.Error())
		meta.SetStatusCondition(&kc.Status.Conditions, metav1.Condition{
			Type:               kubecontainerv1alpha1.ConditionDegraded,
			Status:             metav1.ConditionTrue,
			Reason:             "ReconcileError",
			Message:            err.Error(),
			ObservedGeneration: kc.Generation,
		})
		if statusErr := r.Status().Update(ctx, kc); statusErr != nil {
			log.Error(statusErr, "failed to update status after reconcile error")
		}
		return ctrl.Result{}, err
	}

	if err := r.updateStatus(ctx, kc); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// reconcileChildren creates or updates the owned resources and deletes the
// ones the current spec no longer calls for.
func (r *KubeContainerReconciler) reconcileChildren(ctx context.Context, kc *kubecontainerv1alpha1.KubeContainer) error {
	log := logf.FromContext(ctx)

	deploy := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{Name: kc.Name, Namespace: kc.Namespace}}
	op, err := controllerutil.CreateOrUpdate(ctx, r.Client, deploy, func() error {
		r.buildDeployment(kc, deploy)
		return controllerutil.SetControllerReference(kc, deploy, r.Scheme)
	})
	if err != nil {
		return fmt.Errorf("deployment: %w", err)
	}
	logChildOp(log, r.Recorder, kc, "Deployment", op)

	svc := &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: kc.Name, Namespace: kc.Namespace}}
	op, err = controllerutil.CreateOrUpdate(ctx, r.Client, svc, func() error {
		buildService(kc, svc)
		return controllerutil.SetControllerReference(kc, svc, r.Scheme)
	})
	if err != nil {
		return fmt.Errorf("service: %w", err)
	}
	logChildOp(log, r.Recorder, kc, "Service", op)

	if kc.Spec.Expose.Type == kubecontainerv1alpha1.ExposeIngress {
		ing := &networkingv1.Ingress{ObjectMeta: metav1.ObjectMeta{Name: kc.Name, Namespace: kc.Namespace}}
		op, err = controllerutil.CreateOrUpdate(ctx, r.Client, ing, func() error {
			buildIngress(kc, ing)
			return controllerutil.SetControllerReference(kc, ing, r.Scheme)
		})
		if err != nil {
			return fmt.Errorf("ingress: %w", err)
		}
		logChildOp(log, r.Recorder, kc, "Ingress", op)
	} else if err := r.deleteIfExists(ctx, kc, &networkingv1.Ingress{}); err != nil {
		return fmt.Errorf("deleting orphaned ingress: %w", err)
	}

	if kc.Spec.Scaling.Autoscale != nil {
		hpa := &autoscalingv2.HorizontalPodAutoscaler{ObjectMeta: metav1.ObjectMeta{Name: kc.Name, Namespace: kc.Namespace}}
		op, err = controllerutil.CreateOrUpdate(ctx, r.Client, hpa, func() error {
			buildHPA(kc, hpa)
			return controllerutil.SetControllerReference(kc, hpa, r.Scheme)
		})
		if err != nil {
			return fmt.Errorf("hpa: %w", err)
		}
		logChildOp(log, r.Recorder, kc, "HorizontalPodAutoscaler", op)
	} else if err := r.deleteIfExists(ctx, kc, &autoscalingv2.HorizontalPodAutoscaler{}); err != nil {
		return fmt.Errorf("deleting orphaned hpa: %w", err)
	}

	return nil
}

// updateStatus mirrors the owned Deployment's availability into the CR's
// conditions and computes the reachable endpoint.
func (r *KubeContainerReconciler) updateStatus(ctx context.Context, kc *kubecontainerv1alpha1.KubeContainer) error {
	deploy := &appsv1.Deployment{}
	if err := r.Get(ctx, types.NamespacedName{Name: kc.Name, Namespace: kc.Namespace}, deploy); err != nil {
		return err
	}
	svc := &corev1.Service{}
	if err := r.Get(ctx, types.NamespacedName{Name: kc.Name, Namespace: kc.Namespace}, svc); err != nil {
		return err
	}

	kc.Status.ObservedGeneration = kc.Generation
	kc.Status.AvailableReplicas = deploy.Status.AvailableReplicas
	kc.Status.Endpoint = endpoint(kc, svc)

	desired := int32(1)
	if deploy.Spec.Replicas != nil {
		desired = *deploy.Spec.Replicas
	}
	ready := deploy.Status.AvailableReplicas >= desired && deploy.Status.ObservedGeneration >= deploy.Generation

	readyCond := metav1.Condition{
		Type:               kubecontainerv1alpha1.ConditionReady,
		Status:             metav1.ConditionFalse,
		Reason:             "DeploymentUnavailable",
		Message:            fmt.Sprintf("%d/%d replicas available", deploy.Status.AvailableReplicas, desired),
		ObservedGeneration: kc.Generation,
	}
	progressingCond := metav1.Condition{
		Type:               kubecontainerv1alpha1.ConditionProgressing,
		Status:             metav1.ConditionTrue,
		Reason:             "RolloutInProgress",
		Message:            "waiting for the deployment to become available",
		ObservedGeneration: kc.Generation,
	}
	if ready {
		readyCond.Status = metav1.ConditionTrue
		readyCond.Reason = "DeploymentAvailable"
		progressingCond.Status = metav1.ConditionFalse
		progressingCond.Reason = "RolloutComplete"
		progressingCond.Message = "the deployment is available"
	}
	meta.SetStatusCondition(&kc.Status.Conditions, readyCond)
	meta.SetStatusCondition(&kc.Status.Conditions, progressingCond)
	meta.SetStatusCondition(&kc.Status.Conditions, metav1.Condition{
		Type:               kubecontainerv1alpha1.ConditionDegraded,
		Status:             metav1.ConditionFalse,
		Reason:             "ReconcileSucceeded",
		ObservedGeneration: kc.Generation,
	})

	return r.Status().Update(ctx, kc)
}

func (r *KubeContainerReconciler) buildDeployment(kc *kubecontainerv1alpha1.KubeContainer, deploy *appsv1.Deployment) {
	labels := selectorLabels(kc)
	deploy.Labels = labels
	deploy.Spec.Selector = &metav1.LabelSelector{MatchLabels: labels}
	// Under autoscaling the HPA owns the replica count, so leave it untouched.
	if kc.Spec.Scaling.Autoscale == nil {
		deploy.Spec.Replicas = ptr.To(int32(1))
		if kc.Spec.Scaling.Replicas != nil {
			deploy.Spec.Replicas = ptr.To(*kc.Spec.Scaling.Replicas)
		}
	}

	container := corev1.Container{
		Name:      "workload",
		Image:     kc.Spec.Image,
		Ports:     []corev1.ContainerPort{{Name: "http", ContainerPort: kc.Spec.Port, Protocol: corev1.ProtocolTCP}},
		Env:       kc.Spec.Env,
		Resources: kc.Spec.Resources,
	}
	if hc := kc.Spec.HealthCheck; hc != nil {
		probe := &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{Path: hc.Path, Port: intstr.FromInt32(kc.Spec.Port)},
			},
		}
		container.LivenessProbe = probe.DeepCopy()
		container.ReadinessProbe = probe.DeepCopy()
	}
	deploy.Spec.Template.Labels = labels
	deploy.Spec.Template.Spec.Containers = []corev1.Container{container}
}

func buildService(kc *kubecontainerv1alpha1.KubeContainer, svc *corev1.Service) {
	svc.Labels = selectorLabels(kc)
	svc.Spec.Selector = selectorLabels(kc)
	svc.Spec.Type = corev1.ServiceTypeClusterIP
	if kc.Spec.Expose.Type == kubecontainerv1alpha1.ExposeLoadBalancer {
		svc.Spec.Type = corev1.ServiceTypeLoadBalancer
	}
	svc.Spec.Ports = []corev1.ServicePort{{
		Name:       "http",
		Port:       kc.Spec.Port,
		TargetPort: intstr.FromInt32(kc.Spec.Port),
		Protocol:   corev1.ProtocolTCP,
	}}
}

func buildIngress(kc *kubecontainerv1alpha1.KubeContainer, ing *networkingv1.Ingress) {
	ing.Labels = selectorLabels(kc)
	pathType := networkingv1.PathTypePrefix
	ing.Spec.Rules = []networkingv1.IngressRule{{
		Host: kc.Spec.Expose.Host,
		IngressRuleValue: networkingv1.IngressRuleValue{
			HTTP: &networkingv1.HTTPIngressRuleValue{
				Paths: []networkingv1.HTTPIngressPath{{
					Path:     "/",
					PathType: &pathType,
					Backend: networkingv1.IngressBackend{
						Service: &networkingv1.IngressServiceBackend{
							Name: kc.Name,
							Port: networkingv1.ServiceBackendPort{Number: kc.Spec.Port},
						},
					},
				}},
			},
		},
	}}
}

func buildHPA(kc *kubecontainerv1alpha1.KubeContainer, hpa *autoscalingv2.HorizontalPodAutoscaler) {
	hpa.Labels = selectorLabels(kc)
	as := kc.Spec.Scaling.Autoscale
	hpa.Spec.ScaleTargetRef = autoscalingv2.CrossVersionObjectReference{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Name:       kc.Name,
	}
	hpa.Spec.MinReplicas = ptr.To(as.MinReplicas)
	hpa.Spec.MaxReplicas = as.MaxReplicas
	target := as.TargetCPUUtilization
	if target == 0 {
		target = 80
	}
	hpa.Spec.Metrics = []autoscalingv2.MetricSpec{{
		Type: autoscalingv2.ResourceMetricSourceType,
		Resource: &autoscalingv2.ResourceMetricSource{
			Name: corev1.ResourceCPU,
			Target: autoscalingv2.MetricTarget{
				Type:               autoscalingv2.UtilizationMetricType,
				AverageUtilization: ptr.To(target),
			},
		},
	}}
}

// deleteIfExists removes the same-named child of the given type, ignoring
// objects that are already gone.
func (r *KubeContainerReconciler) deleteIfExists(ctx context.Context, kc *kubecontainerv1alpha1.KubeContainer, obj client.Object) error {
	obj.SetName(kc.Name)
	obj.SetNamespace(kc.Namespace)
	if err := r.Delete(ctx, obj); err != nil && !apierrors.IsNotFound(err) {
		return err
	}
	return nil
}

func endpoint(kc *kubecontainerv1alpha1.KubeContainer, svc *corev1.Service) string {
	switch kc.Spec.Expose.Type {
	case kubecontainerv1alpha1.ExposeIngress:
		return kc.Spec.Expose.Host
	case kubecontainerv1alpha1.ExposeLoadBalancer:
		for _, ing := range svc.Status.LoadBalancer.Ingress {
			host := ing.Hostname
			if host == "" {
				host = ing.IP
			}
			if host != "" {
				return fmt.Sprintf("%s:%d", host, kc.Spec.Port)
			}
		}
		return "" // LB address not yet provisioned
	default:
		return fmt.Sprintf("%s.%s.svc.cluster.local:%d", kc.Name, kc.Namespace, kc.Spec.Port)
	}
}

func selectorLabels(kc *kubecontainerv1alpha1.KubeContainer) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       kc.Name,
		"app.kubernetes.io/managed-by": "kubecontainer",
	}
}

func logChildOp(log logr.Logger, recorder record.EventRecorder, kc *kubecontainerv1alpha1.KubeContainer, kind string, op controllerutil.OperationResult) {
	if op == controllerutil.OperationResultNone {
		return
	}
	log.Info("reconciled child", "kind", kind, "operation", op)
	recorder.Eventf(kc, corev1.EventTypeNormal, "Reconciled", "%s %s", kind, op)
}

// SetupWithManager sets up the controller with the Manager.
func (r *KubeContainerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kubecontainerv1alpha1.KubeContainer{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&networkingv1.Ingress{}).
		Owns(&autoscalingv2.HorizontalPodAutoscaler{}).
		Named("kubecontainer").
		Complete(r)
}
