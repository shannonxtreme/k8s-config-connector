package google

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	container "google.golang.org/api/container/v1beta1"
)

// Matches gke-default scope from https://cloud.google.com/sdk/gcloud/reference/container/clusters/create
var defaultOauthScopes = []string{
	"https://www.googleapis.com/auth/devstorage.read_only",
	"https://www.googleapis.com/auth/logging.write",
	"https://www.googleapis.com/auth/monitoring",
	"https://www.googleapis.com/auth/service.management.readonly",
	"https://www.googleapis.com/auth/servicecontrol",
	"https://www.googleapis.com/auth/trace.append",
}

func schemaNodeConfig() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Computed:    true,
		ForceNew:    true,
		Description: `The configuration of the nodepool`,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"disk_size_gb": {
					Type:         schema.TypeInt,
					Optional:     true,
					Computed:     true,
					ForceNew:     true,
					ValidateFunc: validation.IntAtLeast(10),
					Description:  `Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.`,
				},

				"disk_type": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
					ForceNew:    true,
					Description: `Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd`,
				},

				"guest_accelerator": {
					Type:     schema.TypeList,
					Optional: true,
					Computed: true,
					ForceNew: true,
					// Legacy config mode allows removing GPU's from an existing resource
					// See https://www.terraform.io/docs/configuration/attr-as-blocks.html
					ConfigMode:  schema.SchemaConfigModeAttr,
					Description: `List of the type and count of accelerator cards attached to the instance.`,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"count": {
								Type:        schema.TypeInt,
								Required:    true,
								ForceNew:    true,
								Description: `The number of the accelerator cards exposed to an instance.`,
							},
							"type": {
								Type:             schema.TypeString,
								Required:         true,
								ForceNew:         true,
								DiffSuppressFunc: compareSelfLinkOrResourceName,
								Description:      `The accelerator type resource name.`,
							},
							"gpu_partition_size": {
								Type:        schema.TypeString,
								Optional:    true,
								ForceNew:    true,
								Description: `Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning)`,
							},
						},
					},
				},

				"image_type": {
					Type:             schema.TypeString,
					Optional:         true,
					Computed:         true,
					DiffSuppressFunc: caseDiffSuppress,
					Description:      `The image type to use for this node. Note that for a given image type, the latest version of it will be used.`,
				},

				"labels": {
					Type:     schema.TypeMap,
					Optional: true,
					// Computed=true because GKE Sandbox will automatically add labels to nodes that can/cannot run sandboxed pods.
					Computed:         true,
					ForceNew:         true,
					Elem:             &schema.Schema{Type: schema.TypeString},
					Description:      `The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.`,
					DiffSuppressFunc: containerNodePoolLabelsSuppress,
				},

				"local_ssd_count": {
					Type:         schema.TypeInt,
					Optional:     true,
					Computed:     true,
					ForceNew:     true,
					ValidateFunc: validation.IntAtLeast(0),
					Description:  `The number of local SSD disks to be attached to the node.`,
				},

				"ephemeral_storage_config": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: `Parameters for the ephemeral storage filesystem.`,
					ForceNew:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"local_ssd_count": {
								Type:         schema.TypeInt,
								Required:     true,
								ForceNew:     true,
								ValidateFunc: validation.IntAtLeast(0),
								Description:  `Number of local SSDs to use to back ephemeral storage. Uses NVMe interfaces. Each local SSD is 375 GB in size.`,
							},
						},
					},
				},

				"gcfs_config": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: `GCFS configuration for this node.`,
					ForceNew:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"enabled": {
								Type:        schema.TypeBool,
								Required:    true,
								ForceNew:    true,
								Description: `Whether or not GCFS is enabled`,
							},
						},
					},
				},

				"gvnic": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: `Enable or disable gvnic in the node pool.`,
					ForceNew:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"enabled": {
								Type:        schema.TypeBool,
								Required:    true,
								ForceNew:    true,
								Description: `Whether or not gvnic is enabled`,
							},
						},
					},
				},

				"machine_type": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
					ForceNew:    true,
					Description: `The name of a Google Compute Engine machine type.`,
				},

				"metadata": {
					Type:        schema.TypeMap,
					Optional:    true,
					Computed:    true,
					ForceNew:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: `The metadata key/value pairs assigned to instances in the cluster.`,
				},

				"min_cpu_platform": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
					ForceNew:    true,
					Description: `Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform.`,
				},

				"oauth_scopes": {
					Type:        schema.TypeSet,
					Optional:    true,
					Computed:    true,
					ForceNew:    true,
					Description: `The set of Google API scopes to be made available on all of the node VMs.`,
					Elem: &schema.Schema{
						Type: schema.TypeString,
						StateFunc: func(v interface{}) string {
							return canonicalizeServiceScope(v.(string))
						},
					},
					DiffSuppressFunc: containerClusterAddedScopesSuppress,
					Set:              stringScopeHashcode,
				},

				"preemptible": {
					Type:        schema.TypeBool,
					Optional:    true,
					ForceNew:    true,
					Default:     false,
					Description: `Whether the nodes are created as preemptible VM instances.`,
				},

				"spot": {
					Type:        schema.TypeBool,
					Optional:    true,
					ForceNew:    true,
					Default:     false,
					Description: `Whether the nodes are created as spot VM instances.`,
				},

				"service_account": {
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
					ForceNew:    true,
					Description: `The Google Cloud Platform Service Account to be used by the node VMs.`,
				},

				"tags": {
					Type:        schema.TypeList,
					Optional:    true,
					ForceNew:    true,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Description: `The list of instance tags applied to all nodes.`,
				},

				"shielded_instance_config": {
					Type:        schema.TypeList,
					Optional:    true,
					Computed:    true,
					ForceNew:    true,
					Description: `Shielded Instance options.`,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"enable_secure_boot": {
								Type:        schema.TypeBool,
								Optional:    true,
								ForceNew:    true,
								Default:     false,
								Description: `Defines whether the instance has Secure Boot enabled.`,
							},
							"enable_integrity_monitoring": {
								Type:        schema.TypeBool,
								Optional:    true,
								ForceNew:    true,
								Default:     true,
								Description: `Defines whether the instance has integrity monitoring enabled.`,
							},
						},
					},
				},

				"taint": {
					Type:     schema.TypeList,
					Optional: true,
					// Computed=true because GKE Sandbox will automatically add taints to nodes that can/cannot run sandboxed pods.
					Computed: true,
					ForceNew: true,
					// Legacy config mode allows explicitly defining an empty taint.
					// See https://www.terraform.io/docs/configuration/attr-as-blocks.html
					ConfigMode:       schema.SchemaConfigModeAttr,
					Description:      `List of Kubernetes taints to be applied to each node.`,
					DiffSuppressFunc: containerNodePoolTaintSuppress,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"key": {
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: `Key for taint.`,
							},
							"value": {
								Type:        schema.TypeString,
								Required:    true,
								ForceNew:    true,
								Description: `Value for taint.`,
							},
							"effect": {
								Type:         schema.TypeString,
								Required:     true,
								ForceNew:     true,
								ValidateFunc: validation.StringInSlice([]string{"NO_SCHEDULE", "PREFER_NO_SCHEDULE", "NO_EXECUTE"}, false),
								Description:  `Effect for taint.`,
							},
						},
					},
				},

				"workload_metadata_config": {
					Computed:    true,
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: `The workload metadata configuration for this node.`,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"node_metadata": {
								Type:         schema.TypeString,
								Optional:     true,
								Computed:     true,
								Deprecated:   "Deprecated in favor of mode.",
								ValidateFunc: validation.StringInSlice([]string{"UNSPECIFIED", "SECURE", "EXPOSE", "GKE_METADATA_SERVER"}, false),
								Description:  `NodeMetadata is the configuration for how to expose metadata to the workloads running on the node.`,
							},
							"mode": {
								Type:         schema.TypeString,
								Optional:     true,
								Computed:     true,
								ValidateFunc: validation.StringInSlice([]string{"MODE_UNSPECIFIED", "GCE_METADATA", "GKE_METADATA"}, false),
								Description:  `Mode is the configuration for how to expose metadata to workloads running on the node.`,
							},
						},
					},
				},

				"sandbox_config": {
					Type:        schema.TypeList,
					Optional:    true,
					ForceNew:    true,
					MaxItems:    1,
					Description: `Sandbox configuration for this node.`,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"sandbox_type": {
								Type:         schema.TypeString,
								Required:     true,
								Description:  `Type of the sandbox to use for the node (e.g. 'gvisor')`,
								ValidateFunc: validation.StringInSlice([]string{"gvisor"}, false),
							},
						},
					},
				},
				"boot_disk_kms_key": {
					Type:        schema.TypeString,
					Optional:    true,
					ForceNew:    true,
					Description: `The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool.`,
				},
				// Note that AtLeastOneOf can't be set because this schema is reused by
				// two different resources.
				"kubelet_config": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: `Node kubelet configs.`,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"cpu_manager_policy": {
								Type:         schema.TypeString,
								Required:     true,
								ValidateFunc: validation.StringInSlice([]string{"static", "none", ""}, false),
								Description:  `Control the CPU management policy on the node.`,
							},
							"cpu_cfs_quota": {
								Type:        schema.TypeBool,
								Optional:    true,
								Description: `Enable CPU CFS quota enforcement for containers that specify CPU limits.`,
							},
							"cpu_cfs_quota_period": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: `Set the CPU CFS quota period value 'cpu.cfs_period_us'.`,
							},
						},
					},
				},

				"linux_node_config": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: `Parameters that can be configured on Linux nodes.`,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"sysctls": {
								Type:        schema.TypeMap,
								Required:    true,
								Elem:        &schema.Schema{Type: schema.TypeString},
								Description: `The Linux kernel parameters to be applied to the nodes and all pods running on the nodes.`,
							},
						},
					},
				},
				"node_group": {
					Type:        schema.TypeString,
					Optional:    true,
					ForceNew:    true,
					Description: `Setting this field will assign instances of this pool to run on the specified node group. This is useful for running workloads on sole tenant nodes.`,
				},
			},
		},
	}
}

