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

package webhook

import (
	"context"
	"fmt"
	"net/http"

	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	"github.com/golang/glog"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	apimachinerytypes "k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type managementConflictAnnotationDefaulter struct {
	client                client.Client
	dclSchemaLoader       dclschemaloader.DCLSchemaLoader
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader
	smLoader              *servicemappingloader.ServiceMappingLoader
	tfResourceMap         map[string]*tfschema.Resource
}

func NewManagementConflictAnnotationDefaulter(smLoader *servicemappingloader.ServiceMappingLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) *managementConflictAnnotationDefaulter {
	return &managementConflictAnnotationDefaulter{
		smLoader:              smLoader,
		serviceMetadataLoader: serviceMetadataLoader,
		dclSchemaLoader:       dclSchemaLoader,
		tfResourceMap:         google.ResourceMap(),
	}
}

// managementConflictAnnotationDefaulter implements inject.Client.
var _ inject.Client = &managementConflictAnnotationDefaulter{}

// InjectClient injects the client into the managementConflictAnnotationDefaulter
func (a *managementConflictAnnotationDefaulter) InjectClient(c client.Client) error {
	a.client = c
	return nil
}

func (a *managementConflictAnnotationDefaulter) Handle(ctx context.Context, req admission.Request) admission.Response {
	deserializer := codecs.UniversalDeserializer()
	obj := &unstructured.Unstructured{}
	if _, _, err := deserializer.Decode(req.AdmissionRequest.Object.Raw, nil, obj); err != nil {
		glog.Error(err)
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error decoding object: %v", err))
	}
	ns := &corev1.Namespace{}
	if err := a.client.Get(ctx, apimachinerytypes.NamespacedName{Name: obj.GetNamespace()}, ns); err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting Namespace %v: %v", obj.GetNamespace(), err))
	}
	if dclmetadata.IsDCLBasedResourceKind(obj.GroupVersionKind(), a.serviceMetadataLoader) {
		return defaultManagementConflictAnnotationForDCLBasedResources(obj, ns, a.dclSchemaLoader, a.serviceMetadataLoader)
	}
	return defaultManagementConflictAnnotationForTFBasedResources(obj, ns, a.smLoader, a.tfResourceMap)
}

func defaultManagementConflictAnnotationForDCLBasedResources(obj *unstructured.Unstructured, ns *corev1.Namespace, dclSchemaLoader dclschemaloader.DCLSchemaLoader, serviceMetadataLoader dclmetadata.ServiceMetadataLoader) admission.Response {
	gvk := obj.GroupVersionKind()
	stv, err := dclmetadata.ToServiceTypeVersion(gvk, serviceMetadataLoader)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting DCL ServiceTypeVersion for GroupVersionKind %v: %v", gvk, err))
	}
	schema, err := dclSchemaLoader.GetDCLSchema(stv)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError,
			fmt.Errorf("error getting the DCL Schema for GroupVersionKind %v: %v", gvk, err))
	}
	newObj := obj.DeepCopy()
	if err := k8s.ValidateOrDefaultManagementConflictPreventionAnnotationForDCLBasedResource(newObj, ns, schema); err != nil {
		return admission.Errored(http.StatusBadRequest, fmt.Errorf("error validating or defaulting management conflict policy annotation: %v", err))
	}
	return constructPatchResponse(obj, newObj)
}

func defaultManagementConflictAnnotationForTFBasedResources(obj *unstructured.Unstructured, ns *corev1.Namespace, smLoader *servicemappingloader.ServiceMappingLoader, tfResourceMap map[string]*tfschema.Resource) admission.Response {
	rc, err := smLoader.GetResourceConfig(obj)
	if err != nil {
		return admission.Errored(http.StatusBadRequest,
			fmt.Errorf("error getting ResourceConfig for kind %v: %v", obj.GetKind(), err))
	}
	newObj := obj.DeepCopy()
	if err := k8s.ValidateOrDefaultManagementConflictPreventionAnnotationForTFBasedResource(newObj, ns, rc, tfResourceMap); err != nil {
		return admission.Errored(http.StatusBadRequest, fmt.Errorf("error validating or defaulting management conflict policy annotation: %v", err))
	}
	return constructPatchResponse(obj, newObj)
}
