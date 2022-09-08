package example

import (
	"context"
	"fmt"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

// const (
// 	apiKeyId      = ""
// 	apiSecretFile = ""
// 	endpoint      = ""
// )

// type Config struct {
// 	ApiKey    string
// 	SecretKey string
// 	Endpoint  string
// 	ApiClient *intersight.APIClient
// 	ctx       context.Context
// }

func createBootLocalCdd() *intersight.BootDeviceBase {
	bootLocalCdd := intersight.NewBootDeviceBase("boot.LocalCdd", "boot.LocalCdd")
	return bootLocalCdd
}

func createBootLocalDisk() *intersight.BootDeviceBase {
	bootLocalDisk := intersight.NewBootDeviceBase("boot.LocalDisk", "boot.LocalDisk")
	return bootLocalDisk
}

func createBootSdcard() *intersight.BootDeviceBase {
	bootSdcard := intersight.NewBootDeviceBase("boot.SdCard", "boot.SdCard")
	bootSdcard.SetName("boot_sdcard_test")
	bootSdcard.SetEnabled(true)
	return bootSdcard
}

func createBootIscsi() *intersight.BootDeviceBase {
	bootIscsi := intersight.NewBootDeviceBase("boot.Iscsi", "boot.Iscsi")
	bootIscsi.SetName("boot_iscsi_test")
	bootIscsi.SetEnabled(true)
	return bootIscsi
}

func createBootVirtualMedia() *intersight.BootDeviceBase {
	bootVirtualMedia := intersight.NewBootDeviceBase("boot.VirtualMedia", "boot.VirtualMedia")
	bootVirtualMedia.SetName("boot_virtual_media_test")
	bootVirtualMedia.SetEnabled(true)
	return bootVirtualMedia
}

func createOrganization() intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func createOrganizationWithMoid(moid string) intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organization.Moid = &moid
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func CreateObject(apiKeyId string, apiSecret string, endpoint string) {
	// config := Config{
	// 	ApiKey:    apiKeyId,
	// 	SecretKey: apiSecretFile,
	// 	Endpoint:  endpoint,
	// }
	var err error
	// config.ctx, err = setInputs(apiKeyId, apiSecretFile, endpoint)
	// if err != nil {
	// 	fmt.Println("Unable to get APIClient: ", err)
	// 	return
	// }
	// config.ApiClient,err= getApiClient(config.ctx)
	// if err != nil {
	// 	fmt.Println("Error", err)
	// 	return
	// }

	ApiClient, err := getApiClient(apiKeyId, apiSecret, endpoint)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	bootLocalCdd := createBootLocalCdd()
	bootLocalDisk := createBootLocalDisk()
	organization := createOrganization()
	bootDevices := []intersight.BootDeviceBase{*bootLocalDisk, *bootLocalCdd}
	bootPrecisionPolicy := intersight.NewBootPrecisionPolicyWithDefaults()
	bootPrecisionPolicy.SetName("sample_boot_policy1")
	bootPrecisionPolicy.SetDescription("sample boot precision policy")
	bootPrecisionPolicy.SetBootDevices(bootDevices)
	bootPrecisionPolicy.SetOrganization(organization)
	apiClient := ApiClient
	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.BootApi.CreateBootPrecisionPolicy(context.Background()).BootPrecisionPolicy(*bootPrecisionPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)

	//Update
	id := resp.GetMoid()
	getapiResponse, r, err := apiClient.BootApi.GetBootPrecisionPolicyByMoid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error -> GetBootPrecisionPolicyByMoid: %v\n", err)
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", r)
	}
	objMoid := getapiResponse.GetMoid()
	organizationMoid := getapiResponse.GetOrganization().MoMoRef.GetMoid()
	bootSdcard := createBootSdcard()
	bootIscsi := createBootIscsi()
	organization1 := createOrganizationWithMoid(organizationMoid)
	bootDevices1 := []intersight.BootDeviceBase{*bootSdcard, *bootIscsi}
	updatebootPrecisionPolicy := intersight.NewBootPrecisionPolicyWithDefaults()
	updatebootPrecisionPolicy.PolicyAbstractPolicy.SetName("updated_boot_precision_policy_for_go_test")
	updatebootPrecisionPolicy.PolicyAbstractPolicy.SetDescription("updated description of boot precision policy for testing go example")
	updatebootPrecisionPolicy.SetBootDevices(bootDevices1)
	updatebootPrecisionPolicy.SetOrganization(organization1)
	updateResp, r, err := apiClient.BootApi.UpdateBootPrecisionPolicy(context.Background(), objMoid).BootPrecisionPolicy(*updatebootPrecisionPolicy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error -> UpdateBootPrecisionPolicy: %v\n", err)
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response : %v\n", updateResp)

	//Patch
	bootVirtualMedia := createBootVirtualMedia()
	bootDevices2 := []intersight.BootDeviceBase{*bootVirtualMedia}
	patchbootPrecisionPolicy := intersight.NewBootPrecisionPolicyWithDefaults()
	patchbootPrecisionPolicy.PolicyAbstractPolicy.SetName("updated_boot_precision_policy_using_patch_go_test")
	patchbootPrecisionPolicy.PolicyAbstractPolicy.SetDescription("update the description of boot precision policy with patch for go test")
	patchbootPrecisionPolicy.SetBootDevices(bootDevices2)
	patchbootPrecisionPolicy.SetOrganization(organization1)
	patchResp, r, err := apiClient.BootApi.PatchBootPrecisionPolicy(context.Background(), objMoid).BootPrecisionPolicy(*patchbootPrecisionPolicy).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error -> PatchBootPrecisionPolicy: %v\n", err)
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response : %v\n", patchResp)

	//Delete
	fullResp, err := apiClient.BootApi.DeleteBootPrecisionPolicy(context.Background(), objMoid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error -> DeleteBootPrecisionPolicy: %v\n", err)
		fmt.Fprintf(os.Stderr, "HTTP response: %v\n", fullResp)
	}

}
