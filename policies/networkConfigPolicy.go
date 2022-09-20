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

func CreateNetworkConfigPolicy(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	organization := setOrganization()
	networkConfigPolicy := intersight.NewNetworkConfigPolicyWithDefaults()
	networkConfigPolicy.SetName("tf_network_config1_sdk")
	networkConfigPolicy.SetDescription("test policy")
	networkConfigPolicy.SetEnableDynamicDns(false)
	networkConfigPolicy.SetPreferredIpv4dnsServer("::")
	networkConfigPolicy.SetEnableIpv6(true)
	networkConfigPolicy.SetEnableIpv6dnsFromDhcp(false)
	networkConfigPolicy.SetPreferredIpv4dnsServer("10.10.10.1")
	networkConfigPolicy.SetAlternateIpv4dnsServer("10.10.10.1")
	networkConfigPolicy.SetAlternateIpv6dnsServer("::")
	networkConfigPolicy.SetEnableIpv4dnsFromDhcp(false)
	networkConfigPolicy.SetOrganization(organization)
	resp, r, err := apiClient.CreateNetworkconfigPolicy(ctx).NetworkconfigPolicy(*networkConfigPolicy).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
