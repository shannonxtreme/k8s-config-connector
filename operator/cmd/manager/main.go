// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	goflag "flag"
	"os"

	corev1v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/apis/core/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers/configconnector"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/controllers/configconnectorcontext"
	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/pkg/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/kccmanager/nocache"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp/profiler"

	flag "github.com/spf13/pflag"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = apiextensions.SchemeBuilder.AddToScheme(scheme)

	_ = corev1v1beta1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var repoPath string

	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	profiler.AddFlag(flag.CommandLine)
	flag.StringVar(&repoPath, "local-repo", "./channels", "location of local repository to use")
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(logging.BuildLogger(os.Stderr))
	addon.Init()

	if err := profiler.StartIfEnabled(); err != nil {
		setupLog.Error(err, "problem running profiler")
		os.Exit(1)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		Port:               9443,
		// Disable the caching for the client. The cached reader will lazily list structured resources cross namespaces.
		// The operator mostly only cares about resources in cnrm-system namespace.
		NewClient: nocache.NoCacheClientFunc,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = configconnector.Add(mgr, repoPath); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ConfigConnector")
		os.Exit(1)
	}

	if err = configconnectorcontext.Add(mgr, repoPath); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ConfigConnectorContext")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
