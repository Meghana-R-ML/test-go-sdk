package example

import (
	"context"
	"os"
	"log"
	"fmt"
	intersight "github.com/CiscoDevNet/intersight-go"
)

// type apiClient struct{ api intersight.ApiClient
// 	err error
//  }

type Config struct {
	ApiKey    string
	SecretKey string
	Endpoint  string
	ApiClient *intersight.APIClient
	ctx       context.Context
}

func SetInputs(apiKeyId string, apiSecretFile string, endpoint string) (context.Context, error) {
	ctx := context.Background()
	httpSigningInfo := new(intersight.HttpSignatureAuth)
	httpSigningInfo.KeyId = apiKeyId
	httpSigningInfo.PrivateKeyPath = apiSecretFile
	httpSigningInfo.SigningScheme = intersight.HttpSigningSchemeRsaSha256
	httpSigningInfo.SigningAlgorithm = intersight.HttpSigningAlgorithmRsaPKCS1v15
	httpSigningInfo.SignedHeaders = []string{intersight.HttpSignatureParameterRequestTarget,
		intersight.HttpSignatureParameterCreated,
		intersight.HttpSignatureParameterExpires,
		intersight.HttpHeaderHost,
		intersight.HttpHeaderDate,
		intersight.HttpHeaderDigest,
		"Content-Type"}
	if _, err := os.Stat(apiSecretFile); err != nil {
		err = httpSigningInfo.SetPrivateKey(apiSecretFile)
		if err != nil {
			return nil, err
		}
	} else {
		httpSigningInfo.PrivateKeyPath = apiSecretFile
	}

	ctx, err := httpSigningInfo.ContextWithValue(ctx)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error creating authentication context")
	}
// 	_, err = httpSigningInfo.GetPublicKey()
// 	if err != nil {
// 		return nil, err
// 	}
	return ctx, err
}

func getApiClient(apiKeyId string, apiSecretFile string, endpoint string) (*intersight.APIClient, error) {

	config := Config{
		ApiKey:    apiKeyId,
		SecretKey: apiSecretFile,
		Endpoint:  endpoint,
	}

	ctx, err := SetInputs(config.ApiKey, config.SecretKey, config.Endpoint)
	if err!=nil{
		log.Fatalf("Error: %v",err)
		}
	config.ctx = ctx
	cfg := intersight.NewConfiguration()
	cfg.Host = endpoint
	cfg.Debug = true
	apiClient := intersight.NewAPIClient(cfg)
	return apiClient, err
}
