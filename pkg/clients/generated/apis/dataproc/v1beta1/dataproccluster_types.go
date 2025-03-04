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

type ClusterAccelerators struct {
	/* Immutable. The number of the accelerator cards of this type exposed to this instance. */
	// +optional
	AcceleratorCount *int `json:"acceleratorCount,omitempty"`

	/* Immutable. Full URL, partial URI, or short name of the accelerator type resource to expose to this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes). Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80` * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the accelerator type resource, for example, `nvidia-tesla-k80`. */
	// +optional
	AcceleratorType *string `json:"acceleratorType,omitempty"`
}

type ClusterAutoscalingConfig struct {
	/* Immutable. */
	// +optional
	PolicyRef *v1alpha1.ResourceRef `json:"policyRef,omitempty"`
}

type ClusterConfig struct {
	/* Immutable. Optional. Autoscaling config for the policy associated with the cluster. Cluster does not autoscale if this field is unset. */
	// +optional
	AutoscalingConfig *ClusterAutoscalingConfig `json:"autoscalingConfig,omitempty"`

	/* Immutable. Optional. Encryption settings for the cluster. */
	// +optional
	EncryptionConfig *ClusterEncryptionConfig `json:"encryptionConfig,omitempty"`

	/* Immutable. Optional. Port/endpoint configuration for this cluster */
	// +optional
	EndpointConfig *ClusterEndpointConfig `json:"endpointConfig,omitempty"`

	/* Immutable. Optional. The shared Compute Engine config settings for all instances in a cluster. */
	// +optional
	GceClusterConfig *ClusterGceClusterConfig `json:"gceClusterConfig,omitempty"`

	/* Immutable. Optional. Commands to execute on each node after config is completed. By default, executables are run on master and all worker nodes. You can test a node's `role` metadata to run an executable on a master or worker node, as shown below using `curl` (you can also use `wget`): ROLE=$(curl -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role) if [[ "${ROLE}" == 'Master' ]]; then ... master specific actions ... else ... worker specific actions ... fi */
	// +optional
	InitializationActions []ClusterInitializationActions `json:"initializationActions,omitempty"`

	/* Immutable. Optional. Lifecycle setting for the cluster. */
	// +optional
	LifecycleConfig *ClusterLifecycleConfig `json:"lifecycleConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for the master instance in a cluster. */
	// +optional
	MasterConfig *ClusterMasterConfig `json:"masterConfig,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for additional worker instances in a cluster. */
	// +optional
	SecondaryWorkerConfig *ClusterSecondaryWorkerConfig `json:"secondaryWorkerConfig,omitempty"`

	/* Immutable. Optional. Security settings for the cluster. */
	// +optional
	SecurityConfig *ClusterSecurityConfig `json:"securityConfig,omitempty"`

	/* Immutable. Optional. The config settings for software inside the cluster. */
	// +optional
	SoftwareConfig *ClusterSoftwareConfig `json:"softwareConfig,omitempty"`

	/* Immutable. */
	// +optional
	StagingBucketRef *v1alpha1.ResourceRef `json:"stagingBucketRef,omitempty"`

	/* Immutable. */
	// +optional
	TempBucketRef *v1alpha1.ResourceRef `json:"tempBucketRef,omitempty"`

	/* Immutable. Optional. The Compute Engine config settings for worker instances in a cluster. */
	// +optional
	WorkerConfig *ClusterWorkerConfig `json:"workerConfig,omitempty"`
}

type ClusterDiskConfig struct {
	/* Immutable. Optional. Size in GB of the boot disk (default is 500GB). */
	// +optional
	BootDiskSizeGb *int `json:"bootDiskSizeGb,omitempty"`

	/* Immutable. Optional. Type of the boot disk (default is "pd-standard"). Valid values: "pd-balanced" (Persistent Disk Balanced Solid State Drive), "pd-ssd" (Persistent Disk Solid State Drive), or "pd-standard" (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types). */
	// +optional
	BootDiskType *string `json:"bootDiskType,omitempty"`

	/* Immutable. Optional. Number of attached SSDs, from 0 to 4 (default is 0). If SSDs are not attached, the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html) data. If one or more SSDs are attached, this runtime bulk data is spread across them, and the boot disk contains only basic config and installed binaries. */
	// +optional
	NumLocalSsds *int `json:"numLocalSsds,omitempty"`
}

