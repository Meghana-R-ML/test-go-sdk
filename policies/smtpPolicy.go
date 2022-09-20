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

func ReturnSmtpPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateSmtpPolicy(config)
	smtpPolicy1 := new(intersight.PolicyAbstractPolicy)
	smtpPolicy1.SetClassId("mo.MoRef")
	smtpPolicy1.ObjectType("smtp.Policy")
	smtpPolicy1.SetMoid(moid)
	smtpPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(smtpPolicy1)
	return smtpPolicyRelationship
}

func CreateSmtpPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	smtpPolicy := intersight.NewSmtpPolicyWithDefaults()
	smtpPolicy.PolicyAbstractPolicy.SetName("smtp_policy_test")
	smtpPolicy.PolicyAbstractPolicy.SetDescription("sample testing smtp policy")
	smtpPolicy.SetEnabled(false)
	smtpPolicy.SetSmtpPort(int64(32))
	smtpPolicy.SetMinSeverity("critical")
	smtpPolicy.SetSmtpServer("10.10.10.1")
	smtpPolicy.SetSenderEmail("IMCSQAAutomation@cisco.com")
	smtpRecipients := []string{"aw@cisco.com",
    "cy@cisco.com",
    "dz@cisco.com"}
	smtpPolicy.SetSmtpRecipients(smtpRecipients)
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	smtpPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.SmtpApi.CreateSmtpPolicy(ctx).SmtpPolicy(*smtpPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
	moid := resp.GetMoid()
	return moid
}