package policy

import (
	"log"
	"fmt"
	"os"

	intersight "github.com/CiscoDevNet/intersight-go"
)

var adapterPolicyMoid, ethNetworkMoid, qosMoid, lanConnectivityMoid string

func createOrganizationRelationship(moid string) intersight.OrganizationOrganizationRelationship {
	organization := new(intersight.OrganizationOrganization)
	organization.ClassId = "mo.MoRef"
	organization.ObjectType = "organization.Organization"
	organization.Moid = &moid
	organizationRelationship := intersight.OrganizationOrganizationAsOrganizationOrganizationRelationship(organization)
	return organizationRelationship
}

func ReturnEthAdapterPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateVnicEthAdapterPolicy(config)
	ethAdapterPolicy := new(intersight.PolicyAbstractPolicy)
	ethAdapterPolicy.SetClassId("mo.MoRef")
	ethAdapterPolicy.ObjectType("vnic.EthAdapterPolicy")
	ethAdapterPolicy.SetMoid(moid)
	adapterPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(ethAdapterPolicy)
	return adapterPolicyRelationship
}

func ReturnEthNetworkPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateVnicEthNetworkPolicy(config)
	networkPolicy1 := new(intersight.PolicyAbstractPolicy)
	networkPolicy1.SetClassId("mo.MoRef")
	networkPolicy1.ObjectType("vnic.EthNetworkPolicy")
	networkPolicy1.SetMoid(moid)
	networkPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(networkPolicy1)
	return networkPolicyRelationship
}

func ReturnEthQosPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateVnicEthQosPolicy(config)
	qosPolicy1 := new(intersight.PolicyAbstractPolicy)
	qosPolicy1.SetClassId("mo.MoRef")
	qosPolicy1.ObjectType("vnic.EthQosPolicy")
	qosPolicy1.SetMoid(moid)
	qosPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(qosPolicy1)
	return qosPolicyRelationship
}

func ReturnLanPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateVnicEthLanConnectivityPolicy(config)
	lanConnectivityPolicy := new(intersight.PolicyAbstractPolicy)
	lanConnectivityPolicy.SetClassId("mo.MoRef")
	lanConnectivityPolicy.ObjectType("vnic.LanConnectivityPolicy")
	lanConnectivityPolicy.SetMoid(moid)
	lanConnectivityPolicyRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(lanConnectivityPolicy)
	return lanConnectivityPolicyRelationship
}

func ReturnEthIfPolicyAbstractPolicyRelationship(config *Config) intersight.PolicyAbstractPolicyRelationship {
	moid := CreateVnicEthIf(config)
	ethIf1 := new(intersight.PolicyAbstractPolicy)
	ethIf1.SetClassId("mo.MoRef")
	ethIf1.ObjectType("vnic.EthIf")
	ethIf1.SetMoid(moid)
	ethIfRelationship := intersight.PolicyAbstractPolicyAsPolicyAbstractPolicyRelationship(ethIf1)
	return ethIfRelationship
}

func createEthAdapterPolicyRelationship(moid string) intersight.VnicEthAdapterPolicyRelationship {
	adapterPolicy := new(intersight.VnicEthAdapterPolicy)
	adapterPolicy.ClassId = "mo.MoRef"
	adapterPolicy.ObjectType = "vnic.EthAdapterPolicy"
	adapterPolicy.Moid = &moid
	adapterRelationship := intersight.VnicEthAdapterPolicyAsVnicEthAdapterPolicyRelationship(adapterPolicy)
	return adapterRelationship
}

func createEthNetworkPolicyRelationship(moid string) intersight.VnicEthNetworkPolicyRelationship {
	networkPolicy := new(intersight.VnicEthNetworkPolicy)
	networkPolicy.ClassId = "mo.MoRef"
	networkPolicy.ObjectType = "vnic.EthNetworkPolicy"
	networkPolicy.Moid = &moid
	networkRelationship := intersight.VnicEthNetworkPolicyAsVnicEthNetworkPolicyRelationship(networkPolicy)
	return networkRelationship
}

