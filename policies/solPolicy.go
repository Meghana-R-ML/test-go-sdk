package policy

import (
	"log"
	"fmt"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateSolPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	solPolicy := intersight.NewSolPolicyWithDefaults()
	solPolicy.PolicyAbstractPolicy.SetName("sol_test")
	solPolicy.SetEnabled(false)
	solPolicy.SetBaudRate(int32(9600))
	solPolicy.SetComPort("com1")
	solPolicy.SetSshPort(int64(1096))
	org_moid := getDefaultOrgMoid()
	organizationRelationship := getOrganizationRelationship(org_moid)
	solPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.SnmpApi.CreateSolPolicy(ctx).SolPolicy(*solPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
	moid := resp.GetMoid()
	return moid
}
