// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package capi

import (
	cacpt "github.com/siderolabs/cluster-api-control-plane-provider-talos/api/v1alpha3"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	capiv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	sidero "github.com/siderolabs/sidero/app/caps-controller-manager/api/v1alpha3"
	metal "github.com/siderolabs/sidero/app/sidero-controller-manager/api/v1alpha1"
)

// GetMetalClient builds k8s client with schemes required to access all the CAPI/Sidero/Talos components.
func GetMetalClient(config *rest.Config) (runtimeclient.Client, error) {
	scheme := runtime.NewScheme()

	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		return nil, err
	}

	if err := capiv1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	if err := cacpt.AddToScheme(scheme); err != nil {
		return nil, err
	}

	if err := sidero.AddToScheme(scheme); err != nil {
		return nil, err
	}

	if err := metal.AddToScheme(scheme); err != nil {
		return nil, err
	}

	return runtimeclient.New(config, runtimeclient.Options{Scheme: scheme})
}
