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

func ReturnVmediaPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateVmediaPolicy(config)
	vmediaPolicy1 := new(intersight.PolicyAbstractPolicy)
	vmediaPolicy1.SetClassId("mo.MoRef")
	vmediaPolicy1.ObjectType("vmedia.Policy")
	vmediaPolicy1.SetMoid(moid)
	vmediaPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(vmediaPolicy1)
	return vmediaPolicyRelationship
}

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
	
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
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