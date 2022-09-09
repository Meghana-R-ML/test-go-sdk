package example

import (
	"context"
	"fmt"
	"os"
	"log"
	//intersight "github.com/CiscoDevNet/intersight-go"
)

// const (
// 	apiKey    = ""
// 	SecretKey = ""
// 	endpoint  = "https://intersight.com"
// )

func GetObjectListWithFilter(apiKey string, SecretKey string, endpoint string) {
	filter := "CreateTime gt 2021-08-29T21:58:33Z"
	orderby := "CreateTime"
	top := int32(10)
	skip := int32(100)
	select_ := "CreateTime,ModTime"
	expand := "Biosunits"
	apply := "groupby((Model),aggregate(AvailableMemory with min as MinAvailableMemory))"
	count := false
	inlinecount := "allpages"
	at := "VersionType eq 'Configured'"
	tags := "tags_example"

	apiClient, err := getApiClient(apiKey, SecretKey, endpoint)
	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error while getting api client: %v\n", err)
		log.Printf("Error while getting api client: %v\n", err)
		return
	}
	apiResponse, r, err := apiClient.ComputeApi.GetComputeRackUnitList(context.Background()).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error when calling `ComputeApi.GetComputeRckUnitList``: %v\n", err)
// 		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		log.Printf("Error when calling `ComputeApi.GetComputeRckUnitList``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `ComputeApi.GetComputeRckUnitList`: %v\n", apiResponse)
}
