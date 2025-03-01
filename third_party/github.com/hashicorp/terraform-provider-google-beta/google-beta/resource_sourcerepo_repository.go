// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSourceRepoRepositoryPubSubConfigsHash(v interface{}) int {
	if v == nil {
		return 0
	}

	var buf bytes.Buffer
	m := v.(map[string]interface{})

	buf.WriteString(fmt.Sprintf("%s-", GetResourceNameFromSelfLink(m["topic"].(string))))
	buf.WriteString(fmt.Sprintf("%s-", m["message_format"].(string)))
	if v, ok := m["service_account_email"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", v.(string)))
	}

	return hashcode(buf.String())
}

func resourceSourceRepoRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceSourceRepoRepositoryCreate,
		Read:   resourceSourceRepoRepositoryRead,
		Update: resourceSourceRepoRepositoryUpdate,
		Delete: resourceSourceRepoRepositoryDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSourceRepoRepositoryImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Resource name of the repository, of the form '{{repo}}'.
The repo name may contain slashes. eg, 'name/with/slash'`,
			},
			"pubsub_configs": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `How this repository publishes a change in the repository through Cloud Pub/Sub. 
Keyed by the topic names.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"message_format": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateEnum([]string{"PROTOBUF", "JSON"}),
							Description: `The format of the Cloud Pub/Sub messages. 
- PROTOBUF: The message payload is a serialized protocol buffer of SourceRepoEvent.
- JSON: The message payload is a JSON string of SourceRepoEvent. Possible values: ["PROTOBUF", "JSON"]`,
						},
						"service_account_email": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `Email address of the service account used for publishing Cloud Pub/Sub messages. 
This service account needs to be in the same project as the PubsubConfig. When added, 
the caller needs to have iam.serviceAccounts.actAs permission on this service account. 
If unspecified, it defaults to the compute engine default service account.`,
						},
					},
				},
				Set: resourceSourceRepoRepositoryPubSubConfigsHash,
			},
			"size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: `The disk usage of the repo, in bytes.`,
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `URL to clone the repository from Google Cloud Source Repositories.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSourceRepoRepositoryCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandSourceRepoRepositoryName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	pubsubConfigsProp, err := expandSourceRepoRepositoryPubsubConfigs(d.Get("pubsub_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pubsub_configs"); !isEmptyValue(reflect.ValueOf(pubsubConfigsProp)) && (ok || !reflect.DeepEqual(v, pubsubConfigsProp)) {
		obj["pubsubConfigs"] = pubsubConfigsProp
	}

	url, err := replaceVars(d, config, "{{SourceRepoBasePath}}projects/{{project}}/repos")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Repository: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Repository: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	if v, ok := d.GetOkExists("pubsub_configs"); !isEmptyValue(reflect.ValueOf(pubsubConfigsProp)) && (ok || !reflect.DeepEqual(v, pubsubConfigsProp)) {
		log.Printf("[DEBUG] Calling update after create to patch in pubsub_configs")
		// pubsub_configs cannot be added on create
		return resourceSourceRepoRepositoryUpdate(d, meta)
	}

	log.Printf("[DEBUG] Finished creating Repository %q: %#v", d.Id(), res)

	return resourceSourceRepoRepositoryRead(d, meta)
}

func resourceSourceRepoRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{SourceRepoBasePath}}projects/{{project}}/repos/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("SourceRepoRepository %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	if err := d.Set("name", flattenSourceRepoRepositoryName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("url", flattenSourceRepoRepositoryUrl(res["url"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("size", flattenSourceRepoRepositorySize(res["size"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}
	if err := d.Set("pubsub_configs", flattenSourceRepoRepositoryPubsubConfigs(res["pubsubConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading Repository: %s", err)
	}

	return nil
}

func resourceSourceRepoRepositoryUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	pubsubConfigsProp, err := expandSourceRepoRepositoryPubsubConfigs(d.Get("pubsub_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("pubsub_configs"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, pubsubConfigsProp)) {
		obj["pubsubConfigs"] = pubsubConfigsProp
	}

	obj, err = resourceSourceRepoRepositoryUpdateEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{SourceRepoBasePath}}projects/{{project}}/repos/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Repository %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("pubsub_configs") {
		updateMask = append(updateMask, "pubsubConfigs")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Repository %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Repository %q: %#v", d.Id(), res)
	}

	return resourceSourceRepoRepositoryRead(d, meta)
}

func resourceSourceRepoRepositoryDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Repository: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{SourceRepoBasePath}}projects/{{project}}/repos/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Repository %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Repository")
	}

	log.Printf("[DEBUG] Finished deleting Repository %q: %#v", d.Id(), res)
	return nil
}

func resourceSourceRepoRepositoryImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/repos/(?P<name>.+)",
		"(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenSourceRepoRepositoryName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}

	// We can't use a standard name_from_self_link because the name can include /'s
	parts := strings.SplitAfterN(v.(string), "/", 4)
	return parts[3]
}

func flattenSourceRepoRepositoryUrl(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenSourceRepoRepositorySize(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := stringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenSourceRepoRepositoryPubsubConfigs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.(map[string]interface{})
	transformed := make([]interface{}, 0, len(l))
	for k, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"topic":                 k,
			"message_format":        flattenSourceRepoRepositoryPubsubConfigsMessageFormat(original["messageFormat"], d, config),
			"service_account_email": flattenSourceRepoRepositoryPubsubConfigsServiceAccountEmail(original["serviceAccountEmail"], d, config),
		})
	}
	return transformed
}
func flattenSourceRepoRepositoryPubsubConfigsMessageFormat(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenSourceRepoRepositoryPubsubConfigsServiceAccountEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandSourceRepoRepositoryName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return replaceVars(d, config, "projects/{{project}}/repos/{{name}}")
}

func expandSourceRepoRepositoryPubsubConfigs(v interface{}, d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}
	m := make(map[string]interface{})
	for _, raw := range v.(*schema.Set).List() {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMessageFormat, err := expandSourceRepoRepositoryPubsubConfigsMessageFormat(original["message_format"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMessageFormat); val.IsValid() && !isEmptyValue(val) {
			transformed["messageFormat"] = transformedMessageFormat
		}

		transformedServiceAccountEmail, err := expandSourceRepoRepositoryPubsubConfigsServiceAccountEmail(original["service_account_email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedServiceAccountEmail); val.IsValid() && !isEmptyValue(val) {
			transformed["serviceAccountEmail"] = transformedServiceAccountEmail
		}

		transformedTopic, err := expandSourceRepoRepositoryPubsubConfigsTopic(original["topic"], d, config)
		if err != nil {
			return nil, err
		}
		m[transformedTopic] = transformed
	}
	return m, nil
}

func expandSourceRepoRepositoryPubsubConfigsMessageFormat(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandSourceRepoRepositoryPubsubConfigsServiceAccountEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceSourceRepoRepositoryUpdateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Add "topic" field using pubsubConfig map key
	pubsubConfigsVal := obj["pubsubConfigs"]
	if pubsubConfigsVal != nil {
		pubsubConfigs := pubsubConfigsVal.(map[string]interface{})
		for key := range pubsubConfigs {
			config := pubsubConfigs[key].(map[string]interface{})
			config["topic"] = key
		}
	}

	// Nest request body in "repo" field
	newObj := make(map[string]interface{})
	newObj["repo"] = obj
	return newObj, nil
}
