package policy

import (
	"context"

	intersight "github.com/CiscoDevNet/intersight-go"
)

type Config struct {
	ApiKey    string
	SecretKey string
	Endpoint  string
	ApiClient *intersight.APIClient
	ctx       context.Context
}

func ExecutePolicies(apiKey string, secret string, host string) {

	config := Config{
		ApiKey:    apiKey,
		SecretKey: secret,
		Endpoint:  host,
	}

	CreateAdapterPolicy(&config)
	CreateDeviceConnectorPolicy(&config)
	CreateLdapPolicy(&config)
	CreateIpmiPolicy(&config)
	CreateKvmPolicy(&config)
	CreateNetworkConfigPolicy(&config)
	CreateNtpPolicy(&config)
	CreateSdCardPolicy(&config)
}
