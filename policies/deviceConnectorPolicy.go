package policy

import (
	"fmt"
	"log"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)


func CreateDeviceConnectorPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid()
        organizationRelationship := getOrganizationRelationship(org_moid)
	deviceConnectorPolicy := intersight.NewDeviceconnectorPolicyWithDefaults()
	deviceConnectorPolicy.SetOrganization(organizationRelationship)
	deviceConnectorPolicy.SetLockoutEnabled(true)
	deviceConnectorPolicy.SetName("device_con1_sdk")
	deviceConnectorPolicy.SetDescription("test policy")
	resp, r, err := apiClient.DeviceconnectorApi.CreateDeviceconnectorPolicy(ctx).DeviceconnectorPolicy(*deviceConnectorPolicy).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	return resp.GetMoid()
}
