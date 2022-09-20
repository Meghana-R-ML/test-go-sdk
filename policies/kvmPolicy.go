package policy

import (
	"fmt"
	"log"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func CreateKvmPolicy(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	org_moid := getDefaultOrgMoid()
        organizationRelationship := getOrganizationRelationship(org_moid)
	kvmPolicy := intersight.NewKvmPolicyWithDefaults()
	kvmPolicy.SetName("tf_kvm_sdk")
	kvmPolicy.SetDescription("demo kvm policy")
	kvmPolicy.SetEnabled(true)
	kvmPolicy.SetMaximumSessions(3)
	kvmPolicy.SetRemotePort(2069)
	kvmPolicy.SetEnableVideoEncryption(true)
	kvmPolicy.SetEnableLocalServerVideo(true)
	kvmPolicy.SetOrganization(organizationRelationship)
	resp, r, err := apiClient.KvmApi.CreateKvmPolicy(ctx).KvmPolicy(*kvmPolicy).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