func expandNodeConfig(v interface{}) *container.NodeConfig {
	nodeConfigs := v.([]interface{})
	nc := &container.NodeConfig{
		// Defaults can't be set on a list/set in the schema, so set the default on create here.
		OauthScopes: defaultOauthScopes,
	}
	if len(nodeConfigs) == 0 {
		return nc
	}

	nodeConfig := nodeConfigs[0].(map[string]interface{})

	if v, ok := nodeConfig["machine_type"]; ok {
		nc.MachineType = v.(string)
	}

	if v, ok := nodeConfig["guest_accelerator"]; ok {
		accels := v.([]interface{})
		guestAccelerators := make([]*container.AcceleratorConfig, 0, len(accels))
		for _, raw := range accels {
			data := raw.(map[string]interface{})
			if data["count"].(int) == 0 {
				continue
			}
			guestAccelerators = append(guestAccelerators, &container.AcceleratorConfig{
				AcceleratorCount: int64(data["count"].(int)),
				AcceleratorType:  data["type"].(string),
				GpuPartitionSize: data["gpu_partition_size"].(string),
			})
		}
		nc.Accelerators = guestAccelerators
	}

	if v, ok := nodeConfig["disk_size_gb"]; ok {
		nc.DiskSizeGb = int64(v.(int))
	}

	if v, ok := nodeConfig["disk_type"]; ok {
		nc.DiskType = v.(string)
	}

	if v, ok := nodeConfig["local_ssd_count"]; ok {
		nc.LocalSsdCount = int64(v.(int))
	}

	if v, ok := nodeConfig["ephemeral_storage_config"]; ok && len(v.([]interface{})) > 0 {
		conf := v.([]interface{})[0].(map[string]interface{})
		nc.EphemeralStorageConfig = &container.EphemeralStorageConfig{
			LocalSsdCount: int64(conf["local_ssd_count"].(int)),
		}
	}

	if v, ok := nodeConfig["gcfs_config"]; ok && len(v.([]interface{})) > 0 {
		conf := v.([]interface{})[0].(map[string]interface{})
		nc.GcfsConfig = &container.GcfsConfig{
			Enabled: conf["enabled"].(bool),
		}
	}

	if v, ok := nodeConfig["gvnic"]; ok && len(v.([]interface{})) > 0 {
		conf := v.([]interface{})[0].(map[string]interface{})
		nc.Gvnic = &container.VirtualNIC{
			Enabled: conf["enabled"].(bool),
		}
	}

	if scopes, ok := nodeConfig["oauth_scopes"]; ok {
		scopesSet := scopes.(*schema.Set)
		scopes := make([]string, scopesSet.Len())
		for i, scope := range scopesSet.List() {
			scopes[i] = canonicalizeServiceScope(scope.(string))
		}

		nc.OauthScopes = scopes
	}

	if v, ok := nodeConfig["service_account"]; ok {
		nc.ServiceAccount = v.(string)
	}

	if v, ok := nodeConfig["metadata"]; ok {
		m := make(map[string]string)
		for k, val := range v.(map[string]interface{}) {
			m[k] = val.(string)
		}
		nc.Metadata = m
	}

	if v, ok := nodeConfig["image_type"]; ok {
		nc.ImageType = v.(string)
	}

	if v, ok := nodeConfig["labels"]; ok {
		m := make(map[string]string)
		for k, val := range v.(map[string]interface{}) {
			m[k] = val.(string)
		}
		nc.Labels = m
	}

	if v, ok := nodeConfig["tags"]; ok {
		tagsList := v.([]interface{})
		tags := []string{}
		for _, v := range tagsList {
			if v != nil {
				tags = append(tags, v.(string))
			}
		}
		nc.Tags = tags
	}

	if v, ok := nodeConfig["shielded_instance_config"]; ok && len(v.([]interface{})) > 0 {
		conf := v.([]interface{})[0].(map[string]interface{})
		nc.ShieldedInstanceConfig = &container.ShieldedInstanceConfig{
			EnableSecureBoot:          conf["enable_secure_boot"].(bool),
			EnableIntegrityMonitoring: conf["enable_integrity_monitoring"].(bool),
		}
	}

	// Preemptible Is Optional+Default, so it always has a value
	nc.Preemptible = nodeConfig["preemptible"].(bool)

	// Spot Is Optional+Default, so it always has a value
	nc.Spot = nodeConfig["spot"].(bool)

	if v, ok := nodeConfig["min_cpu_platform"]; ok {
		nc.MinCpuPlatform = v.(string)
	}

	if v, ok := nodeConfig["taint"]; ok && len(v.([]interface{})) > 0 {
		taints := v.([]interface{})
		nodeTaints := make([]*container.NodeTaint, 0, len(taints))
		for _, raw := range taints {
			data := raw.(map[string]interface{})
			taint := &container.NodeTaint{
				Key:    data["key"].(string),
				Value:  data["value"].(string),
				Effect: data["effect"].(string),
			}
			nodeTaints = append(nodeTaints, taint)
		}
		nc.Taints = nodeTaints
	}

	if v, ok := nodeConfig["workload_metadata_config"]; ok {
		nc.WorkloadMetadataConfig = expandWorkloadMetadataConfig(v)
	}

	if v, ok := nodeConfig["sandbox_config"]; ok && len(v.([]interface{})) > 0 {
		conf := v.([]interface{})[0].(map[string]interface{})
		nc.SandboxConfig = &container.SandboxConfig{
			SandboxType: conf["sandbox_type"].(string),
		}
	}
	if v, ok := nodeConfig["boot_disk_kms_key"]; ok {
		nc.BootDiskKmsKey = v.(string)
	}

	if v, ok := nodeConfig["kubelet_config"]; ok {
		nc.KubeletConfig = expandKubeletConfig(v)
	}

	if v, ok := nodeConfig["linux_node_config"]; ok {
		nc.LinuxNodeConfig = expandLinuxNodeConfig(v)
	}

	if v, ok := nodeConfig["node_group"]; ok {
		nc.NodeGroup = v.(string)
	}

	return nc
}

