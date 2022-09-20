package policy

import (
	"log"
	"fmt"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func createOrganizationRelationship(moid string) intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organization.Moid = &moid
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func ReturnSolPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateSolPolicy(config)
	solPolicy1 := new(intersight.PolicyAbstractPolicy)
	solPolicy1.SetClassId("mo.MoRef")
	solPolicy1.ObjectType("sol.Policy")
	solPolicy1.SetMoid(moid)
	solPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(solPolicy1)
	return solPolicyRelationship
}

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
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
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
