package policy

import (
	"fmt"
	"log"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func setOrganization() intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func setBaseProperties() intersight.IamLdapBaseProperties {
	ldapBaseProperties := intersight.NewIamLdapBasePropertiesWithDefaults()
	ldapBaseProperties.SetAttribute("CiscoAvPair")
	ldapBaseProperties.SetBaseDn("DC=QATCSLABTPI02,DC=cisco,DC=com")
	ldapBaseProperties.SetBindDn("CN=administrator,CN=Users,DC=QATCSLABTPI02,DC=cisco,DC=com")
	ldapBaseProperties.SetBindMethod("Anonymous")
	ldapBaseProperties.SetDomain("QATCSLABTPI02.cisco.com")
	ldapBaseProperties.SetEnableEncryption(true)
	ldapBaseProperties.SetEnableGroupAuthorization(true)
	ldapBaseProperties.SetFilter("sAMAccountName")
	ldapBaseProperties.SetGroupAttribute("memberOf")
	ldapBaseProperties.SetNestedGroupSearchDepth(128)
	ldapBaseProperties.SetTimeout(180)
	return ldapBaseProperties
}

func setDnsProperties() intersight.IamLdapDnsParameters {
	ldapDnsProperties := intersight.NewIamLdapDnsParametersWithDefaults()
	ldapDnsProperties.SetSource("Extracted")
	ldapDnsProperties.SetSearchForest("xyz")
	ldapDnsProperties.SetDomain("abc")
	return ldapDnsProperties
}

func CreateLdapPolicy(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	organization := setOrganization()
	baseProperties := setBaseProperties()
	dnsProperties := setDnsProperties()
	ldapPolicy := intersight.NewIamLdapPolicyWithDefaults()
	ldapPolicy.SetName("tf_ldap1_sdk")
	ldapPolicy.SetDescription("test policy")
	ldapPolicy.SetEnabled(true)
	ldapPolicy.SetEnableDns(true)
	ldapPolicy.SetUserSearchPrecedence("LocalUserDb")
	ldapPolicy.SetOrganization(organization)
	ldapPolicy.SetBaseProperties(baseProperties)
	ldapPolicy.SetDnsPaarameters(dnsProperties)
	resp, r, err := apiClient.IamApi.CreateIamLdapPolicy(ctx).IamLdapPolicy(*ldapPolicy).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