func expandWorkloadMetadataConfig(v interface{}) *container.WorkloadMetadataConfig {
	if v == nil {
		return nil
	}
	ls := v.([]interface{})
	if len(ls) == 0 {
		return nil
	}
	wmc := &container.WorkloadMetadataConfig{}

	cfg := ls[0].(map[string]interface{})

	if v, ok := cfg["mode"]; ok {
		wmc.Mode = v.(string)
	}

	if v, ok := cfg["node_metadata"]; ok {
		wmc.NodeMetadata = v.(string)
	}

	return wmc
}

func expandKubeletConfig(v interface{}) *container.NodeKubeletConfig {
	if v == nil {
		return nil
	}
	ls := v.([]interface{})
	if len(ls) == 0 {
		return nil
	}
	cfg := ls[0].(map[string]interface{})
	kConfig := &container.NodeKubeletConfig{}
	if cpuManagerPolicy, ok := cfg["cpu_manager_policy"]; ok {
		kConfig.CpuManagerPolicy = cpuManagerPolicy.(string)
	}
	if cpuCfsQuota, ok := cfg["cpu_cfs_quota"]; ok {
		kConfig.CpuCfsQuota = cpuCfsQuota.(bool)
		kConfig.ForceSendFields = append(kConfig.ForceSendFields, "CpuCfsQuota")
	}
	if cpuCfsQuotaPeriod, ok := cfg["cpu_cfs_quota_period"]; ok {
		kConfig.CpuCfsQuotaPeriod = cpuCfsQuotaPeriod.(string)
	}
	return kConfig
}

