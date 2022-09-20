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

func ReturnUserPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateIamEndPointUserPolicy(config)
	userPolicy1 := new(intersight.PolicyAbstractPolicy)
	userPolicy1.SetClassId("mo.MoRef")
	userPolicy1.ObjectType("iam.EndPointUserPolicy")
	userPolicy1.SetMoid(moid)
	userPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(userPolicy1)
	return userPolicyRelationship
}

func CreateIamEndPointUserPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	userPolicy := intersight.NewIamEndPointUserPolicyWithDefaults()
	userPolicy.PolicyAbstractPolicy.SetName("iam_end_point_user_policy_test")
	userPolicy.PolicyAbstractPolicy.SetDescription("test user policy")
	
	passwordPropertiesVal := intersight.NewIamEndPointPasswordPropertiesWithDefaults()
	passwordPropertiesVal.SetEnforceStrongPassword(true)
	passwordPropertiesVal.SetEnablePasswordExpiry(true)
	passwordPropertiesVal.SetPasswordExpiryDuration(int64(50))
	passwordPropertiesVal.SetPasswordHistory(int64(5))
	passwordPropertiesVal.SetNotificationPeriod(int64(1))
	passwordPropertiesVal.SetGracePeriod(int64(2))
	passproperties := intersight.NewNullableIamEndPointPasswordProperties(passwordPropertiesVal)
	passwordProperties := passproperties.Get()
	userPolicy.SetPasswordProperties(*passwordProperties)
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	userPolicy.SetOrganization(organizationRelationship)
	

	ifMatch := ""
	ifNoneMatch := ""
	resp, r, err := apiClient.IamApi.CreateIamEndPointUserPolicy(ctx).IamEndPointUserPolicy(*userPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", resp)
	moid := resp.GetMoid()
	return moid
}