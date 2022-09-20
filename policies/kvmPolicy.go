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

func CreateKvmPolicy(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	organization := setOrganization()
	kvmPolicy := intersight.NewKvmPolicyWithDefaults()
	kvmPolicy.SetName("tf_kvm_sdk")
	kvmPolicy.SetDescription("demo kvm policy")
	kvmPolicy.SetEnabled(true)
	kvmPolicy.SetMaximumSessions(3)
	kvmPolicy.SetRemotePort(2069)
	kvmPolicy.SetEnableVideoEncryption(true)
	kvmPolicy.SetEnableLocalServerVideo(true)
	kvmPolicy.SetOrganization(organization)
	resp, r, err := apiClient.KvmApi.CreateKvmPolicy(ctx).KvmPolicy(*kvmPolicy).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