func expandLinuxNodeConfig(v interface{}) *container.LinuxNodeConfig {
	if v == nil {
		return nil
	}
	ls := v.([]interface{})
	if len(ls) == 0 {
		return nil
	}
	cfg := ls[0].(map[string]interface{})
	sysCfgRaw, ok := cfg["sysctls"]
	if !ok {
		return nil
	}
	m := make(map[string]string)
	for k, v := range sysCfgRaw.(map[string]interface{}) {
		m[k] = v.(string)
	}
	return &container.LinuxNodeConfig{
		Sysctls: m,
	}
}

func flattenNodeConfig(c *container.NodeConfig) []map[string]interface{} {
	config := make([]map[string]interface{}, 0, 1)

	if c == nil {
		return config
	}

	config = append(config, map[string]interface{}{
		"machine_type":             c.MachineType,
		"disk_size_gb":             c.DiskSizeGb,
		"disk_type":                c.DiskType,
		"guest_accelerator":        flattenContainerGuestAccelerators(c.Accelerators),
		"local_ssd_count":          c.LocalSsdCount,
		"ephemeral_storage_config": flattenEphemeralStorageConfig(c.EphemeralStorageConfig),
		"gcfs_config":              flattenGcfsConfig(c.GcfsConfig),
		"gvnic":                    flattenGvnic(c.Gvnic),
		"service_account":          c.ServiceAccount,
		"metadata":                 c.Metadata,
		"image_type":               c.ImageType,
		"labels":                   c.Labels,
		"tags":                     c.Tags,
		"preemptible":              c.Preemptible,
		"spot":                     c.Spot,
		"min_cpu_platform":         c.MinCpuPlatform,
		"shielded_instance_config": flattenShieldedInstanceConfig(c.ShieldedInstanceConfig),
		"taint":                    flattenTaints(c.Taints),
		"workload_metadata_config": flattenWorkloadMetadataConfig(c.WorkloadMetadataConfig),
		"sandbox_config":           flattenSandboxConfig(c.SandboxConfig),
		"boot_disk_kms_key":        c.BootDiskKmsKey,
		"kubelet_config":           flattenKubeletConfig(c.KubeletConfig),
		"linux_node_config":        flattenLinuxNodeConfig(c.LinuxNodeConfig),
		"node_group":               c.NodeGroup,
	})

	if len(c.OauthScopes) > 0 {
		config[0]["oauth_scopes"] = schema.NewSet(stringScopeHashcode, convertStringArrToInterface(c.OauthScopes))
	}

	return config
}

