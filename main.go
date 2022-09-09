package main

import (
// 	"context"
// 	"fmt"
// 	"log"
	example "test-go-sdk/Examples"
// 	intersight "github.com/CiscoDevNet/intersight-go"
)

func main() {
	keyID := "API_KEY_ID"
	keyFile := "API_KEY_PATH"
	host := "TARGET_SERVER"

	// Set up the authentication configuration struct
	// authCfg := intersight.HttpSignatureAuth{
	// 	KeyId:          keyID,
	// 	PrivateKeyPath: keyFile,

	// 	SigningScheme: intersight.HttpSigningSchemeRsaSha256,
	// 	SignedHeaders: []string{
	// 		intersight.HttpSignatureParameterRequestTarget,
	// 		"Host",
	// 		"Date",
	// 		"Digest",
	// 	},
	// 	SigningAlgorithm: intersight.HttpSigningAlgorithmRsaPKCS1v15,
	// }

	// ctx, err := authCfg.ContextWithValue(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal("Error creating authentication context")
	// }

	//     config := intersight.NewConfiguration()
	//     config.Host = host
	//     config.Debug = true

	//     client := intersight.NewAPIClient(config)

	// res, _, err := client.SnmpApi.GetSnmpPolicyList(ctx).Execute()
	// if err != nil {
	// 	log.Fatalf("Error getting SNMP policies: %v", err)
	// }

	// for _, snmpPolicy := range res.SnmpPolicyList.Results {
	// 	fmt.Printf("Snmp Policy: %s\n", *snmpPolicy.Name)
	// }

	//invoking the CreateObject to create a new instance of the bootPrecisionPolicy
	example.CreateObject(keyID, keyFile, host)

	//invoke getObjectList to get list of smtp policy
	example.GetObjectList(keyID, keyFile, host)

	//invoke getOjectListWithFilter
	example.GetObjectListWithFilter(keyID, keyFile, host)

	//invoke GetObjectWithMoid to get list of smtp policy by moid
	example.GetObjectWithMoid(keyID, keyFile, host)
}
