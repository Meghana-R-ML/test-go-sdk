package example

import (
	"context"
	//	intersight "github.com/CiscoDevNet/intersight-go"
)

type Config struct {
	ApiKey    string
	SecretKey string
	Endpoint  string
	ApiClient *intersight.APIClient
	ctx       context.Context
}

func executeExamples(apiKey string, secret string, host string) {

	config := Config{
		ApiKey:    apiKey,
		SecretKey: secret,
		Endpoint:  host,
	}

	CreateObject(&config)
	GetObjectList(&config)
	GetObjectListWithFilter(&config)
	GetObjectWithMoid(&config)
}
