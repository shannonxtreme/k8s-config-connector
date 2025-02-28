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

package k8s

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/types"
)

type ManagementConflictPreventionPolicy string

// TODO: clean up old conditions used in handcrafted controllers
const (
	CNRMGroup                            = "cnrm.cloud.google.com"
	ApiDomainSuffix                      = ".cnrm.cloud.google.com"
	SystemNamespace                      = "cnrm-system"
	ControllerMaxConcurrentReconciles    = 20
	ReconcileDeadline                    = 59 * time.Minute
	TimeToLeaseExpiration                = 40 * time.Minute
	TimeToLeaseRenewal                   = 20 * time.Minute
	MeanReconcileReenqueuePeriod         = 10 * time.Minute
	JitterFactor                         = 2.0
	UpToDate                             = "UpToDate"
	UpToDateMessage                      = "The resource is up to date"
	Created                              = "Created"
	CreatedMessage                       = "Successfully created"
	CreateFailed                         = "CreateFailed"
	CreateFailedMessageTmpl              = "Create call failed: %v"
	Updating                             = "Updating"
	UpdatingMessage                      = "Update in progress"
	UpdateFailed                         = "UpdateFailed"
	UpdateFailedMessageTmpl              = "Update call failed: %v"
	Deleting                             = "Deleting"
	DeletingMessage                      = "Deletion in progress"
	Deleted                              = "Deleted"
	DeletedMessage                       = "Successfully deleted"
	DeleteFailed                         = "DeleteFailed"
	NoCondition                          = "NoCondition"
	DeleteFailedMessageTmpl              = "Delete call failed: %v"
	ControllerFinalizerName              = "cnrm.cloud.google.com/finalizer"
	DeletionDefenderFinalizerName        = "cnrm.cloud.google.com/deletion-defender"
	DependencyNotReady                   = "DependencyNotReady"
	DependencyNotFound                   = "DependencyNotFound"
	DependencyInvalid                    = "DependencyInvalid"
	ManagementConflict                   = "ManagementConflict"
	PreActuationTransformFailed          = "PreActuationTransformFailed"
	PostActuationTransformFailed         = "PostActuationTransformFailed"
	DeletionPolicyDelete                 = "delete"
	DeletionPolicyAbandon                = "abandon"
	AnnotationPrefix                     = CNRMGroup
	NamespaceEnvVar                      = "NAMESPACE"
	ImmediateReconcileRequestsBufferSize = 10000
	MaxNumResourceWatcherRoutines        = 10000

	ReadinessServerPort = 23232
	ReadinessServerPath = "/ready"

	ControllerManagedFieldManager = "cnrm-controller-manager"
	SupportsSSAManager            = "supports-ssa"

	// Management conflict prevention policies
	ManagementConflictPreventionPolicyNone     = "none"
	ManagementConflictPreventionPolicyResource = "resource"

	// State into spec annotation values
	StateMergeIntoSpec = "merge"
	StateAbsentInSpec  = "absent"

	// Core kubernetes constants
	LastAppliedConfigurationAnnotation = "kubectl.kubernetes.io/last-applied-configuration"
	ManagedFieldsTypeFieldsV1          = "FieldsV1"

	ResourceIDFieldName = "resourceID"
	ResourceIDFieldPath = "spec." + ResourceIDFieldName

	StabilityLevelStable = "stable"
	StabilityLevelAlpha  = "alpha"

	KCCAPIVersion = "v1beta1"
)

var (
	DeletionPolicyAnnotation = FormatAnnotation("deletion-policy")

	// Annotations for Container objects
	ProjectIDAnnotation  = FormatAnnotation("project-id")
	FolderIDAnnotation   = FormatAnnotation("folder-id")
	OrgIDAnnotation      = FormatAnnotation("organization-id")
	ContainerAnnotations = []string{
		ProjectIDAnnotation,
		FolderIDAnnotation,
		OrgIDAnnotation,
	}

	ManagementConflictPreventionPolicyAnnotation               = "management-conflict-prevention-policy"
	ManagementConflictPreventionPolicyFullyQualifiedAnnotation = FormatAnnotation(ManagementConflictPreventionPolicyAnnotation)
	ManagementConflictPreventionPolicyValues                   = []string{
		ManagementConflictPreventionPolicyNone,
		ManagementConflictPreventionPolicyResource,
	}

	KCCComponentLabel = FormatAnnotation("component")
	KCCSystemLabel    = FormatAnnotation("system")
	KCCVersionLabel   = FormatAnnotation("version")
	DCL2CRDLabel      = FormatAnnotation("dcl2crd")
	KCCStabilityLabel = FormatAnnotation("stability-level")

	MutableButUnreadableFieldsAnnotation = FormatAnnotation("mutable-but-unreadable-fields")
	ObservedSecretVersionsAnnotation     = FormatAnnotation("observed-secret-versions")

	SupportsSSAAnnotation = FormatAnnotation("supports-ssa")

	BlueprintAttributionAnnotation = FormatAnnotation("blueprint")

	StateIntoSpecAnnotation       = FormatAnnotation("state-into-spec")
	StateIntoSpecAnnotationValues = []string{
		StateMergeIntoSpec,
		StateAbsentInSpec,
	}
	// TODO(kcc-eng): Adjust the timeout back down after b/237398742 is fixed.
	WebhookTimeoutSeconds = int32(10)

	ReservedStatusFieldNamesForFutureUse = []string{"generation"}

	NamespaceIDConfigMapNN = types.NamespacedName{
		Namespace: SystemNamespace,
		Name:      "namespace-id",
	}
)

func FormatAnnotation(annotationName string) string {
	return fmt.Sprintf("%v/%v", AnnotationPrefix, annotationName)
}
