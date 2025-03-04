{# AUTOGENERATED. DO NOT EDIT. #}

{% extends "config-connector/_base.html" %}

{% block page_title %}BillingBudgetsBudget{% endblock %}
{% block body %}


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
<td>Billing Budgets</td>
</tr>
<tr>
<td>{{gcp_name_short}} Service Documentation</td>
<td><a href="/billing/docs/">/billing/docs/</a></td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Name</td>
<td>billingAccounts.budgets</td>
</tr>
<tr>
<td>{{gcp_name_short}} REST Resource Documentation</td>
<td><a href="/billing/docs/reference/budget/rest/v1beta1/billingAccounts.budgets">/billing/docs/reference/budget/rest/v1beta1/billingAccounts.budgets</a></td>
</tr>
<tr>
<td>{{product_name_short}} Resource Short Names</td>
<td>gcpbillingbudgetsbudget<br>gcpbillingbudgetsbudgets<br>billingbudgetsbudget</td>
</tr>
<tr>
<td>{{product_name_short}} Service Name</td>
<td>billingbudgets.googleapis.com</td>
</tr>
<tr>
<td>{{product_name_short}} Resource Fully Qualified Name</td>
<td>billingbudgetsbudgets.billingbudgets.cnrm.cloud.google.com</td>
</tr>

<tr>
    <td>Can Be Referenced by IAMPolicy/IAMPolicyMember</td>
    <td>No</td>
</tr>


</tbody>
</table>

## Custom Resource Definition Properties



### Spec
#### Schema
```yaml
allUpdatesRule:
  disableDefaultIamRecipients: boolean
  monitoringNotificationChannels:
  - external: string
    name: string
    namespace: string
  pubsubTopicRef:
    external: string
    name: string
    namespace: string
  schemaVersion: string
amount:
  lastPeriodAmount: {}
  specifiedAmount:
    currencyCode: string
    nanos: integer
    units: integer
billingAccountRef:
  external: string
  name: string
  namespace: string
budgetFilter:
  calendarPeriod: string
  creditTypes:
  - string
  creditTypesTreatment: string
  customPeriod:
    endDate:
      day: integer
      month: integer
      year: integer
    startDate:
      day: integer
      month: integer
      year: integer
  labels:
    string: object
  projects:
  - external: string
    name: string
    namespace: string
  services:
  - string
  subaccounts:
  - external: string
    name: string
    namespace: string
displayName: string
resourceID: string
thresholdRules:
- spendBasis: string
  thresholdPercent: float
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
            <p><code>allUpdatesRule</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Rules to apply to notifications sent based on budget spend and thresholds.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.disableDefaultIamRecipients</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">boolean</code></p>
            <p>{% verbatim %}Optional. When set to true, disables default notifications sent when a threshold is exceeded. Default notifications are sent to those with Billing Account Administrator and Billing Account User IAM roles for the target account.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.monitoringNotificationChannels</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.monitoringNotificationChannels[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.monitoringNotificationChannels[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The Google Cloud resource name of a `MonitoringNotificationChannel` resource (format: `projects/{{project}}/notificationChannels/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.monitoringNotificationChannels[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.monitoringNotificationChannels[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.pubsubTopicRef</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.pubsubTopicRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. The name of the Pub/Sub topic where budget related messages will be published, in the form `projects/{project_id}/topics/{topic_id}`. Updates are sent at regular intervals to the topic. The topic needs to be created before the budget is created; see https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications for more details. Caller is expected to have `pubsub.topics.setIamPolicy` permission on the topic when it's set for a budget, otherwise, the API call will fail with PERMISSION_DENIED. See https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#permissions_required_for_this_task for more details on Pub/Sub roles and permissions.

Allowed value: The Google Cloud resource name of a `PubSubTopic` resource (format: `projects/{{project}}/topics/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.pubsubTopicRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.pubsubTopicRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>allUpdatesRule.schemaVersion</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. Required when NotificationsRule.pubsub_topic is set. The schema version of the notification sent to NotificationsRule.pubsub_topic. Only "1.0" is accepted. It represents the JSON schema as defined in https://cloud.google.com/billing/docs/how-to/budgets-programmatic-notifications#notification_format.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>amount</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Required. Budgeted amount.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>amount.lastPeriodAmount</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Use the last period's actual spend as the budget for the present period. LastPeriodAmount can only be set when the budget's time period is a .{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>amount.specifiedAmount</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}A specified amount to use as the budget. `currency_code` is optional. If specified when creating a budget, it must match the currency of the billing account. If specified when updating a budget, it must match the currency_code of the existing budget. The `currency_code` is provided on output.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>amount.specifiedAmount.currencyCode</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. The three-letter currency code defined in ISO 4217.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>amount.specifiedAmount.nanos</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If `units` is positive, `nanos` must be positive or zero. If `units` is zero, `nanos` can be positive, zero, or negative. If `units` is negative, `nanos` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>amount.specifiedAmount.units</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}The whole units of the amount. For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>billingAccountRef</code></p>
            <p><i>Required</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>billingAccountRef.external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}The billing account of the resource

Allowed value: The Google Cloud resource name of a Google Cloud Billing Account (format: `billingAccounts/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>billingAccountRef.name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}[WARNING] BillingAccount not yet supported in Config Connector, use 'external' field to reference existing resources.
Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>billingAccountRef.namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.calendarPeriod</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. Specifies to track usage for recurring calendar period. For example, assume that CalendarPeriod.QUARTER is set. The budget will track usage from April 1 to June 30, when the current calendar month is April, May, June. After that, it will track usage from July 1 to September 30 when the current calendar month is July, August, September, so on. Possible values: CALENDAR_PERIOD_UNSPECIFIED, MONTH, QUARTER, YEAR{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.creditTypes</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. If Filter.credit_types_treatment is INCLUDE_SPECIFIED_CREDITS, this is a list of credit types to be subtracted from gross cost to determine the spend for threshold calculations. See a list of acceptable credit type values. If Filter.credit_types_treatment is not INCLUDE_SPECIFIED_CREDITS, this field must be empty.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.creditTypes[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.creditTypesTreatment</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. If not set, default behavior is `INCLUDE_ALL_CREDITS`.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Optional. Specifies to track usage from any start date (required) to any end date (optional). This time period is static, it does not recur.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.endDate</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. Optional. The end date of the time period. Budgets with elapsed end date won't be processed. If unset, specifies to track all usage incurred since the start_date.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.endDate.day</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.endDate.month</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.endDate.year</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.startDate</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}Immutable. Required. The start date must be after January 1, 2017.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.startDate.day</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.startDate.month</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.customPeriod.startDate.year</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}Immutable. Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.labels</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">map (key: string, value: object)</code></p>
            <p>{% verbatim %}Optional. A single label and value pair specifying that usage from only this set of labeled resources should be included in the budget. Currently, multiple entries or multiple values per entry are not allowed. If omitted, the report will include all labeled and unlabeled usage.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.projects</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.projects[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.projects[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.projects[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.projects[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.services</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (string)</code></p>
            <p>{% verbatim %}Optional. A set of services of the form `services/{service_id}`, specifying that usage from only this set of services should be included in the budget. If omitted, the report will include usage for all the services. The service names are available through the Catalog API: https://cloud.google.com/billing/v1/how-tos/catalog-api.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.services[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.subaccounts</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.subaccounts[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.subaccounts[].external</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.subaccounts[].name</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}[WARNING] CloudBillingBillingAccount not yet supported in Config Connector, use 'external' field to reference existing resources.
Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>budgetFilter.subaccounts[].namespace</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>displayName</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}User data for display name in UI. The name must be less than or equal to 60 characters.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>resourceID</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>thresholdRules</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">list (object)</code></p>
            <p>{% verbatim %}Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>thresholdRules[]</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">object</code></p>
            <p>{% verbatim %}{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>thresholdRules[].spendBasis</code></p>
            <p><i>Optional</i></p>
        </td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. The type of basis used to determine if spend has passed the threshold. Behavior defaults to CURRENT_SPEND if not set. Possible values: BASIS_UNSPECIFIED, CURRENT_SPEND, FORECASTED_SPEND{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td>
            <p><code>thresholdRules[].thresholdPercent</code></p>
            <p><i>Required*</i></p>
        </td>
        <td>
            <p><code class="apitype">float</code></p>
            <p>{% verbatim %}Required. Send an alert when this threshold is exceeded. This is a 1.0-based percentage, so 0.5 = 50%. Validation: non-negative number.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>


<p>{% verbatim %}* Field is required when parent field is specified{% endverbatim %}</p>


### Status
#### Schema
```yaml
conditions:
- lastTransitionTime: string
  message: string
  reason: string
  status: string
  type: string
etag: string
observedGeneration: integer
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
        <td><code>etag</code></td>
        <td>
            <p><code class="apitype">string</code></p>
            <p>{% verbatim %}Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.{% endverbatim %}</p>
        </td>
    </tr>
    <tr>
        <td><code>observedGeneration</code></td>
        <td>
            <p><code class="apitype">integer</code></p>
            <p>{% verbatim %}ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.{% endverbatim %}</p>
        </td>
    </tr>
</tbody>
</table>

## Sample YAML(s)

### Calendar Budget
```yaml
# Copyright 2021 Google LLC
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

apiVersion: billingbudgets.cnrm.cloud.google.com/v1beta1
kind: BillingBudgetsBudget
metadata:
  name: billingbudgetsbudget-sample-calendarbudget
spec:
  billingAccountRef:
    # Replace "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}" with the numeric ID for your billing account
    external: "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}"
  displayName: "sample-budget"
  budgetFilter:
    projects:
    - name: "billingbudgetsbudget-dep-calb"
    creditTypes:
    - "DISCOUNT"
    creditTypesTreatment: "INCLUDE_SPECIFIED_CREDITS"
    services:
    # This is the service name for the Geolocation API.
    - "services/0245-C3C9-3864"
    labels:
      label-one:
        values:
        - "value-one"
    calendarPeriod: "MONTH"
  amount:
    specifiedAmount:
      currencyCode: "USD"
      units: 9000000
      nanos: 0
  thresholdRules:
  - thresholdPercent: 0.5
    spendBasis: "CURRENT_SPEND"
  allUpdatesRule:
    pubsubTopicRef:
      name: "billingbudgetsbudget-dep-calendarbudget"
    schemaVersion: "1.0"
    monitoringNotificationChannels:
    - name: "billingbudgetsbudget-dep-calendarbudget"
    disableDefaultIamRecipients: false
---
apiVersion: monitoring.cnrm.cloud.google.com/v1beta1
kind: MonitoringNotificationChannel
metadata:
  name: billingbudgetsbudget-dep-calendarbudget
spec:
  labels:
    email_address: test@example.com
  type: "email"
---
apiVersion: pubsub.cnrm.cloud.google.com/v1beta1
kind: PubSubTopic
metadata:
  name: billingbudgetsbudget-dep-calendarbudget
---
apiVersion: resourcemanager.cnrm.cloud.google.com/v1beta1
kind: Project
metadata:
  name: billingbudgetsbudget-dep-calb
spec:
  organizationRef:
    # Replace "${ORG_ID?}" with the numeric ID for your organization
    external: "${ORG_ID?}"
  name: "billingbudgetsbudget-dep-calb"
  billingAccountRef:
    # Replace "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}" with the numeric ID for your billing account
    external: "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}"
```

### Custom Budget
```yaml
# Copyright 2021 Google LLC
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

apiVersion: billingbudgets.cnrm.cloud.google.com/v1beta1
kind: BillingBudgetsBudget
metadata:
  name: billingbudgetsbudget-sample-custombudget
spec:
  billingAccountRef:
    # Replace "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}" with the numeric ID for your billing account
    external: "${BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES?}"
  budgetFilter:
    creditTypes:
    - "DISCOUNT"
    creditTypesTreatment: "INCLUDE_SPECIFIED_CREDITS"
    customPeriod:
      startDate:
        year: 2140
        month: 1
        day: 1
      endDate:
        year: 2312
        month: 3
        day: 14
  amount:
    specifiedAmount:
      currencyCode: "USD"
      units: 9000000
      nanos: 0
```


{% endblock %}