func createEthQosPolicyRelationship(moid string) intersight.VnicEthQosPolicyRelationship {
	qosPolicy := new(intersight.VnicEthQosPolicy)
	qosPolicy.ClassId = "mo.MoRef"
	qosPolicy.ObjectType = "vnic.EthQosPolicy"
	qosPolicy.Moid = &moid
	qosRelationship := intersight.VnicEthQosPolicyAsVnicEthQosPolicyRelationship(qosPolicy)
	return qosRelationship
}

func createLanConnectivityPolicy(moid string) intersight.VnicLanConnectivityPolicyRelationship {
	lanPolicy := new(intersight.VnicLanConnectivityPolicy)
	lanPolicy.ClassId = "mo.MoRef"
	lanPolicy.ObjectType = "vnic.LanConnectivityPolicy"
	lanPolicy.Moid = &moid
	lanRelationship := intersight.VnicLanConnectivityPolicyAsVnicLanConnectivityPolicyRelationship(lanPolicy)
	return lanRelationship
}

func CreateVnicEthAdapterPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethadapterPolicy := intersight.NewVnicEthAdapterPolicyWithDefaults()
	ethadapterPolicy.PolicyAbstractPolicy.SetName("v_eth_adapter_test")
	ethadapterPolicy.SetRssSettings(true)
	ethadapterPolicy.SetUplinkFailbackTimeout(int64(5))
	
	vxlanSettingsVal := intersight.NewVnicVxlanSettingsWithDefaults()
	vxlanSettingsVal.SetEnabled(false)
	vxlanSetting := intersight.NewNullableVnicVxlanSettings(vxlanSettingsVal)
	vxlanSettings := vxlanSetting.Get()
	ethadapterPolicy.SetVxlanSettings(*vxlanSettings)
	
	nvgreSettingsVal := intersight.NewVnicNvgreSettingsWithDefaults()
	nvgreSettingsVal.SetEnabled(true)
	nvgreSetting := intersight.NewNullableVnicNvgreSettings(nvgreSettingsVal)
	nvgreSettings := nvgreSetting.Get()
	ethadapterPolicy.SetNvgreSettings(*nvgreSettings)
	
	arfsSettingsVal := intersight.NewVnicArfsSettingsWithDefaults()
	arfsSettingsVal.SetEnabled(true)
	arfsSetting := intersight.NewNullableVnicArfsSettings(arfsSettingsVal)
	arfsSettings := arfsSetting.Get()
	ethadapterPolicy.SetArfsSettings(*arfsSettings)
	
	interruptSettingsVal := intersight.NewVnicEthInterruptSettingsWithDefaults()
	interruptSettingsVal.SetCoalescingTime(int64(125))
	interruptSettingsVal.SetCoalescingType("MIN")
	interruptSettingsVal.SetCount(int64(4))
	interruptSettingsVal.SetMode("MSI")
	interruptSetting := intersight.NewNullableVnicEthInterruptSettings(interruptSettingsVal)
	interruptSettings := interruptSetting.Get()
	ethadapterPolicy.SetInterruptSettings(*interruptSettings)
	
	completionQueueSettingsVal := intersight.NewVnicCompletionQueueSettingsWithDefaults()
	completionQueueSettingsVal.SetRingSize(int64(1))
	interruptSettingsVal.SetCount(int64(4))
	completionQueueSetting := intersight.NewNullableVnicCompletionQueueSettings(completionQueueSettingsVal)
	completionQueueSettings := completionQueueSetting.Get()
	ethadapterPolicy.SetCompletionQueueSettings(*completionQueueSettings)
	
	rxQueueSettingsVal := intersight.NewVnicEthRxQueueSettingsWithDefaults()
	rxQueueSettingsVal.SetRingSize(int64(512))
	rxQueueSettingsVal.SetCount(int64(4))
	rxQueueSetting := intersight.NewNullableVnicEthRxQueueSettings(rxQueueSettingsVal)
	rxQueueSettings := rxQueueSetting.Get()
	ethadapterPolicy.SetRxQueueSettings(*rxQueueSettings)
	
	txQueueSettingsVal := intersight.NewVnicEthTxQueueSettingsWithDefaults()
	txQueueSettingsVal.SetRingSize(int64(512))
	txQueueSettingsVal.SetCount(int64(4))
	txQueueSetting := intersight.NewNullableVnicEthTxQueueSettings(txQueueSettingsVal)
	txQueueSettings := txQueueSetting.Get()
	ethadapterPolicy.SetTxQueueSettings(*txQueueSettings)
	
	tcpOffloadSettingsVal := intersight.NewVnicTcpOffloadSettingsWithDefaults()
	tcpOffloadSettingsVal.SetLargeReceive(true)
	tcpOffloadSettingsVal.SetLargeSend(true)
	tcpOffloadSettingsVal.SetRxChecksum(true)
	tcpOffloadSettingsVal.SetTxChecksum(true)
	tcpOffloadSetting := intersight.NewNullableVnicTcpOffloadSettings(tcpOffloadSettingsVal)
	tcpOffloadSettings := tcpOffloadSetting.Get()
	ethadapterPolicy.SetTcpOffloadSettings(*tcpOffloadSettings)
	
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	ethadapterPolicy.SetOrganization(organizationRelationship)

	ifMatch := ""
	ifNoneMatch := ""
	respAdapter, r, err := apiClient.VnicApi.CreateVnicEthAdapterPolicy(ctx).VnicEthAdapterPolicy(*ethadapterPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", respAdapter)
	adapterPolicyMoid = respAdapter.GetMoid()
	return adapterPolicyMoid
}

func CreateVnicEthNetworkPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethNetworkPolicy := intersight.NewVnicEthNetworkPolicyWithDefaults()
	ethNetworkPolicy.PolicyAbstractPolicy.SetName("v_eth_network_test")
	
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	ethNetworkPolicy.SetOrganization(organizationRelationship)
	
	vlanSettingsVal := intersight.NewVnicVlanSettingsWithDefaults()
	vlanSettingsVal.SetDefaultVlan(int64(1))
	vlanSettingsVal.SetMode("ACCESS")
	vlanSetting := intersight.NewNullableVnicVlanSettings(vlanSettingsVal)
	vlanSettings := vlanSetting.Get()
	ethNetworkPolicy.SetVlanSettings(*vlanSettings)
	
	ifMatch := ""
	ifNoneMatch := ""
	respNetwork, r, err := apiClient.VnicApi.CreateVnicEthNetworkPolicy(ctx).VnicEthNetworkPolicy(*ethNetworkPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Printf("HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", respNetwork)
	ethNetworkMoid = respNetwork.GetMoid()
	return ethNetworkMoid
}

func CreateVnicEthQosPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethQosPolicy := intersight.NewVnicEthQosPolicyWithDefaults()
	ethQosPolicy.PolicyAbstractPolicy.SetName("v_eth_qos_test")
	ethQosPolicy.SetMtu(int64(1500))
	ethQosPolicy.SetRateLimit(int64(0))
	ethQosPolicy.SetCos(int64(0))
	ethQosPolicy.SetTrustHostCos(false)
	
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	ethQosPolicy.SetOrganization(organizationRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	respQos, r, err := apiClient.VnicApi.CreateVnicEthQosPolicy(ctx).VnicEthQosPolicy(*ethQosPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", respQos)
	qosMoid = respQos.GetMoid()
	return qosMoid
}

func CreateVnicEthLanConnectivityPolicy(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethlanConnectivityPolicy := intersight.NewVnicLanConnectivityPolicyWithDefaults()
	ethlanConnectivityPolicy.PolicyAbstractPolicy.SetName("vnic_lan_test")
	
	org_resp,_,org_err := apiClient.OrganizationApi.GetOrganizationOrganizationList(ctx).Filter("Name eq 'default'").Execute()
	if org_err != nil {
		log.Fatalf("Error: %v\n", org_err)
	}
	orgMoid := org_resp.GetMoid()
	organizationRelationship := createOrganizationRelationship(orgMoid)
	ethlanConnectivityPolicy.SetOrganization(organizationRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	respLan, r, err := apiClient.VnicApi.CreateVnicEthLanConnectivityPolicy(ctx).VnicLanConnectivityPolicy(*ethlanConnectivityPolicy).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", respLan)
	lanConnectivityMoid = respLan.GetMoid()
	return lanConnectivityMoid
}

func CreateVnicEthIf(config *Config) string {
	var err error
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	ethIf := intersight.NewVnicEthIfWithDefaults()
	ethIf.SetName("vnic_eth_if_test")
	ethIf.SetOrder(int64(0))
	
	placementSettingVal := intersight.NewVnicPlacementSettingsWithDefaults()
	placementSettingVal.SetId("1")
	placementSettingVal.SetPciLink(int64(0))
	placementSettingVal.SetUplink(int64(0))
	placementSetting := intersight.NewNullableVnicPlacementSettings(placementSettingVal)
	placementSettings := placementSetting.Get()
	ethIf.SetPlacement(*placementSettings)
	
	cdnVal := intersight.NewVnicCdnWithDefaults()
	cdnVal.SetValue("VIC-1-eth00")
	cdnVal.SetSource("user")
	cdn := intersight.NewNullableVnicCdn(cdnVal)
	cdn1 := cdn.Get()
	ethIf.SetCdn(*cdn1)
	
	usnicSettingsVal := intersight.NewVnicUsnicSettingsWithDefaults()
	usnicSettingsVal.SetCos(int64(5))
	usnicSettingsVal.SetCount(int64(0))
	usnicSetting := intersight.NewNullableVnicUsnicSettings(usnicSettingsVal)
	usnicSettings := usnicSetting.Get()
	ethIf.SetUsnicSettings(*usnicSettings)
	
	vmqSettingsVal := intersight.NewVnicVmqSettingsWithDefaults()
	vmqSettingsVal.SetEnabled(true)
	vmqSettingsVal.SetMultiQueueSupport(false)
	vmqSettingsVal.SetNumInterrupts(int64(1))
	vmqSettingsVal.SetNumVmqs(int64(1))
	vmqSetting := intersight.NewNullableVnicVmqSettings(vmqSettingsVal)
	vmqSettings := vmqSetting.Get()
	ethIf.SetVmqSettings(*vmqSettings)
	
	ethAdapterPolicyRelationship := createEthAdapterPolicyRelationship(adapterPolicyMoid)
	ethIf.SetEthAdapterPolicy(ethAdapterPolicyRelationship)
	
	ethNetworkPolicyRelationship := createEthNetworkPolicyRelationship(ethNetworkMoid)
	ethIf.SetEthNetworkPolicy(ethNetworkPolicyRelationship)
	
	ethQosPolicyRelationship := createEthQosPolicyRelationship(qosMoid)
	ethIf.SetEthQosPolicy(ethQosPolicyRelationship)
	
	lanConnectivityPolicyRelationship := createLanConnectivityPolicy(lanConnectivityMoid)
	ethIf.SetLanConnectivityPolicy(lanConnectivityPolicyRelationship)
	
	ifMatch := ""
	ifNoneMatch := ""
	ethIfResp, r, err := apiClient.VnicApi.CreateVnicEthIf(ctx).VnicEthIf(*ethIf).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		log.Fatalf("HTTP response: %v\n", r)
	}
	fmt.Fprintf(os.Stdout, "Response: %v\n", ethIfResp)
	moid := ethIfResp.GetMoid()
	return moid
}
	