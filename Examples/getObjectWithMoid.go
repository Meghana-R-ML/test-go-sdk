package example

import (
	"context"
	"fmt"
	"os"
	//intersight "github.com/CiscoDevNet/intersight-go"
)

// const (
// 	apiKey    = ""
// 	SecretKey = ""
// 	endpoint  = "https://intersight.com"
// )

func GetObjectWithMoid(apiKey string, SecretKey string, endpoint string) {
	moid := "moid_of_the_Mo_document"

	apiClient, err := getApiClient(apiKey, SecretKey, endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while getting api client: %v\n", err)
	}
	apiResponse, r, err := apiClient.SmtpApi.GetSmtpPolicyByMoid(context.Background(), moid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmtpApi.GetSmtpPolicyByMoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	fmt.Fprintf(os.Stdout, "Response from `SmtpApi.GetSmtpPolicyByMoid`: %v\n", apiResponse)
}
