package example

import (
	"fmt"
	"log"
	"os"
	intersight "github.com/CiscoDevNet/intersight-go"
)

func GetObjectWithMoid(config *Config) {
	moid := "moid_of_the_Mo_document"
	cfg := getApiClient(config)
	apiClient := cfg.ApiClient
	ctx := cfg.ctx
	apiResponse, r, err := apiClient.SmtpApi.GetSmtpPolicyByMoid(ctx, moid).Execute()
	if err != nil {
		// 		fmt.Fprintf(os.Stderr, "Error when calling `SmtpApi.GetSmtpPolicyByMoid``: %v\n", err)
		// 		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		log.Printf("Error when calling `SmtpApi.GetSmtpPolicyByMoid``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `SmtpApi.GetSmtpPolicyByMoid`: %v\n", apiResponse)
}
