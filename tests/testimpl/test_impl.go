package testimpl

import (
	"context"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestMonitorPrivateLinkScoped(t *testing.T, ctx types.TestContext) {
	subscriptionId := os.Getenv("ARM_SUBSCRIPTION_ID")
	if len(subscriptionId) == 0 {
		t.Fatalf("ARM_SUBSCRIPTION_ID environment variable is not set")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %v\n", err)
	}

	resourceGroupName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	privateLinkScopeName := terraform.Output(t, ctx.TerratestTerraformOptions(), "private_link_scope_name")
	privateLinkScopeId := terraform.Output(t, ctx.TerratestTerraformOptions(), "private_link_scope_id")

	client := getPrivateLinkScopesClient(t, subscriptionId, cred)

	privateLinkScope, err := client.Get(context.TODO(), resourceGroupName, privateLinkScopeName, nil)
	if err != nil {
		log.Fatalf("Failed to get the Private Link Scope: %v", err)
	}

	t.Run("TestMonitorPrivateLinkScopedId", func(t *testing.T) {
		assert.Equal(t, strings.ToLower(privateLinkScopeId), strings.ToLower(*privateLinkScope.ID), "Private Link Scope ID does not match")
	})

	t.Run("TestOpenIngestionAndQuery", func(t *testing.T) {
		ctx.EnabledOnlyForTests(t, "with_app_insights")

		assert.Equal(t, armmonitor.AccessModeOpen, *privateLinkScope.Properties.AccessModeSettings.IngestionAccessMode, "Ingestion Access Mode is not Open")
		assert.Equal(t, armmonitor.AccessModeOpen, *privateLinkScope.Properties.AccessModeSettings.QueryAccessMode, "Query Access Mode is not Open")
	})

	t.Run("TestPrivateIngestionAndQuery", func(t *testing.T) {
		ctx.EnabledOnlyForTests(t, "with_no_resources")

		assert.Equal(t, armmonitor.AccessModePrivateOnly, *privateLinkScope.Properties.AccessModeSettings.IngestionAccessMode, "Ingestion Access Mode is not PrivateOnly")
		assert.Equal(t, armmonitor.AccessModePrivateOnly, *privateLinkScope.Properties.AccessModeSettings.QueryAccessMode, "Query Access Mode is not PrivateOnly")
	})
}

func getPrivateLinkScopesClient(t *testing.T, subscriptionId string, cred *azidentity.DefaultAzureCredential) *armmonitor.PrivateLinkScopesClient {
	client, err := armmonitor.NewPrivateLinkScopesClient(subscriptionId, cred, nil)
	if err != nil {
		t.Fatalf("Error creating PrivateLinkScopesClient: %v", err)
	}
	return client
}
