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

package testgcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcp"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/storage/v1"
)

const (
	TestFolderId                            = "TEST_FOLDER_ID"
	TestFolder2Id                           = "TEST_FOLDER_2_ID"
	TestOrgId                               = "TEST_ORG_ID"
	IAMIntegrationTestsOrganizationId       = "IAM_INTEGRATION_TESTS_ORGANIZATION_ID"
	IsolatedTestOrgName                     = "ISOLATED_TEST_ORG_NAME"
	TestBillingAccountId                    = "TEST_BILLING_ACCOUNT_ID"
	TestBillingAccountIDForBillingResources = "BILLING_ACCOUNT_ID_FOR_BILLING_RESOURCES"
	IAMIntegrationTestsBillingAccountId     = "IAM_INTEGRATION_TESTS_BILLING_ACCOUNT_ID"
	FirestoreTestProject                    = "FIRESTORE_TEST_PROJECT"
	CloudFunctionsTestProject               = "CLOUD_FUNCTIONS_TEST_PROJECT"
	IdentityPlatformTestProject             = "IDENTITY_PLATFORM_TEST_PROJECT"
	InterconnectTestProject                 = "INTERCONNECT_TEST_PROJECT"
	HighCPUQuotaTestProject                 = "HIGH_CPU_QUOTA_TEST_PROJECT"
	RecaptchaEnterpriseTestProject          = "RECAPTCHA_ENTERPRISE_TEST_PROJECT"
)

var (
	testFolderID                            = os.Getenv(TestFolderId)
	testFolder2Id                           = os.Getenv(TestFolder2Id)
	testOrgID                               = os.Getenv(TestOrgId)
	isolatedTestOrgName                     = os.Getenv(IsolatedTestOrgName)
	iamIntegrationTestsOrganizationId       = os.Getenv(IAMIntegrationTestsOrganizationId)
	testBillingAccountID                    = os.Getenv(TestBillingAccountId)
	testBillingAccountIDForBillingResources = os.Getenv(TestBillingAccountIDForBillingResources)
	iamIntegrationTestsBillingAccountId     = os.Getenv(IAMIntegrationTestsBillingAccountId)
	firestoreTestProject                    = os.Getenv(FirestoreTestProject)
	cloudFunctionsTestProject               = os.Getenv(CloudFunctionsTestProject)
	identityPlatformTestProject             = os.Getenv(IdentityPlatformTestProject)
	interconnectTestProject                 = os.Getenv(InterconnectTestProject)
	highCpuQuotaTestProject                 = os.Getenv(HighCPUQuotaTestProject)
	recaptchaEnterpriseTestProject          = os.Getenv(RecaptchaEnterpriseTestProject)
)

// GetDefaultProjectID returns the ID of user's configured default GCP project.
func GetDefaultProjectID(t *testing.T) string {
	t.Helper()
	projectID, err := gcp.GetDefaultProjectID()
	if err != nil {
		t.Fatalf("error retrieving gcloud sdk credentials: %v", err)
	}
	return projectID
}

// GetDefaultProjectNumber returns the project number of user's configured default GCP project.
func GetDefaultProjectNumber(t *testing.T) string {
	t.Helper()
	projectNumber, err := gcp.GetDefaultProjectNumber()
	if err != nil {
		t.Fatalf("error retrieving gcloud sdk credentials: %v", err)
	}
	return projectNumber
}

// GetDefaultServiceAccount returns the service account used to access the user's configured default GCP project.
func GetDefaultServiceAccount(t *testing.T) string {
	creds, err := google.FindDefaultCredentials(context.TODO(), cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		t.Fatalf("error getting credentials: %v", err)
	}
	if creds == nil {
		t.Fatalf("test running in environment without default credentials, can't continue")
	}

	var rawCreds map[string]string
	if err := json.Unmarshal(creds.JSON, &rawCreds); err != nil {
		t.Fatalf("creds file malformed: %v", err)
	}

	return rawCreds["client_email"]
}

func GetFolderID(t *testing.T) string {
	return testFolderID
}

func GetFolder2ID(t *testing.T) string {
	return testFolder2Id
}

