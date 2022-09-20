package policy

import (
	"fmt"
	"log"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)


func setServers() []string {
	servers := []string{"ntp.esl.cisco.com", "time-a-g.nist.gov", "time-b-g.nist.gov"}
	return servers
}

func CreateNtpPolicy(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid()
        organizationRelationship := getOrganizationRelationship(org_moid)
	ntpPolicy := intersight.NewNtpPolicyWithDefaults()
	ntpPolicy.SetName("tf_ntp1_sdk")
	ntpPolicy.SetEnabled(true)
	ntpPolicy.SetOrganization(organizationRelationship)
	servers := setServers()
	ntpPolicy.SetNtpServers(servers)
	resp, r, err := apiClient.NtpApi.CreateNtpPolicy(ctx).NtpPolicy(*ntpPolicy).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