type ClusterEncryptionConfig struct {
	/* Immutable. */
	// +optional
	GcePdKmsKeyRef *v1alpha1.ResourceRef `json:"gcePdKmsKeyRef,omitempty"`
}

type ClusterEndpointConfig struct {
	/* Immutable. Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false. */
	// +optional
	EnableHttpPortAccess *bool `json:"enableHttpPortAccess,omitempty"`
}

type ClusterGceClusterConfig struct {
	/* Immutable. Optional. If true, all instances in the cluster will only have internal IP addresses. By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. This `internal_ip_only` restriction can only be enabled for subnetwork enabled networks, and all off-cluster dependencies must be configured to be accessible without external IP addresses. */
	// +optional
	InternalIPOnly *bool `json:"internalIPOnly,omitempty"`

	/* Immutable. The Compute Engine metadata entries to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)). */
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`

	/* Immutable. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. Optional. Node Group Affinity for sole-tenant clusters. */
	// +optional
	NodeGroupAffinity *ClusterNodeGroupAffinity `json:"nodeGroupAffinity,omitempty"`

	/* Immutable. Optional. The type of IPv6 access for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNETWORK, OUTBOUND, BIDIRECTIONAL */
	// +optional
	PrivateIPv6GoogleAccess *string `json:"privateIPv6GoogleAccess,omitempty"`

	/* Immutable. Optional. Reservation Affinity for consuming Zonal reservation. */
	// +optional
	ReservationAffinity *ClusterReservationAffinity `json:"reservationAffinity,omitempty"`

	/* Immutable. */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* Immutable. Optional. The URIs of service account scopes to be included in Compute Engine instances. The following base set of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly * https://www.googleapis.com/auth/devstorage.read_write * https://www.googleapis.com/auth/logging.write If no scopes are specified, the following defaults are also provided: * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control */
	// +optional
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	/* Immutable. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`

	/* Immutable. The Compute Engine tags to add to all instances (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)). */
	// +optional
	Tags []string `json:"tags,omitempty"`

	/* Immutable. Optional. The zone where the Compute Engine cluster will be located. On a create request, it is required in the "global" region. If omitted in a non-global Dataproc region, the service will pick a zone in the corresponding Compute Engine region. On a get request, zone will always be present. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]` * `projects/[project_id]/zones/[zone]` * `us-central1-f` */
	// +optional
	Zone *string `json:"zone,omitempty"`
}

