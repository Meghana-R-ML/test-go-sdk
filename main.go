package main

import (
	example "test-go-sdk/Examples"
)

func main() {
	keyID := "API_KEY_ID"
	keyFile := "API_KEY_PATH"
	host := "TARGET_SERVER"


	example.ExecuteExamples(keyID, keyFile, host)
}