func GetBillingAccountID(t *testing.T) string {
	return testBillingAccountID
}

func GetTestBillingAccountIDForBillingResources(t *testing.T) string {
	return testBillingAccountIDForBillingResources
}

func GetOrgID(t *testing.T) string {
	return testOrgID
}

func GetIsolatedTestOrgName(t *testing.T) string {
	return isolatedTestOrgName
}

func GetIAMIntegrationTestsBillingAccountId(t *testing.T) string {
	return iamIntegrationTestsBillingAccountId
}

func GetIAMIntegrationTestsOrganizationId(t *testing.T) string {
	return iamIntegrationTestsOrganizationId
}

func GetFirestoreTestProject(t *testing.T) string {
	return firestoreTestProject
}

func GetCloudFunctionsTestProject(t *testing.T) string {
	return cloudFunctionsTestProject
}

func GetIdentityPlatformTestProject(t *testing.T) string {
	return identityPlatformTestProject
}

func GetInterconnectTestProject(t *testing.T) string {
	return interconnectTestProject
}

func GetHighCpuQuotaTestProject(t *testing.T) string {
	return highCpuQuotaTestProject
}
func GetRecaptchaEnterpriseTestProject(t *testing.T) string {
	return recaptchaEnterpriseTestProject
}

// attempts to return a valid IAM policy binding for the current credential by searching for an email in the cloud credentials file and defaulting to the
// current user if on a corp machine.
func GetIAMPolicyBindingMember(t *testing.T) string {
	currentUser, err := user.Current()
	if err != nil {
		t.Fatalf("unable to find current user: %v", err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		t.Fatalf("unable to get hostname: %v", err)
	}
	if serviceAccountEmail := GetDefaultServiceAccount(t); serviceAccountEmail != "" {
		return fmt.Sprintf("serviceAccount:%v", serviceAccountEmail)
	}
	if strings.HasSuffix(hostname, ".corp.google.com") {
		return fmt.Sprintf("user:%s@google.com", currentUser.Username)
	}
	if strings.HasSuffix(hostname, ".c.googlers.com") {
		return fmt.Sprintf("user:%s@google.com", currentUser.Username)
	}
	t.Fatalf("Unable to get safety binding member")
	return ""
}

func NewDefaultHTTPClient(t *testing.T) *http.Client {
	t.Helper()
	client, err := google.DefaultClient(context.TODO(), compute.CloudPlatformScope)
	if err != nil {
		t.Fatalf("error creating default google client: %v", err)
	}
	return client
}

func NewStorageClient(t *testing.T) *storage.Service {
	t.Helper()
	client, err := gcp.NewStorageClient(context.TODO())
	if err != nil {
		t.Fatalf("error creating storage client: %v", err)
	}
	return client
}

func NewResourceManagerClient(t *testing.T) *cloudresourcemanager.Service {
	t.Helper()
	client, err := gcp.NewCloudResourceManagerClient(context.TODO())
	if err != nil {
		t.Fatalf("error creating cloud resource manager client: %v", err)
	}
	return client
}

func NewIAMClient(t *testing.T) *iam.Service {
	t.Helper()
	client, err := gcp.NewIAMClient(context.TODO())
	if err != nil {
		t.Fatalf("error creating IAM client: %v", err)
	}
	return client
}

func ResourceSupportsDeletion(resourceKind string) bool {
	switch resourceKind {
	case "BigQueryJob",
		"BinaryAuthorizationPolicy",
		"ComputeProjectMetadata",
		"DataflowFlexTemplateJob",
		"DataflowJob",
		"IAMCustomRole",
		"IAMWorkforcePool",
		"IAMWorkforcePoolProvider",
		"IAMWorkloadIdentityPool",
		"IAMWorkloadIdentityPoolProvider",
		"KMSCryptoKey",
		"KMSKeyRing",
		"LoggingLogBucket",
		"PrivateCACertificate",
		"PrivateCACertificateAuthority",
		"ResourceManagerPolicy",
		"SecretManagerSecretVersion":
		return false
	default:
		return true
	}
}
