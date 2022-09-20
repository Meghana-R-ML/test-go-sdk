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

func ReturnPolicyAbstractPolicyRelationshipStoragePolicy(config *Config) intersight.PolicyAbstractPolicyRelationship, string {
	moid := CreateStorageStoragePolicy(config)
	storagePolicy := new(intersight.PolicyAbstractPolicy)
	storagePolicy.SetClassId("mo.MoRef")
	storagePolicy.ObjectType("storage.DriveGroup")
	storagePolicy.SetMoid(moid)
	storagePolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(storagePolicy)
	return storagePolicyRelationship,moid
}

func CreateStorageStoragePolicy(config *Config) string{
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	storageStoragePolicy := intersight.NewStorageStoragePolicyWithDefaults()
	storageStoragePolicy.PolicyAbstractPolicy.SetName("storage_policy_test")
	storageStoragePolicy.PolicyAbstractPolicy.SetDescription( "storage policy test")
	storageStoragePolicy.SetUseJbodForVdCreation(true)
	storageStoragePolicy.SetUnusedDisksState("UnconfiguredGood")
	storageStoragePolicy.SetGlobalHotSpares("3")
	
	m2VirtualDrive := intersight.NewStorageM2VirtualDriveConfigWithDefaults()
	m2VirtualDrive.SetEnable(false)
	m2VirtualDrive1 := intersight.NewNullableStorageM2VirtualDriveConfig(m2VirtualDrive)
	m2Virtual1 := m2VirtualDrive1.Get()
	storageStoragePolicy.SetM2VirtualDrive(*m2Virtual1)
	
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	storageStoragePolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.StorageApi.CreateStorageStoragePolicy(ctx).StorageStoragePolicy(*storageStoragePolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
	
	moid := resp.GetMoid()
	return moid
}