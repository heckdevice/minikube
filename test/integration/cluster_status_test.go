// +build integration

/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package integration

import (
	"fmt"
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/minikube/test/integration/util"
)

func TestCluster(t *testing.T) {
	minikubeRunner := util.MinikubeRunner{BinaryPath: *binaryPath, T: t}
	minikubeRunner.RunCommand("start", true)
	minikubeRunner.CheckStatus("Running")

	kubectlRunner := util.NewKubectlRunner(t)
	cs := api.ComponentStatusList{}
	kubectlRunner.RunCommand([]string{"get", "cs"}, &cs)

	for _, i := range cs.Items {
		status := api.ConditionFalse
		for _, c := range i.Conditions {
			if c.Type != api.ComponentHealthy {
				continue
			}
			fmt.Printf("Component: %s, Healthy: %s.\n", i.GetName(), c.Status)
			status = c.Status
		}
		if status != api.ConditionTrue {
			t.Fatalf("Component %s is not Healthy! Status: %s", i.GetName(), status)
		}
	}
}
