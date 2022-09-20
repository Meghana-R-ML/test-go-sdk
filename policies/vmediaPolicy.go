package policy

import (
	"log"
	"fmt"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateVmediaPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	vmediaPolicy := intersight.NewVmediaPolicyWithDefaults()
	vmediaPolicy.PolicyAbstractPolicy.SetName("vmedia_policy_test")
	vmediaPolicy.PolicyAbstractPolicy.SetDescription("vmedia policy test")
	vmediaPolicy.SetEnabled(true)
	vmediaPolicy.SetEncryption(true)
	vmediaPolicy.SetLowPowerUsb(true)
	
	org_moid := getDefaultOrgMoid()
	organizationRelationship := getOrganizationRelationship(org_moid)
	vmediaPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.VmediaApi.CreateVmediaPolicy(ctx).VmediaPolicy(*vmediaPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
	moid := resp.GetMoid()
	return moid
}