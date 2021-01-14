// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

package main

import (
	"flag"
	"github.com/DataDog/watermarkpodautoscaler/pkg/config"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	datadoghqv1alpha1 "github.com/DataDog/watermarkpodautoscaler/api/v1alpha1"
	"github.com/DataDog/watermarkpodautoscaler/controllers"
	// +kubebuilder:scaffold:imports
)

var (
	scheme            = runtime.NewScheme()
	setupLog          = ctrl.Log.WithName("setup")
	host              = "0.0.0.0"
	metricsPort int32 = 8383
	healthPort        = 9440
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(datadoghqv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var printVersionArg bool
	flag.BoolVar(&printVersionArg, "version", printVersionArg, "print version")
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", true,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.IntVar(&healthPort, "health-port", healthPort, "Port to use for the health probe")

	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), config.ManagerOptionsWithNamespaces(setupLog, ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: fmt.Sprintf("%s:%d", host, metricsPort),
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "watermarkpodautoscaler-lock",
		HealthProbeBindAddress: fmt.Sprintf("%s:%d", host, healthPort),
	}))
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.WatermarkPodAutoscalerReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("WatermarkPodAutoscaler"),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "WatermarkPodAutoscaler")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
