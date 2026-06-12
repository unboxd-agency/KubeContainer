// DeployRehearsal — deploy the declared kube into the rehearsal chamber:
// a real control plane (envtest), the real CRD, the real reconciler, and
// the real declaration from deploy/. The venv law, executed: compile and
// simulate in the venv; certify in the cloud; deliver at the ports.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/yaml"

	kubecontainerv1alpha1 "github.com/unboxd-cloud/kubecontainer/api/v1alpha1"
	"github.com/unboxd-cloud/kubecontainer/internal/controller"
)

func main() {
	ctrl.SetLogger(zap.New(zap.UseDevMode(false)))
	scheme := runtime.NewScheme()
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(kubecontainerv1alpha1.AddToScheme(scheme))

	env := &envtest.Environment{CRDDirectoryPaths: []string{"config/crd/bases"}}
	cfg, err := env.Start()
	if err != nil {
		fmt.Println("[fail] control plane:", err)
		os.Exit(1)
	}
	defer env.Stop()
	fmt.Println("[pass] control plane up (rehearsal chamber: real kube-apiserver + etcd)")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{Scheme: scheme,
		Metrics: metricsserver.Options{BindAddress: "0"}, HealthProbeBindAddress: "0"})
	if err != nil {
		fmt.Println("[fail] manager:", err)
		os.Exit(1)
	}
	if err := (&controller.KubeContainerReconciler{
		Client: mgr.GetClient(), Scheme: mgr.GetScheme(),
		Recorder: mgr.GetEventRecorder("kubecontainer-controller"),
	}).SetupWithManager(mgr); err != nil {
		fmt.Println("[fail] reconciler:", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() { _ = mgr.Start(ctx) }()
	if !mgr.GetCache().WaitForCacheSync(ctx) {
		fmt.Println("[fail] cache sync")
		os.Exit(1)
	}
	fmt.Println("[pass] operator running (the loop is live)")

	raw, err := os.ReadFile("deploy/arithmetic-kube.yaml")
	if err != nil {
		fmt.Println("[fail] declaration:", err)
		os.Exit(1)
	}
	kc := &kubecontainerv1alpha1.KubeContainer{}
	if err := yaml.UnmarshalStrict(raw, kc); err != nil {
		fmt.Println("[fail] declaration inadmissible:", err)
		os.Exit(1)
	}
	kc.Namespace = "default"
	if err := mgr.GetClient().Create(ctx, kc); err != nil {
		fmt.Println("[fail] admission:", err)
		os.Exit(1)
	}
	fmt.Println("[pass] declaration admitted: KubeContainer/arithmetic (CEL-validated at the gate)")

	key := types.NamespacedName{Name: kc.Name, Namespace: "default"}
	var dep appsv1.Deployment
	var svc corev1.Service
	deadline := time.Now().Add(30 * time.Second)
	haveDep, haveSvc := false, false
	for time.Now().Before(deadline) && !(haveDep && haveSvc) {
		if !haveDep && mgr.GetClient().Get(ctx, key, &dep) == nil {
			haveDep = true
			fmt.Printf("[pass] child converged: Deployment/%s (image %s)\n",
				dep.Name, dep.Spec.Template.Spec.Containers[0].Image)
		}
		if !haveSvc && mgr.GetClient().Get(ctx, key, &svc) == nil {
			haveSvc = true
			fmt.Printf("[pass] child converged: Service/%s (port %d)\n",
				svc.Name, svc.Spec.Ports[0].Port)
		}
		time.Sleep(200 * time.Millisecond)
	}
	if !haveDep || !haveSvc {
		fmt.Println("[fail] children did not converge in 30s")
		os.Exit(1)
	}
	got := &kubecontainerv1alpha1.KubeContainer{}
	_ = mgr.GetClient().Get(ctx, key, got)
	for _, c := range got.Status.Conditions {
		fmt.Printf("[info] condition %s=%s (%s)\n", c.Type, c.Status, c.Reason)
	}
	fmt.Println("verdict: the kube is deployed in the rehearsal chamber — declared, admitted, converged, recorded")
	fmt.Println("note: Ready follows real pods; envtest runs no kubelet — the real-traffic verdict belongs to the e2e gate and any conformant cluster (kubectl apply -f deploy/arithmetic-kube.yaml)")
}
