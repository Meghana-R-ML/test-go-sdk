package policy

import (
	"fmt"
	"log"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)


func CreateIpmiPolicy(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid()
        organizationRelationship := getOrganizationRelationship(org_moid)
	ipmiPolicy := intersight.NewIpmioverlanPolicyWithDefaults()
	ipmiPolicy.SetName("tf_ipmi_sdk")
	ipmiPolicy.SetDescription("demo ipmi policy")
	ipmiPolicy.SetEnabled(true)
	ipmiPolicy.SetPrivilege("admin")
	ipmiPolicy.SetOrganization(organizationRelationship)
	resp, r, err := apiClient.IpmioverlanApi.CreateIpmioverlanPolicy(ctx).IpmioverlanPolicy(*ipmiPolicy).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
