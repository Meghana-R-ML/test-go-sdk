package policy

import (
	"context"

	intersight "github.com/CiscoDevNet/intersight-go"
)

type Config struct {
	ApiKey    string
	SecretKey string
	Endpoint  string
	ApiClient *intersight.APIClient
	ctx       context.Context
}

func ExecutePolicies(apiKey string, secret string, host string) {

	config := Config{
		ApiKey:    apiKey,
		SecretKey: secret,
		Endpoint:  host,
	}

	CreateAdapterPolicy(&config)
	CreateDeviceConnectorPolicy(&config)
	CreateLdapPolicy(&config)
	CreateIpmiPolicy(&config)
	CreateKvmPolicy(&config)
	CreateNetworkConfigPolicy(&config)
	CreateNtpPolicy(&config)
	CreateSdCardPolicy(&config)
}

func getOrganizationRelationship(moid string) intersight.OrganizationOrganizationRelationship {
        organization := new(intersight.OrganizationOrganization)
        organization.ClassId = "mo.MoRef"
        organization.ObjectType = "organization.Organization"
	organization.Moid = &moid
        organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
        return organizationRelationship
}

func getDefaultOrgMoid() string {
        cfg := getApiClient(*config)
        apiClient := cfg.ApiClient
        ctx := cfg.ctx

	org_resp, r, org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Printf("Error: %v\n", err)
                log.Printf("HTTP response: %v\n", r)
                return
	}
	org_list := org_resp.OrganizationOrganizationList.GetResults()
	if len(org_list) == 0 {
		log.Printf("Couldn't find the organization specified")
                return
	}
	org_moid := org_list[0].MoBaseMo.GetMoid()
	return org_moid
}

func getPolicyRelationship(policy string) intersight.PolicyAbstractPolicyRelationship {
	switch policy {
		case "vmedia.Policy":
			policy_moid = CreateVmediaPolicy(&config)
		case "vnic.EthNetworkPolicy":
			policy_moid = CreateVnicEthIf(&config)
	}
	policy := new(intersight.PolicyAbstractPolicy)
	policy.SetClassId("mo.MoRef")
	policy.ObjectType(policy)
	policy.SetMoid(policy_moid)
	policyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(policy)
	return policyRelationship
}