func flattenContainerGuestAccelerators(c []*container.AcceleratorConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, accel := range c {
		result = append(result, map[string]interface{}{
			"count":              accel.AcceleratorCount,
			"type":               accel.AcceleratorType,
			"gpu_partition_size": accel.GpuPartitionSize,
		})
	}
	return result
}

func flattenShieldedInstanceConfig(c *container.ShieldedInstanceConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"enable_secure_boot":          c.EnableSecureBoot,
			"enable_integrity_monitoring": c.EnableIntegrityMonitoring,
		})
	}
	return result
}

func flattenEphemeralStorageConfig(c *container.EphemeralStorageConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"local_ssd_count": c.LocalSsdCount,
		})
	}
	return result
}

func flattenGcfsConfig(c *container.GcfsConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"enabled": c.Enabled,
		})
	}
	return result
}

func flattenGvnic(c *container.VirtualNIC) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"enabled": c.Enabled,
		})
	}
	return result
}

func flattenTaints(c []*container.NodeTaint) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, taint := range c {
		result = append(result, map[string]interface{}{
			"key":    taint.Key,
			"value":  taint.Value,
			"effect": taint.Effect,
		})
	}
	return result
}

func flattenWorkloadMetadataConfig(c *container.WorkloadMetadataConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"mode":          c.Mode,
			"node_metadata": c.NodeMetadata,
		})
	}
	return result
}
func flattenSandboxConfig(c *container.SandboxConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"sandbox_type": c.SandboxType,
		})
	}
	return result
}

