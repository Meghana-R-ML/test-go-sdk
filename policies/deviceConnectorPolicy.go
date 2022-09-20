package policy

import (
	"fmt"
	"log"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)


func CreateDeviceConnectorPolicy(config *Config) {
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
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
