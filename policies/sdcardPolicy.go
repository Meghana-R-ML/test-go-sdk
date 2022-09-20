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

func createVirtualDrives() intersight.SdcardVirtualDrive {
	virtualDrives := intersight.NewSdcardVirtualDriveWithDefaults()
	virtualDrives.SetEnable(true)
	virtualDrives.SetObjectType("sdcard.OperatingSystem")
	return virtualDrives
}

func setPartitions() intersight.SdcardPartition {
	partitions := intersight.NewSdcardPartitionWithDefaults()
	partitions.SetType("OS")
	partitions.SetObjectType("sdcard.Partition")
	virtualDrives := createVirtualDrives()
	partitions.SetVirtualDrives(*virtualDrives)
	return partitions
}

func CreateSdCardPolicy(config *Config) {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	organization := setOrganization()
	partitions := setPartitions()
	sdCardPolicy := intersight.NewSdcardPolicyWithDefaults()
	sdCardPolicy.SetName("tf_sdcard_sdk")
	sdCardPolicy.SetDescription("test policy")
	sdCardPolicy.SetOrganization(organization)
	sdCardPolicy.SetPartitions(*partitions)
	resp, r, err := apiClient.SdcardApi.CreateSdcardPolicy(ctx).SdcardPolicy(*sdCardPolicy).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
}
