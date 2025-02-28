// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkServicesMeshSpec struct {
	/* Optional. A free-text description of the resource. Max length 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to be redirected to this port regardless of its actual ip:port destination. If unset, a port '15001' is used as the interception port. This field is only valid if the type of Mesh is SIDECAR. */
	// +optional
	InterceptionPort *int `json:"interceptionPort,omitempty"`

	/* Optional. Set of label tags associated with the Mesh resource. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type NetworkServicesMeshStatus struct {
	/* Conditions represent the latest available observations of the
	   NetworkServicesMesh's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The timestamp when the resource was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. Server-defined URL of this resource */
	SelfLink string `json:"selfLink,omitempty"`
	/* Output only. The timestamp when the resource was updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesMesh is the Schema for the networkservices API
// +k8s:openapi-gen=true
type NetworkServicesMesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkServicesMeshSpec   `json:"spec,omitempty"`
	Status NetworkServicesMeshStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkServicesMeshList contains a list of NetworkServicesMesh
type NetworkServicesMeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesMesh `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesMesh{}, &NetworkServicesMeshList{})
}
