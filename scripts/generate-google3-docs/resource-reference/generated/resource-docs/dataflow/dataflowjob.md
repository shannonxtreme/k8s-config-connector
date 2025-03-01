{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}DataflowJob{% endblock %}
{% block body %}



Note: You can update the spec of DataflowJob resources when the type is
<code>JOB_TYPE_STREAMING</code>. When the type is <code>JOB_TYPE_BATCH</code>,
the spec cannot be updated. When you update a DataflowJob,
{{product_name_short}} replaces the Cloud Dataflow Job with a new job.
Read more about updating Dataflow jobs at
<a href="/dataflow/docs/guides/updating-a-pipeline">Updating an existing pipeline</a>.

<table>
<thead>
<tr>
<th><strong>Property</strong></th>
<th><strong>Value</strong></th>
</tr>
</thead>
<tbody>
<tr>
<td>{{gcp_name_short}} Service Name</td>
<td>Cloud Dataflow</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/dataflow/docs/">/dataflow/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>v1b3.projects.jobs</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/dataflow/docs/reference/rest/v1b3/projects.jobs">/dataflow/docs/reference/rest/v1b3/projects.jobs</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpdataflowjob<br>gcpdataflowjobs<br>dataflowjob</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>dataflow.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>dataflowjobs.dataflow.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties


### Annotations
<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>cnrm.cloud.google.com/on-delete</code></td>
    </tr>
    <tr>
        <td><code>cnrm.cloud.google.com/project-id</code></td>
    </tr>
    <tr>
        <td><code>cnrm.cloud.google.com/skip-wait-on-job-termination</code></td>
    </tr>
</tbody>
</table>


### Spec
#### Schema
```yaml
additionalExperiments:
- string
enableStreamingEngine: boolean
ipConfiguration: string
kmsKeyRef:
  external: string
  name: string
  namespace: string
machineType: string
maxWorkers: integer
networkRef:
  external: string
  name: string
  namespace: string
parameters: {}
region: string
resourceID: string
serviceAccountRef:
  external: string
  name: string
  namespace: string
subnetworkRef:
  external: string
  name: string
  namespace: string
tempGcsLocation: string
templateGcsPath: string
transformNameMapping: {}
zone: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td>
            <p><code>additionalExperiments</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>additionalExperiments[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>enableStreamingEngine</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Indicates if the job should use the streaming engine feature.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>ipConfiguration</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>kmsKeyRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}The name for the Cloud KMS key for the job.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>kmsKeyRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `KMSCryptoKey` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>kmsKeyRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>kmsKeyRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>machineType</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The machine type to use for the job.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>maxWorkers</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. The number of workers permitted to work on the job. More workers may improve processing speed at additional cost.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeNetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>networkRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>parameters</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Key/Value pairs to be passed to the Dataflow job (as used in the template).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>region</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The region in which the created job should run.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>serviceAccountRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>serviceAccountRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `email` field of an `IAMServiceAccount` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>serviceAccountRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>serviceAccountRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The `selfLink` field of a `ComputeSubnetwork` resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>subnetworkRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>tempGcsLocation</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>templateGcsPath</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The Google Cloud Storage path to the Dataflow job template.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>transformNameMapping</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>zone</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The zone in which the created job should run. If it is not provided, the provider zone is used.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>



### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
jobId: string
observedGeneration: integer
state: string
type: string
```

<table class="properties responsive">
<thead>
    <tr>
        <th colspan="2">Fields</th>
    </tr>
</thead>
<tbody>
    <tr>
        <td><code>conditions</code></td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Conditions represent the latest available observation of the resource's current state.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[]</code></td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].lastTransitionTime</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Last time the condition transitioned from one status to another.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].message</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Human-readable message indicating details about last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].reason</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Unique, one-word, CamelCase reason for the condition's last transition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].status</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Status is the status of the condition. Can be True, False, Unknown.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>conditions[].type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Type is the type of the condition.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>jobId</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The unique ID of this job.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>state</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The current state of the resource, selected from the JobState enum.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>type</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The type of this job, selected from the JobType enum.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Batch Dataflow Job
```yaml
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: dataflow.cnrm.cloud.google.com/v1beta1
kind: DataflowJob
metadata:
  annotations:
    cnrm.cloud.google.com/on-delete: "cancel"
  labels:
    label-one: "value-one"
  name: dataflowjob-sample-batch
spec:
  tempGcsLocation: gs://${PROJECT_ID?}-dataflowjob-dep-batch/tmp
  # This is a public, Google-maintained Dataflow Job template of a batch job
  templateGcsPath: gs://dataflow-templates/2020-02-03-01_RC00/Word_Count
  parameters:
    # This is a public, Google-maintained text file
    inputFile: gs://dataflow-samples/shakespeare/various.txt
    output: gs://${PROJECT_ID?}-dataflowjob-dep-batch/output
  zone: us-central1-a
  machineType: "n1-standard-1"
  maxWorkers: 3
  ipConfiguration: "WORKER_IP_PUBLIC"
---
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  annotations:
    cnrm.cloud.google.com/force-destroy: "true"
  # StorageBucket names must be globally unique. Replace ${PROJECT_ID?} with your project ID.
  name: ${PROJECT_ID?}-dataflowjob-dep-batch
```

### Streaming Dataflow Job
```yaml
# Copyright 2020 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

apiVersion: dataflow.cnrm.cloud.google.com/v1beta1
kind: DataflowJob
metadata:
  annotations:
    cnrm.cloud.google.com/on-delete: "cancel"
  labels:
    label-one: "value-one"
  name: dataflowjob-sample-streaming
spec:
  tempGcsLocation: gs://${PROJECT_ID?}-dataflowjob-dep-streaming/tmp
  # This is a public, Google-maintained Dataflow Job template of a streaming job
  templateGcsPath: gs://dataflow-templates/2020-02-03-01_RC00/PubSub_to_BigQuery
  parameters:
    # replace ${PROJECT_ID?} with your project name
    inputTopic: projects/${PROJECT_ID?}/topics/dataflowjob-dep-streaming
    outputTableSpec: ${PROJECT_ID?}:dataflowjobdepstreaming.dataflowjobdepstreaming
  zone: us-central1-a
  machineType: "n1-standard-1"
  maxWorkers: 3
  ipConfiguration: "WORKER_IP_PUBLIC"
---
apiVersion: bigquery.cnrm.cloud.google.com/v1beta1
kind: BigQueryDataset
metadata:
  name: dataflowjobdepstreaming
---
apiVersion: bigquery.cnrm.cloud.google.com/v1beta1
kind: BigQueryTable
metadata:
  name: dataflowjobdepstreaming
spec:
  datasetRef:
    name: dataflowjobdepstreaming
---
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: dataflowjob-dep-streaming
---
apiVersion: storage.cnrm.cloud.google.com/v1beta1
kind: StorageBucket
metadata:
  annotations:
    cnrm.cloud.google.com/force-destroy: "true"
  # StorageBucket names must be globally unique. Replace ${PROJECT_ID?} with your project ID.
  name: ${PROJECT_ID?}-dataflowjob-dep-streaming
```


{% endblock %}