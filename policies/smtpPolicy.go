package policy

import (
	"log"
	"fmt"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

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
	org_moid := getDefaultOrgMoid()
	organizationRelationship := getOrganizationRelationship(org_moid)
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