type ClusterInitializationActions struct {
	/* Immutable. Required. Cloud Storage URI of executable file. */
	ExecutableFile string `json:"executableFile"`

	/* Immutable. Optional. Amount of time executable has to complete. Default is 10 minutes (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period. */
	// +optional
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`
}

type ClusterKerberosConfig struct {
	/* Immutable. Optional. The admin server (IP or hostname) for the remote trusted realm in a cross realm trust relationship. */
	// +optional
	CrossRealmTrustAdminServer *string `json:"crossRealmTrustAdminServer,omitempty"`

	/* Immutable. Optional. The KDC (IP or hostname) for the remote trusted realm in a cross realm trust relationship. */
	// +optional
	CrossRealmTrustKdc *string `json:"crossRealmTrustKdc,omitempty"`

	/* Immutable. Optional. The remote realm the Dataproc on-cluster KDC will trust, should the user enable cross realm trust. */
	// +optional
	CrossRealmTrustRealm *string `json:"crossRealmTrustRealm,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the shared password between the on-cluster Kerberos realm and the remote trusted realm, in a cross realm trust relationship. */
	// +optional
	CrossRealmTrustSharedPassword *string `json:"crossRealmTrustSharedPassword,omitempty"`

	/* Immutable. Optional. Flag to indicate whether to Kerberize the cluster (default: false). Set this field to true to enable Kerberos on a cluster. */
	// +optional
	EnableKerberos *bool `json:"enableKerberos,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the master key of the KDC database. */
	// +optional
	KdcDbKey *string `json:"kdcDbKey,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided key. For the self-signed certificate, this password is generated by Dataproc. */
	// +optional
	KeyPassword *string `json:"keyPassword,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of the keystore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate. */
	// +optional
	Keystore *string `json:"keystore,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided keystore. For the self-signed certificate, this password is generated by Dataproc. */
	// +optional
	KeystorePassword *string `json:"keystorePassword,omitempty"`

	/* Immutable. */
	// +optional
	KmsKeyRef *v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`

	/* Immutable. Optional. The name of the on-cluster Kerberos realm. If not specified, the uppercased domain of hostnames will be the realm. */
	// +optional
	Realm *string `json:"realm,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the root principal password. */
	// +optional
	RootPrincipalPassword *string `json:"rootPrincipalPassword,omitempty"`

	/* Immutable. Optional. The lifetime of the ticket granting ticket, in hours. If not specified, or user specifies 0, then default value 10 will be used. */
	// +optional
	TgtLifetimeHours *int `json:"tgtLifetimeHours,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of the truststore file used for SSL encryption. If not provided, Dataproc will provide a self-signed certificate. */
	// +optional
	Truststore *string `json:"truststore,omitempty"`

	/* Immutable. Optional. The Cloud Storage URI of a KMS encrypted file containing the password to the user provided truststore. For the self-signed certificate, this password is generated by Dataproc. */
	// +optional
	TruststorePassword *string `json:"truststorePassword,omitempty"`
}

type ClusterLifecycleConfig struct {
	/* Immutable. Optional. The time when cluster will be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +optional
	AutoDeleteTime *string `json:"autoDeleteTime,omitempty"`

	/* Immutable. Optional. The lifetime duration of cluster. The cluster will be auto-deleted at the end of this period. Minimum value is 10 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +optional
	AutoDeleteTtl *string `json:"autoDeleteTtl,omitempty"`

	/* Immutable. Optional. The duration to keep the cluster alive while idling (when no jobs are running). Passing this threshold will cause the cluster to be deleted. Minimum value is 5 minutes; maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	// +optional
	IdleDeleteTtl *string `json:"idleDeleteTtl,omitempty"`
}

type ClusterMasterConfig struct {
	/* Immutable. Optional. The Compute Engine accelerator configuration for these instances. */
	// +optional
	Accelerators []ClusterAccelerators `json:"accelerators,omitempty"`

	/* Immutable. Optional. Disk option config settings. */
	// +optional
	DiskConfig *ClusterDiskConfig `json:"diskConfig,omitempty"`

	/* Immutable. */
	// +optional
	ImageRef *v1alpha1.ResourceRef `json:"imageRef,omitempty"`

	/* Immutable. Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu). */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**. */
	// +optional
	NumInstances *int `json:"numInstances,omitempty"`

	/* Immutable. Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE */
	// +optional
	Preemptibility *string `json:"preemptibility,omitempty"`
}

type ClusterNodeGroupAffinity struct {
	/* Immutable. */
	NodeGroupRef v1alpha1.ResourceRef `json:"nodeGroupRef"`
}

type ClusterReservationAffinity struct {
	/* Immutable. Optional. Type of reservation to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION, SPECIFIC_RESERVATION */
	// +optional
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	/* Immutable. Optional. Corresponds to the label key of reservation resource. */
	// +optional
	Key *string `json:"key,omitempty"`

	/* Immutable. Optional. Corresponds to the label values of reservation resource. */
	// +optional
	Values []string `json:"values,omitempty"`
}

type ClusterSecondaryWorkerConfig struct {
	/* Immutable. Optional. The Compute Engine accelerator configuration for these instances. */
	// +optional
	Accelerators []ClusterAccelerators `json:"accelerators,omitempty"`

	/* Immutable. Optional. Disk option config settings. */
	// +optional
	DiskConfig *ClusterDiskConfig `json:"diskConfig,omitempty"`

	/* Immutable. */
	// +optional
	ImageRef *v1alpha1.ResourceRef `json:"imageRef,omitempty"`

	/* Immutable. Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu). */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**. */
	// +optional
	NumInstances *int `json:"numInstances,omitempty"`

	/* Immutable. Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE */
	// +optional
	Preemptibility *string `json:"preemptibility,omitempty"`
}

type ClusterSecurityConfig struct {
	/* Immutable. Optional. Kerberos related configuration. */
	// +optional
	KerberosConfig *ClusterKerberosConfig `json:"kerberosConfig,omitempty"`
}

type ClusterSoftwareConfig struct {
	/* Immutable. Optional. The version of software inside the cluster. It must be one of the supported [Dataproc Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported_dataproc_versions), such as "1.2" (including a subminor version, such as "1.2.29"), or the ["preview" version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions). If unspecified, it defaults to the latest Debian version. */
	// +optional
	ImageVersion *string `json:"imageVersion,omitempty"`

	/* Immutable. Optional. The set of components to activate on the cluster. */
	// +optional
	OptionalComponents []string `json:"optionalComponents,omitempty"`

	/* Immutable. Optional. The properties to set on daemon config files. Property keys are specified in `prefix:property` format, for example `core:hadoop.tmp.dir`. The following are supported prefixes and their mappings: * capacity-scheduler: `capacity-scheduler.xml` * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml` * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties` * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties). */
	// +optional
	Properties map[string]string `json:"properties,omitempty"`
}

type ClusterWorkerConfig struct {
	/* Immutable. Optional. The Compute Engine accelerator configuration for these instances. */
	// +optional
	Accelerators []ClusterAccelerators `json:"accelerators,omitempty"`

	/* Immutable. Optional. Disk option config settings. */
	// +optional
	DiskConfig *ClusterDiskConfig `json:"diskConfig,omitempty"`

	/* Immutable. */
	// +optional
	ImageRef *v1alpha1.ResourceRef `json:"imageRef,omitempty"`

	/* Immutable. Optional. The Compute Engine machine type used for cluster instances. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2` * `n1-standard-2` **Auto Zone Exception**: If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement) feature, you must use the short name of the machine type resource, for example, `n1-standard-2`. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Immutable. Optional. Specifies the minimum cpu platform for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu). */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* Immutable. Optional. The number of VM instances in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability) [master_config](#FIELDS.master_config) groups, **must be set to 3**. For standard cluster [master_config](#FIELDS.master_config) groups, **must be set to 1**. */
	// +optional
	NumInstances *int `json:"numInstances,omitempty"`

	/* Immutable. Optional. Specifies the preemptibility of the instance group. The default value for master and worker groups is `NON_PREEMPTIBLE`. This default cannot be changed. The default value for secondary instances is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE */
	// +optional
	Preemptibility *string `json:"preemptibility,omitempty"`
}

