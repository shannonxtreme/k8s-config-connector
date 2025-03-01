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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccApiGatewayApiConfigIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfigIamBinding_basicGenerated(context),
			},
			{
				// Test Iam Binding update
				Config: testAccApiGatewayApiConfigIamBinding_updateGenerated(context),
			},
		},
	})
}

func TestAccApiGatewayApiConfigIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccApiGatewayApiConfigIamMember_basicGenerated(context),
			},
		},
	})
}

func TestAccApiGatewayApiConfigIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/apigateway.viewer",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProvidersOiCS,
		Steps: []resource.TestStep{
			{
				Config: testAccApiGatewayApiConfigIamPolicy_basicGenerated(context),
			},
			{
				Config: testAccApiGatewayApiConfigIamPolicy_emptyBinding(context),
			},
		},
	})
}

func testAccApiGatewayApiConfigIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "cfg%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_api_config_iam_member" "foo" {
  provider = google-beta
  api = google_api_gateway_api_config.api_cfg.api
  api_config = google_api_gateway_api_config.api_cfg.api_config_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccApiGatewayApiConfigIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "cfg%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_api_gateway_api_config_iam_policy" "foo" {
  provider = google-beta
  api = google_api_gateway_api_config.api_cfg.api
  api_config = google_api_gateway_api_config.api_cfg.api_config_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccApiGatewayApiConfigIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "cfg%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

data "google_iam_policy" "foo" {
  provider = google-beta
}

resource "google_api_gateway_api_config_iam_policy" "foo" {
  provider = google-beta
  api = google_api_gateway_api_config.api_cfg.api
  api_config = google_api_gateway_api_config.api_cfg.api_config_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccApiGatewayApiConfigIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "cfg%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_api_config_iam_binding" "foo" {
  provider = google-beta
  api = google_api_gateway_api_config.api_cfg.api
  api_config = google_api_gateway_api_config.api_cfg.api_config_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccApiGatewayApiConfigIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_api_gateway_api" "api_cfg" {
  provider = google-beta
  api_id = "tf-test-api-cfg%{random_suffix}"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api_cfg.api_id
  api_config_id = "cfg%{random_suffix}"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("test-fixtures/apigateway/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_api_config_iam_binding" "foo" {
  provider = google-beta
  api = google_api_gateway_api_config.api_cfg.api
  api_config = google_api_gateway_api_config.api_cfg.api_config_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
