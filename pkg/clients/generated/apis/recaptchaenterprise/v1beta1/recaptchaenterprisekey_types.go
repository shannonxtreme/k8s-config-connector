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

type KeyAndroidSettings struct {
	/* If set to true, it means allowed_package_names will not be enforced. */
	// +optional
	AllowAllPackageNames *bool `json:"allowAllPackageNames,omitempty"`

	/* Android package names of apps allowed to use the key. Example: 'com.companyname.appname' */
	// +optional
	AllowedPackageNames []string `json:"allowedPackageNames,omitempty"`
}

type KeyIosSettings struct {
	/* If set to true, it means allowed_bundle_ids will not be enforced. */
	// +optional
	AllowAllBundleIds *bool `json:"allowAllBundleIds,omitempty"`

	/* iOS bundle ids of apps allowed to use the key. Example: 'com.companyname.productname.appname' */
	// +optional
	AllowedBundleIds []string `json:"allowedBundleIds,omitempty"`
}

type KeyTestingOptions struct {
	/* Immutable. For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if UNSOLVABLE_CHALLENGE. Possible values: TESTING_CHALLENGE_UNSPECIFIED, NOCAPTCHA, UNSOLVABLE_CHALLENGE */
	// +optional
	TestingChallenge *string `json:"testingChallenge,omitempty"`

	/* Immutable. All assessments for this Key will return this score. Must be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive. */
	// +optional
	TestingScore *float64 `json:"testingScore,omitempty"`
}

type KeyWebSettings struct {
	/* If set to true, it means allowed_domains will not be enforced. */
	// +optional
	AllowAllDomains *bool `json:"allowAllDomains,omitempty"`

	/* If set to true, the key can be used on AMP (Accelerated Mobile Pages) websites. This is supported only for the SCORE integration type. */
	// +optional
	AllowAmpTraffic *bool `json:"allowAmpTraffic,omitempty"`

	/* Domains or subdomains of websites allowed to use the key. All subdomains of an allowed domain are automatically allowed. A valid domain requires a host and must not include any path, port, query or fragment. Examples: 'example.com' or 'subdomain.example.com' */
	// +optional
	AllowedDomains []string `json:"allowedDomains,omitempty"`

	/* Settings for the frequency and difficulty at which this key triggers captcha challenges. This should only be specified for IntegrationTypes CHECKBOX and INVISIBLE. Possible values: CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED, USABILITY, BALANCE, SECURITY */
	// +optional
	ChallengeSecurityPreference *string `json:"challengeSecurityPreference,omitempty"`

	/* Immutable. Required. Describes how this key is integrated with the website. Possible values: SCORE, CHECKBOX, INVISIBLE */
	IntegrationType string `json:"integrationType"`
}

type RecaptchaEnterpriseKeySpec struct {
	/* Settings for keys that can be used by Android apps. */
	// +optional
	AndroidSettings *KeyAndroidSettings `json:"androidSettings,omitempty"`

	/* Human-readable display name of this key. Modifiable by user. */
	DisplayName string `json:"displayName"`

	/* Settings for keys that can be used by iOS apps. */
	// +optional
	IosSettings *KeyIosSettings `json:"iosSettings,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Options for user acceptance testing. */
	// +optional
	TestingOptions *KeyTestingOptions `json:"testingOptions,omitempty"`

	/* Settings for keys that can be used by websites. */
	// +optional
	WebSettings *KeyWebSettings `json:"webSettings,omitempty"`
}

type RecaptchaEnterpriseKeyStatus struct {
	/* Conditions represent the latest available observations of the
	   RecaptchaEnterpriseKey's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The timestamp corresponding to the creation of this Key. */
	CreateTime string `json:"createTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RecaptchaEnterpriseKey is the Schema for the recaptchaenterprise API
// +k8s:openapi-gen=true
type RecaptchaEnterpriseKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RecaptchaEnterpriseKeySpec   `json:"spec,omitempty"`
	Status RecaptchaEnterpriseKeyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RecaptchaEnterpriseKeyList contains a list of RecaptchaEnterpriseKey
type RecaptchaEnterpriseKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecaptchaEnterpriseKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RecaptchaEnterpriseKey{}, &RecaptchaEnterpriseKeyList{})
}
