package example

import (
	"fmt"
	"os"
	"log"
	//intersight "github.com/CiscoDevNet/intersight-go"
	"context"
)

// const (
// 	apiKey    = ""
// 	SecretKey = ""
// 	endpoint  = "https://intersight.com"
// )

func GetObjectList(apiKey string, SecretKey string, endpoint string) {
	apiClient, err := getApiClient(apiKey, SecretKey, endpoint)
	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error while getting api client: %v\n", err)
		log.Printf("Error while getting api client: %v\n", err)
		return
	}

	apiResponse, r, err := apiClient.SmtpApi.GetSmtpPolicyList(context.Background()).Execute()
	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error when calling `SmtpApi.GetSmtpPolicyList``: %v\n", err)
// 		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		log.Printf("Error when calling `SmtpApi.GetSmtpPolicyList``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", r)
		return
	}
	fmt.Fprintf(os.Stdout, "Response from `SmtpApi.GetSmtpPolicyList`: %v\n", apiResponse)
}
