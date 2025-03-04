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

type BucketAction struct {
	/* The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. */
	// +optional
	StorageClass *string `json:"storageClass,omitempty"`

	/* The type of the action of this Lifecycle Rule. Supported values include: Delete and SetStorageClass. */
	Type string `json:"type"`
}

type BucketCondition struct {
	/* Minimum age of an object in days to satisfy this condition. */
	// +optional
	Age *int `json:"age,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	// +optional
	CreatedBefore *string `json:"createdBefore,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	// +optional
	CustomTimeBefore *string `json:"customTimeBefore,omitempty"`

	/* Number of days elapsed since the user-specified timestamp set on an object. */
	// +optional
	DaysSinceCustomTime *int `json:"daysSinceCustomTime,omitempty"`

	/* Number of days elapsed since the noncurrent timestamp of an object. This
	condition is relevant only for versioned objects. */
	// +optional
	DaysSinceNoncurrentTime *int `json:"daysSinceNoncurrentTime,omitempty"`

	/* One or more matching name prefixes to satisfy this condition. */
	// +optional
	MatchesPrefix []string `json:"matchesPrefix,omitempty"`

	/* Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY. */
	// +optional
	MatchesStorageClass []string `json:"matchesStorageClass,omitempty"`

	/* One or more matching name suffixes to satisfy this condition. */
	// +optional
	MatchesSuffix []string `json:"matchesSuffix,omitempty"`

	/* Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition. */
	// +optional
	NoncurrentTimeBefore *string `json:"noncurrentTimeBefore,omitempty"`

	/* Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition. */
	// +optional
	NumNewerVersions *int `json:"numNewerVersions,omitempty"`

	/* Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY". */
	// +optional
	WithState *string `json:"withState,omitempty"`
}

type BucketCors struct {
	/* The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses. */
	// +optional
	MaxAgeSeconds *int `json:"maxAgeSeconds,omitempty"`

	/* The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "*" is permitted in the list of methods, and means "any method". */
	// +optional
	Method []string `json:"method,omitempty"`

	/* The list of Origins eligible to receive CORS response headers. Note: "*" is permitted in the list of origins, and means "any Origin". */
	// +optional
	Origin []string `json:"origin,omitempty"`

	/* The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains. */
	// +optional
	ResponseHeader []string `json:"responseHeader,omitempty"`
}

type BucketEncryption struct {
	/*  */
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef"`
}

type BucketLifecycleRule struct {
	/* The Lifecycle Rule's action configuration. A single block of this type is supported. */
	Action BucketAction `json:"action"`

	/* The Lifecycle Rule's condition configuration. */
	Condition BucketCondition `json:"condition"`
}

type BucketLogging struct {
	/* The bucket that will receive log objects. */
	LogBucket string `json:"logBucket"`

	/* The object prefix for log objects. If it's not provided, by default Google Cloud Storage sets this to this bucket's name. */
	// +optional
	LogObjectPrefix *string `json:"logObjectPrefix,omitempty"`
}

type BucketRetentionPolicy struct {
	/* If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action. */
	// +optional
	IsLocked *bool `json:"isLocked,omitempty"`

	/* The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds. */
	RetentionPeriod int `json:"retentionPeriod"`
}

type BucketVersioning struct {
	/* While set to true, versioning is fully enabled for this bucket. */
	Enabled bool `json:"enabled"`
}

type BucketWebsite struct {
	/* Behaves as the bucket's directory index where missing objects are treated as potential directories. */
	// +optional
	MainPageSuffix *string `json:"mainPageSuffix,omitempty"`

	/* The custom object to return when a requested resource is not found. */
	// +optional
	NotFoundPage *string `json:"notFoundPage,omitempty"`
}

type StorageBucketSpec struct {
	/* DEPRECATED. Please use the `uniformBucketLevelAccess` field as this field has been renamed by Google. The `uniformBucketLevelAccess` field will supersede this field.
	Enables Bucket PolicyOnly access to a bucket. */
	// +optional
	BucketPolicyOnly *bool `json:"bucketPolicyOnly,omitempty"`

	/* The bucket's Cross-Origin Resource Sharing (CORS) configuration. */
	// +optional
	Cors []BucketCors `json:"cors,omitempty"`

	/* Whether or not to automatically apply an eventBasedHold to new objects added to the bucket. */
	// +optional
	DefaultEventBasedHold *bool `json:"defaultEventBasedHold,omitempty"`

	/* The bucket's encryption configuration. */
	// +optional
	Encryption *BucketEncryption `json:"encryption,omitempty"`

	/* The bucket's Lifecycle Rules configuration. */
	// +optional
	LifecycleRule []BucketLifecycleRule `json:"lifecycleRule,omitempty"`

	/* Immutable. The Google Cloud Storage location. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* The bucket's Access & Storage Logs configuration. */
	// +optional
	Logging *BucketLogging `json:"logging,omitempty"`

	/* Prevents public access to a bucket. */
	// +optional
	PublicAccessPrevention *string `json:"publicAccessPrevention,omitempty"`

	/* Enables Requester Pays on a storage bucket. */
	// +optional
	RequesterPays *bool `json:"requesterPays,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. */
	// +optional
	RetentionPolicy *BucketRetentionPolicy `json:"retentionPolicy,omitempty"`

	/* The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE. */
	// +optional
	StorageClass *string `json:"storageClass,omitempty"`

	/* Enables uniform bucket-level access on a bucket. */
	// +optional
	UniformBucketLevelAccess *bool `json:"uniformBucketLevelAccess,omitempty"`

	/* The bucket's Versioning configuration. */
	// +optional
	Versioning *BucketVersioning `json:"versioning,omitempty"`

	/* Configuration if the bucket acts as a website. */
	// +optional
	Website *BucketWebsite `json:"website,omitempty"`
}

type StorageBucketStatus struct {
	/* Conditions represent the latest available observations of the
	   StorageBucket's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* The URI of the created resource. */
	SelfLink string `json:"selfLink,omitempty"`
	/* The base URL of the bucket, in the format gs://<bucket-name>. */
	Url string `json:"url,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageBucket is the Schema for the storage API
// +k8s:openapi-gen=true
type StorageBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageBucketSpec   `json:"spec,omitempty"`
	Status StorageBucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageBucketList contains a list of StorageBucket
type StorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageBucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageBucket{}, &StorageBucketList{})
}