type DataprocClusterSpec struct {
	/* Immutable. Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated. */
	// +optional
	Config *ClusterConfig `json:"config,omitempty"`

	/* Immutable. The location for the resource, usually a GCP region. */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	// +optional
	ProjectRef *v1alpha1.ResourceRef `json:"projectRef,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ClusterConfigStatus struct {
	/*  */
	EndpointConfig ClusterEndpointConfigStatus `json:"endpointConfig,omitempty"`

	/*  */
	LifecycleConfig ClusterLifecycleConfigStatus `json:"lifecycleConfig,omitempty"`

	/*  */
	MasterConfig ClusterMasterConfigStatus `json:"masterConfig,omitempty"`

	/*  */
	SecondaryWorkerConfig ClusterSecondaryWorkerConfigStatus `json:"secondaryWorkerConfig,omitempty"`

	/*  */
	WorkerConfig ClusterWorkerConfigStatus `json:"workerConfig,omitempty"`
}

type ClusterEndpointConfigStatus struct {
	/* Output only. The map of port descriptions to URLs. Will only be populated if enable_http_port_access is true. */
	HttpPorts map[string]string `json:"httpPorts,omitempty"`
}

type ClusterLifecycleConfigStatus struct {
	/* Output only. The time when cluster became idle (most recent job finished) and became eligible for deletion due to idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	IdleStartTime string `json:"idleStartTime,omitempty"`
}

type ClusterManagedGroupConfigStatus struct {
	/* Output only. The name of the Instance Group Manager for this group. */
	InstanceGroupManagerName string `json:"instanceGroupManagerName,omitempty"`

	/* Output only. The name of the Instance Template used for the Managed Instance Group. */
	InstanceTemplateName string `json:"instanceTemplateName,omitempty"`
}

type ClusterMasterConfigStatus struct {
	/* Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group. */
	InstanceNames []string `json:"instanceNames,omitempty"`

	/* Output only. Specifies that this instance group contains preemptible instances. */
	IsPreemptible bool `json:"isPreemptible,omitempty"`

	/* Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups. */
	ManagedGroupConfig ClusterManagedGroupConfigStatus `json:"managedGroupConfig,omitempty"`
}

type ClusterMetricsStatus struct {
	/* The HDFS metrics. */
	HdfsMetrics map[string]string `json:"hdfsMetrics,omitempty"`

	/* The YARN metrics. */
	YarnMetrics map[string]string `json:"yarnMetrics,omitempty"`
}

type ClusterSecondaryWorkerConfigStatus struct {
	/* Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group. */
	InstanceNames []string `json:"instanceNames,omitempty"`

	/* Output only. Specifies that this instance group contains preemptible instances. */
	IsPreemptible bool `json:"isPreemptible,omitempty"`

	/* Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups. */
	ManagedGroupConfig ClusterManagedGroupConfigStatus `json:"managedGroupConfig,omitempty"`
}

type ClusterStatusHistoryStatus struct {
	/* Optional. Output only. Details of cluster's state. */
	Detail string `json:"detail,omitempty"`

	/* Output only. The cluster's state. Possible values: UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING, STOPPED, STARTING */
	State string `json:"state,omitempty"`

	/* Output only. Time when this state was entered (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	StateStartTime string `json:"stateStartTime,omitempty"`

	/* Output only. Additional state information that includes status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY, STALE_STATUS */
	Substate string `json:"substate,omitempty"`
}

type ClusterStatusStatus struct {
	/* Optional. Output only. Details of cluster's state. */
	Detail string `json:"detail,omitempty"`

	/* Output only. The cluster's state. Possible values: UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING, STOPPED, STARTING */
	State string `json:"state,omitempty"`

	/* Output only. Time when this state was entered (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)). */
	StateStartTime string `json:"stateStartTime,omitempty"`

	/* Output only. Additional state information that includes status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY, STALE_STATUS */
	Substate string `json:"substate,omitempty"`
}

type ClusterWorkerConfigStatus struct {
	/* Output only. The list of instance names. Dataproc derives the names from `cluster_name`, `num_instances`, and the instance group. */
	InstanceNames []string `json:"instanceNames,omitempty"`

	/* Output only. Specifies that this instance group contains preemptible instances. */
	IsPreemptible bool `json:"isPreemptible,omitempty"`

	/* Output only. The config for Compute Engine Instance Group Manager that manages this group. This is only used for preemptible instance groups. */
	ManagedGroupConfig ClusterManagedGroupConfigStatus `json:"managedGroupConfig,omitempty"`
}

type DataprocClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   DataprocCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster. */
	ClusterUuid string `json:"clusterUuid,omitempty"`
	/*  */
	Config ClusterConfigStatus `json:"config,omitempty"`
	/* Output only. Contains cluster daemon metrics such as HDFS and YARN stats. **Beta Feature**: This report is available for testing purposes only. It may be changed before final release. */
	Metrics ClusterMetricsStatus `json:"metrics,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. Cluster status. */
	Status ClusterStatusStatus `json:"status,omitempty"`
	/* Output only. The previous cluster status. */
	StatusHistory []ClusterStatusHistoryStatus `json:"statusHistory,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataprocCluster is the Schema for the dataproc API
// +k8s:openapi-gen=true
type DataprocCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DataprocClusterSpec   `json:"spec,omitempty"`
	Status DataprocClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataprocClusterList contains a list of DataprocCluster
type DataprocClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocCluster{}, &DataprocClusterList{})
}
