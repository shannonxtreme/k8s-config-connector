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

package crdgeneration

import (
	"fmt"
	"strings"

	corekccv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/slice"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

const (
	TF2CRDLabel = "cnrm.cloud.google.com/tf2crd"
)

func GenerateTF2CRD(sm *corekccv1alpha1.ServiceMapping, resourceConfig *corekccv1alpha1.ResourceConfig) (*apiextensions.CustomResourceDefinition, error) {
	resource := resourceConfig.Name
	p := google.Provider()
	r, ok := p.ResourcesMap[resource]
	if !ok {
		return nil, fmt.Errorf("unknown resource %v", resource)
	}
	s := r.Schema
	specFields := make(map[string]*schema.Schema)
	statusFields := make(map[string]*schema.Schema)
	for k, v := range s {
		if isConfigurableField(v) {
			specFields[k] = v
		} else {
			statusFields[k] = v
		}
	}
	openAPIV3Schema := crdboilerplate.GetOpenAPIV3SchemaSkeleton()
	specJSONSchema := tfObjectSchemaToJSONSchema(specFields)
	removeIgnoredAndOverwrittenFields(resourceConfig, specJSONSchema)
	markRequiredLocationalFieldsRequired(resourceConfig, specJSONSchema)
	addResourceIDFieldIfSupported(resourceConfig, specJSONSchema)
	handleHierarchicalReferences(resourceConfig, specJSONSchema)

	if len(specJSONSchema.Properties) > 0 {
		openAPIV3Schema.Properties["spec"] = *specJSONSchema
		if len(specJSONSchema.Required) > 0 {
			openAPIV3Schema.Required = slice.IncludeString(openAPIV3Schema.Required, "spec")
		}
	}
	statusSchema := tfObjectSchemaToJSONSchema(statusFields)
	for k, v := range statusSchema.Properties {
		openAPIV3Schema.Properties["status"].Properties[k] = v
	}

	group := strings.ToLower(sm.Spec.Name) + "." + ApiDomain

	kind := text.SnakeCaseToUpperCamelCase(resource)
	if resourceConfig != nil && resourceConfig.Kind != "" {
		kind = resourceConfig.Kind
	}
	crd := GetCustomResourceDefinition(kind, group, sm.Spec.Version, openAPIV3Schema, TF2CRDLabel)
	// All Terraform-based CRDs are stable
	crd.ObjectMeta.Labels[k8s.KCCStabilityLabel] = k8s.StabilityLevelStable
	return crd, nil
}

func tfObjectSchemaToJSONSchema(s map[string]*schema.Schema) *apiextensions.JSONSchemaProps {
	jsonSchema := apiextensions.JSONSchemaProps{
		Type:       "object",
		Properties: make(map[string]apiextensions.JSONSchemaProps),
	}
	for k, v := range s {
		key := text.SnakeCaseToLowerCamelCase(k)
		if v.Required {
			jsonSchema.Required = slice.IncludeString(jsonSchema.Required, key)
		}
		js := *tfSchemaToJSONSchema(v)
		description := js.Description
		if description != "" {
			description = ensureEndsInPeriod(description)
		}
		if v.ForceNew {
			description = strings.TrimSpace("Immutable. " + description)
		}
		if v.Deprecated != "" {
			deprecationMsg := ensureEndsInPeriod(fmt.Sprintf("DEPRECATED. %v", v.Deprecated))
			description = strings.TrimSpace(fmt.Sprintf("%v %v", deprecationMsg, description))
		}
		// if the description contains "terraform", ignore the description field
		for _, word := range []string{"terraform", "Terraform"} {
			if !strings.Contains(description, word) {
				continue
			}
			if v.Deprecated != "" {
				panic(fmt.Errorf("about to strip field description since it contains "+
					"the word '%v', but we likely must avoid stripping the "+
					"description entirely since it contains a deprecation message "+
					"that likely should stay included. Suggest changing field's "+
					"description and/or deprecation message to drop the word '%v'. "+
					"Description:\n%v",
					word, word, description))
			}
			description = ""
		}
		js.Description = description
		jsonSchema.Properties[key] = js
	}
	return &jsonSchema
}

func ensureEndsInPeriod(str string) string {
	if !strings.HasSuffix(str, ".") {
		return str + "."
	}
	return str
}

func tfSchemaToJSONSchema(tfSchema *schema.Schema) *apiextensions.JSONSchemaProps {
	jsonSchema := apiextensions.JSONSchemaProps{}
	switch tfSchema.Type {
	case schema.TypeBool:
		jsonSchema.Type = "boolean"
	case schema.TypeFloat:
		jsonSchema.Type = "number"
	case schema.TypeInt:
		jsonSchema.Type = "integer"
	case schema.TypeSet:
		// schema.TypeSet is just like schema.TypeList; the validation for no duplicates happens elsewhere.
		fallthrough
	case schema.TypeList:
		jsonSchema.Type = "array"
		switch v := tfSchema.Elem.(type) {
		case *schema.Resource:
			// MaxItems == 1 actually signifies that this is a nested object, and not actually a
			// list, due to limitations of the TF schema type.
			if tfSchema.MaxItems == 1 {
				jsonSchema = *tfObjectSchemaToJSONSchema(v.Schema)
				break
			}
			jsonSchema.Items = &apiextensions.JSONSchemaPropsOrArray{
				Schema: tfObjectSchemaToJSONSchema(v.Schema),
			}
		case *schema.Schema:
			// List of primitives
			jsonSchema.Items = &apiextensions.JSONSchemaPropsOrArray{
				Schema: tfSchemaToJSONSchema(v),
			}
		default:
			panic("could not parse elem attribute of TF list/set schema")
		}
	case schema.TypeMap:
		// schema.TypeMap is only used for basic map[primitive]primitive resources; maps with schemas for the keys
		// are handled by schema.TypeList with MaxItems == 1
		jsonSchema.Type = "object"
		if mapSchema, ok := tfSchema.Elem.(*schema.Schema); ok {
			jsonSchema.AdditionalProperties = &apiextensions.JSONSchemaPropsOrBool{
				Schema: tfSchemaToJSONSchema(mapSchema),
			}
		}
	case schema.TypeString:
		if tfSchema.Sensitive && isConfigurableField(tfSchema) {
			jsonSchema = crdboilerplate.GetSensitiveFieldSchemaBoilerplate()
		} else {
			jsonSchema.Type = "string"
		}
	case schema.TypeInvalid:
		panic(fmt.Errorf("schema type is invalid"))
	default:
		panic(fmt.Errorf("unknown schema type %v", tfSchema.Type))
	}
	jsonSchema.Description = tfSchema.Description
	return &jsonSchema
}

func removeIgnoredAndOverwrittenFields(rc *corekccv1alpha1.ResourceConfig, s *apiextensions.JSONSchemaProps) {
	if rc.MetadataMapping.Name != "" {
		removeField(rc.MetadataMapping.Name, s)
	}
	if rc.MetadataMapping.Labels != "" {
		removeField(rc.MetadataMapping.Labels, s)
	}
	for _, refConfig := range rc.ResourceReferences {
		handleResourceReference(refConfig, s)
	}
	for _, d := range rc.Directives {
		removeField(d, s)
	}
	for _, f := range rc.IgnoredFields {
		removeField(f, s)
	}
	if !krmtotf.SupportsHierarchicalReferences(rc) {
		// TODO(b/193177782): Delete this if-block once all resources support
		// hierarchical references.
		for _, c := range rc.Containers {
			removeField(c.TFField, s)
		}
	}
}

func markRequiredLocationalFieldsRequired(rc *corekccv1alpha1.ResourceConfig, s *apiextensions.JSONSchemaProps) {
	if rc.IDTemplate == "" {
		return
	}

	locationalFields := []string{"region", "zone", "location"}
	for _, field := range locationalFields {
		// It is assumed that locational fields (region, zone, location) would
		// always be at the base level.
		if _, ok := s.Properties[field]; !ok {
			continue
		}
		if !strings.Contains(rc.IDTemplate, fmt.Sprintf("{{%v}}", field)) {
			continue
		}
		s.Required = slice.IncludeString(s.Required, field)
	}
}

func handleResourceReference(refConfig corekccv1alpha1.ReferenceConfig, s *apiextensions.JSONSchemaProps) {
	*s = populateReference(strings.Split(refConfig.TFField, "."), refConfig, s)
}

func populateReference(path []string, refConfig corekccv1alpha1.ReferenceConfig, s *apiextensions.JSONSchemaProps) apiextensions.JSONSchemaProps {
	field := text.SnakeCaseToLowerCamelCase(path[0])
	if len(path) > 1 {
		subSchema := s.Properties[field]
		switch subSchema.Type {
		case "array":
			itemSchema := populateReference(path[1:], refConfig, subSchema.Items.Schema)
			subSchema.Items.Schema = &itemSchema
			return *s
		case "object":
			objSchema := populateReference(path[1:], refConfig, &subSchema)
			s.Properties[field] = objSchema
			return *s
		default:
			panic(fmt.Errorf("error parsing reference %v: cannot iterate into type that is not object or array of objects", path))
		}
	}

	// Base case; we have found the field representing the reference
	isList := s.Properties[field].Type == "array"
	var refSchema *apiextensions.JSONSchemaProps
	key := field
	if len(refConfig.Types) == 0 {
		if refConfig.Key != "" {
			key = refConfig.Key
			delete(s.Properties, field)
			if slice.StringSliceContains(s.Required, field) {
				s.Required = slice.RemoveStringFromStringSlice(s.Required, field)
				s.Required = slice.IncludeString(s.Required, key)
			}
		}
		refSchema = GetResourceReferenceSchemaFromTypeConfig(refConfig.TypeConfig)
	} else {
		refSchema = &apiextensions.JSONSchemaProps{
			Type:       "object",
			Properties: map[string]apiextensions.JSONSchemaProps{},
		}
		for _, v := range refConfig.Types {
			if v.JSONSchemaType == "" {
				refSchema.Properties[v.Key] = *GetResourceReferenceSchemaFromTypeConfig(v)
			} else {
				refSchema.Properties[v.Key] = apiextensions.JSONSchemaProps{
					Type: v.JSONSchemaType,
				}
			}
		}
	}

	refSchema.Description = refConfig.Description

	if isList {
		s.Properties[key] = apiextensions.JSONSchemaProps{
			Type: "array",
			Items: &apiextensions.JSONSchemaPropsOrArray{
				Schema: refSchema,
			},
		}
	} else {
		s.Properties[key] = *refSchema
	}

	return *s
}

func getDescriptionForExternalRef(typeConfig corekccv1alpha1.TypeConfig) string {
	targetField := typeConfig.TargetField
	if targetField == "" {
		targetField = "name"
	}
	targetField = text.SnakeCaseToLowerCamelCase(targetField)
	article := text.IndefiniteArticleFor(typeConfig.GVK.Kind)
	if typeConfig.ValueTemplate != "" {
		return fmt.Sprintf(
			"Allowed value: string of the format `%v`, where {{value}} is the `%v` field of %v `%v` resource.",
			typeConfig.ValueTemplate, targetField, article, typeConfig.GVK.Kind,
		)
	}
	return fmt.Sprintf("Allowed value: The `%v` field of %v `%v` resource.", targetField, article, typeConfig.GVK.Kind)
}

func GetResourceReferenceSchemaFromTypeConfig(typeConfig corekccv1alpha1.TypeConfig) *apiextensions.JSONSchemaProps {
	description := getDescriptionForExternalRef(typeConfig)
	return crdboilerplate.GetResourceReferenceSchemaBoilerplate(description)
}

func removeField(tfField string, s *apiextensions.JSONSchemaProps) {
	*s = removeNestedField(strings.Split(tfField, "."), *s)
}

func removeNestedField(path []string, s apiextensions.JSONSchemaProps) apiextensions.JSONSchemaProps {
	field := text.SnakeCaseToLowerCamelCase(path[0])
	if len(path) > 1 {
		subSchema := s.Properties[field]
		switch subSchema.Type {
		case "array":
			itemSchema := removeNestedField(path[1:], *subSchema.Items.Schema)
			subSchema.Items.Schema = &itemSchema
		case "object":
			subSchema = removeNestedField(path[1:], subSchema)
		default:
			panic(fmt.Errorf("error parsing field %v: cannot iterate into type that is not object or array of objects", path))
		}
		s.Properties[field] = subSchema
		return s
	}
	delete(s.Properties, field)
	s.Required = slice.RemoveStringFromStringSlice(s.Required, field)
	return s
}

func isConfigurableField(tfSchema *schema.Schema) bool {
	return tfSchema.Required || tfSchema.Optional
}

func addResourceIDFieldIfSupported(rc *corekccv1alpha1.ResourceConfig, spec *apiextensions.JSONSchemaProps) {
	if !krmtotf.SupportsResourceIDField(rc) {
		return
	}

	spec.Properties[k8s.ResourceIDFieldName] = apiextensions.JSONSchemaProps{
		Type:        "string",
		Description: generateResourceIDFieldDescription(rc),
	}
}

func generateResourceIDFieldDescription(rc *corekccv1alpha1.ResourceConfig) string {
	targetFieldCamelCase := text.SnakeCaseToLowerCamelCase(rc.ResourceID.TargetField)
	isServerGeneratedResourceID := krmtotf.IsResourceIDFieldServerGenerated(rc)
	return GenerateResourceIDFieldDescription(targetFieldCamelCase, isServerGeneratedResourceID)
}

func handleHierarchicalReferences(rc *corekccv1alpha1.ResourceConfig, spec *apiextensions.JSONSchemaProps) {
	if len(rc.Containers) > 0 {
		// If resource supports resource-level container annotations, mark
		// hierarchical references optional since users can use the annotations
		// to configure the references.
		*spec = *MarkHierarchicalReferencesOptionalButMutuallyExclusive(spec, rc.HierarchicalReferences)
	} else {
		*spec = *MarkHierarchicalReferencesRequiredButMutuallyExclusive(spec, rc.HierarchicalReferences)
	}
}