func containerNodePoolLabelsSuppress(k, old, new string, d *schema.ResourceData) bool {
	// Node configs are embedded into multiple resources (container cluster and
	// container node pool) so we determine the node config key dynamically.
	idx := strings.Index(k, ".labels.")
	if idx < 0 {
		return false
	}

	root := k[:idx]

	// Right now, GKE only applies its own out-of-band labels when you enable
	// Sandbox. We only need to perform diff suppression in this case;
	// otherwise, the default Terraform behavior is fine.
	o, n := d.GetChange(root + ".sandbox_config.0.sandbox_type")
	if o == nil || n == nil {
		return false
	}

	// Pull the entire changeset as a list rather than trying to deal with each
	// element individually.
	o, n = d.GetChange(root + ".labels")
	if o == nil || n == nil {
		return false
	}

	labels := n.(map[string]interface{})

	// Remove all current labels, skipping GKE-managed ones if not present in
	// the new configuration.
	for key, value := range o.(map[string]interface{}) {
		if nv, ok := labels[key]; ok && nv == value {
			delete(labels, key)
		} else if !strings.HasPrefix(key, "sandbox.gke.io/") {
			// User-provided label removed in new configuration.
			return false
		}
	}

	// If, at this point, the map still has elements, the new configuration
	// added an additional taint.
	if len(labels) > 0 {
		return false
	}

	return true
}

func containerNodePoolTaintSuppress(k, old, new string, d *schema.ResourceData) bool {
	// Node configs are embedded into multiple resources (container cluster and
	// container node pool) so we determine the node config key dynamically.
	idx := strings.Index(k, ".taint.")
	if idx < 0 {
		return false
	}

	root := k[:idx]

	// Right now, GKE only applies its own out-of-band labels when you enable
	// Sandbox. We only need to perform diff suppression in this case;
	// otherwise, the default Terraform behavior is fine.
	o, n := d.GetChange(root + ".sandbox_config.0.sandbox_type")
	if o == nil || n == nil {
		return false
	}

	// Pull the entire changeset as a list rather than trying to deal with each
	// element individually.
	o, n = d.GetChange(root + ".taint")
	if o == nil || n == nil {
		return false
	}

	type taintType struct {
		Key, Value, Effect string
	}

	taintSet := make(map[taintType]struct{})

	// Add all new taints to set.
	for _, raw := range n.([]interface{}) {
		data := raw.(map[string]interface{})
		taint := taintType{
			Key:    data["key"].(string),
			Value:  data["value"].(string),
			Effect: data["effect"].(string),
		}
		taintSet[taint] = struct{}{}
	}

	// Remove all current taints, skipping GKE-managed keys if not present in
	// the new configuration.
	for _, raw := range o.([]interface{}) {
		data := raw.(map[string]interface{})
		taint := taintType{
			Key:    data["key"].(string),
			Value:  data["value"].(string),
			Effect: data["effect"].(string),
		}
		if _, ok := taintSet[taint]; ok {
			delete(taintSet, taint)
		} else if !strings.HasPrefix(taint.Key, "sandbox.gke.io/") {
			// User-provided taint removed in new configuration.
			return false
		}
	}

	// If, at this point, the set still has elements, the new configuration
	// added an additional taint.
	if len(taintSet) > 0 {
		return false
	}

	return true
}

func flattenKubeletConfig(c *container.NodeKubeletConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"cpu_cfs_quota":        c.CpuCfsQuota,
			"cpu_cfs_quota_period": c.CpuCfsQuotaPeriod,
			"cpu_manager_policy":   c.CpuManagerPolicy,
		})
	}
	return result
}

func flattenLinuxNodeConfig(c *container.LinuxNodeConfig) []map[string]interface{} {
	result := []map[string]interface{}{}
	if c != nil {
		result = append(result, map[string]interface{}{
			"sysctls": c.Sysctls,
		})
	}
	return result
}
