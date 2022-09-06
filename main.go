package main

import (
	"context"
	"fmt"
	"log"

	intersight "github.com/CiscoDevNet/intersight-go"
)

func main() {
	keyID := "API_KEY_ID"
	keyFile := "API_KEY_PATH"
	host := "TARGET_SERVER"


	config := intersight.NewConfiguration()
	config.Host = host

	// Uncomment this line if you want to see the Intersight API requests/responses
	config.Debug = true

	client := intersight.NewAPIClient(config)

	// Set up the authentication configuration struct
	authConfig := intersight.HttpSignatureAuth{
		KeyId:          keyID,
		PrivateKeyPath: keyFile,

		SigningScheme: intersight.HttpSigningSchemeRsaSha256,
		SignedHeaders: []string{
			intersight.HttpSignatureParameterRequestTarget, // The special (request-target) parameter expresses the HTTP request target.
			"Host",   // The Host request header specifies the domain name of the server, and optionally the TCP port number.
			"Date",   // The date and time at which the message was originated.
			"Digest", // A cryptographic digest of the request body.
		},
		SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPKCS1v15,
	}

	// Get a context that includes the authentication config
	authCtx, err := authConfig.ContextWithValue(context.Background())
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error creating authentication context")
	}

	// Execute the GetNtpPolicyList operation
	res, _, err := client.NtpApi.GetNtpPolicyList(authCtx).Execute()
	if err != nil {
		log.Fatalf("Error getting NTP policy list: %v", err)
	}

	// Loop through the NtpPolicy list and print the names
	for _, ntpPolicy := range res.NtpPolicyList.Results {
		fmt.Printf("Ntp Policy: %s\n", *ntpPolicy.Name)
	}
}
