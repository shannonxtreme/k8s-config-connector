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

package resourceoverrides

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// CRDDecorate decorates the given CRD to ensure that its schemas are authored correctly.
// It could be used to preserve legacy fields, to mark fields optional with defaults, etc.
type CRDDecorate func(crd *apiextensions.CustomResourceDefinition) error

// PreActuationTransform transforms the original spec to the golden format that the resource actuator can understand.
// For example, it could be used to fetch value from the legacy field and place it to the field that the resource actuator actually understands and supports.
type PreActuationTransform func(r *k8s.Resource) error

// PostActuationTransform transform the reconciled resource object.
// A typical example of post-actuation transformations is to preserve the user specified fields.
type PostActuationTransform func(original, reconciled *k8s.Resource) error

// ConfigValidate validates the input configuration in the webhook.
type ConfigValidate func(r *unstructured.Unstructured) error

// ResourceOverride holds all pieces of changes needed, i.e. decoration, transformation and validation to author
// a resource-specific behavior override.
// Since one particular resource kind could have multiple overrides, each ResourceOverride should be logically orthogonal to each other and neutral to order of execution.
type ResourceOverride struct {
	CRDDecorate            CRDDecorate
	ConfigValidate         ConfigValidate
	PreActuationTransform  PreActuationTransform
	PostActuationTransform PostActuationTransform
}

type ResourceOverrides struct {
	Kind      string
	Overrides []ResourceOverride
}

type ResourceOverridesHandler struct {
	overridesPerKindMap map[string]ResourceOverrides
}

func NewResourceOverridesHandler() *ResourceOverridesHandler {
	return &ResourceOverridesHandler{
		overridesPerKindMap: make(map[string]ResourceOverrides),
	}
}

var Handler = NewResourceOverridesHandler()

func (h *ResourceOverridesHandler) CRDDecorate(crd *apiextensions.CustomResourceDefinition) error {
	kind := crd.Spec.Names.Kind
	ro, found := h.registration(kind)
	if !found {
		return nil
	}
	for _, o := range ro.Overrides {
		if o.CRDDecorate != nil {
			if err := o.CRDDecorate(crd); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *ResourceOverridesHandler) ConfigValidate(r *unstructured.Unstructured) error {
	kind := r.GetKind()
	ro, found := h.registration(kind)
	if !found {
		return nil
	}
	for _, o := range ro.Overrides {
		if o.ConfigValidate != nil {
			if err := o.ConfigValidate(r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *ResourceOverridesHandler) PreActuationTransform(r *k8s.Resource) error {
	ro, found := h.registration(r.Kind)
	if !found {
		return nil
	}
	for _, o := range ro.Overrides {
		if o.PreActuationTransform != nil {
			if err := o.PreActuationTransform(r); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *ResourceOverridesHandler) PostActuationTransform(original, post *k8s.Resource) error {
	ro, found := h.registration(original.Kind)
	if !found {
		return nil
	}
	for _, o := range ro.Overrides {
		if o.PostActuationTransform != nil {
			if err := o.PostActuationTransform(original, post); err != nil {
				return err
			}
		}
	}
	return nil
}

func (h *ResourceOverridesHandler) HasOverrides(kind string) bool {
	_, found := h.registration(kind)
	return found
}

func (h *ResourceOverridesHandler) HasConfigValidate(kind string) bool {
	ro, found := h.registration(kind)
	if !found {
		return false
	}
	for _, o := range ro.Overrides {
		if o.ConfigValidate != nil {
			return true
		}
	}
	return false
}

func (h *ResourceOverridesHandler) registration(kind string) (*ResourceOverrides, bool) {
	ro, found := h.overridesPerKindMap[kind]
	if !found {
		return nil, false
	}
	return &ro, found
}

func (h *ResourceOverridesHandler) Register(ro ResourceOverrides) {
	h.overridesPerKindMap[ro.Kind] = ro
}

func init() {
	Handler.Register(GetStorageBucketResourceOverrides())
	Handler.Register(GetSQLInstanceResourceOverrides())
	Handler.Register(GetContainerClusterResourceOverrides())
	Handler.Register(GetLoggingLogSinkResourceOverrides())
	Handler.Register(GetComputeInstanceResourceOverrides())
	Handler.Register(GetDNSRecordSetOverrides())
	Handler.Register(GetComputeBackendServiceResourceOverrides())
}
