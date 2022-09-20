package main

import (
	example "test-go-sdk/Examples"
	policy "test-go-sdk/policies"
)

func main() {
	keyID := "API_KEY_ID"
	keyFile := "API_KEY_PATH"
	host := "TARGET_SERVER"


	example.ExecuteExamples(keyID, keyFile, host)
	policy.ExecutePolicies(keyID, keyFile, host)
}
