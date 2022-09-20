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

func ReturnSyslogPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateSyslogPolicy(config)
	syslogPolicy1 := new(intersight.PolicyAbstractPolicy)
	syslogPolicy1.SetClassId("mo.MoRef")
	syslogPolicy1.ObjectType("syslog.Policy")
	syslogPolicy1.SetMoid(moid)
	syslogPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(syslogPolicy1)
	return syslogPolicyRelationship
}

func CreateSyslogPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	syslogPolicy := intersight.NewSyslogPolicyWithDefaults()
	syslogPolicy.PolicyAbstractPolicy.SetName("syslog_test")
	syslogPolicy.PolicyAbstractPolicy.SetDescription("demo syslog policy")
	
	remoteClient1 := intersight.NewSyslogRemoteClientBaseWithDefaults()
	remoteClient1.SetEnabled(true)
	remoteClient1.SetHostname("10.10.10.10")
	remoteClient1.SetPort(int64(514))
	remoteClient1.SetProtocol("tcp")
	remoteClient1.SetMinSeverity("emergency")
	
	
	remoteClient2 := intersight.NewSyslogRemoteClientBaseWithDefaults()
	remoteClient2.SetEnabled(true)
	remoteClient2.SetHostname("2001:0db8:0a0b:12f0:0000:0000:0000:0004")
	remoteClient2.SetPort(int64(64000))
	remoteClient2.SetProtocol("udp")
	remoteClient2.SetMinSeverity("emergency")
	remoteClients := []intersight.SyslogRemoteClientBase{*remoteClient1, *remoteClient2}
	syslogPolicy.SetRemoteClients(remoteClients)
	
	localClient1 := intersight.NewSyslogLocalClientBaseWithDefaults()
	localClient1.SetMinSeverity("emergency")
	localClients := []intersight.SyslogLocalClientBase{*localClient1}
	syslogPolicy.SetLocalClients(localClients)
	
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	syslogPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.SyslogApi.CreateSyslogPolicy(ctx).SyslogPolicy(*syslogPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
	moid := resp.GetMoid()
	return moid
}