package policy

import (
	"log"
	"fmt"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func createSnmpTrap() *intersight.SnmpTrap {
	snmpTrap := intersight.NewSnmpTrap("snmp.Trap", "snmp.Trap")
	snmpTrap.SetDestination("10.10.10.1")
	snmpTrap.SetEnabled(false)
	snmpTrap.SetPort(int64(660))
	snmpTrap.SetType("Trap")
	snmpTrap.SetUser("demouser")
	snmpTrap.SetVersion("V3")
	return snmpTrap
}

func createSnmpUser() *intersight.SnmpUser {
	snmpUser := intersight.NewSnmpUser("snmp.User","snmp.User")
	var auth_pass,priv_pass string
	snmpUser.SetName("demouser")
	snmpUser.SetPrivacyType("AES")
	snmpUser.SetAuthPassword(auth_pass)
	snmpUser.SetPrivacyPassword(priv_pass)
	snmpUser.SetSecurityLevel("AuthPriv")
	snmpUser.SetAuthType("SHA")
	return snmpUser
}

func createOrganizationRelationship(moid string) intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organization.Moid = &moid
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func ReturnSnmpPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateSnmpPolicy(config)
	snmpPolicy1 := new(intersight.PolicyAbstractPolicy)
	snmpPolicy1.SetClassId("mo.MoRef")
	snmpPolicy1.ObjectType("snmp.Policy")
	snmpPolicy1.SetMoid(moid)
	snmpPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(snmpPolicy1)
	return snmpPolicyRelationship
}

func CreateSnmpPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	snmpPolicy := intersight.NewSnmpPolicyWithDefaults()
	snmpPolicy.PolicyAbstractPolicy.SetName("snmp_policy_test")
	snmpPolicy.PolicyAbstractPolicy.SetDescription("sample testing snmp policy")
	snmpPolicy.SetEnabled(true)
	snmpPolicy.SetSnmpPort(int64(1983))
	snmpPolicy.SetAccessCommunityString("dummy123")
	snmpPolicy.SetCommunityAccess("Disabled")
	snmpPolicy.SetTrapCommunity("TrapCommunity")
	snmpPolicy.SetSysContact("aanimish")
	snmpPolicy.SetSysLocation("Karnataka")
	snmpPolicy.SetEngineId("vvb")
	snmpTrap1 := createSnmpTrap()
	snmpTraps := []intersight.SnmpTrap{*snmpTrap1}
	snmpPolicy.SetSnmpTraps(snmpTraps)
	snmpUser1 := createSnmpUser()
	snmpUsers := []intersight.SnmpUser{*snmpUser1}
	snmpPolicy.SetSnmpUsers(snmpUsers)

	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	snmpPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.SnmpApi.CreateSnmpPolicy(ctx).SnmpPolicy(*snmpPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
	moid := resp.GetMoid()
	return moid
}