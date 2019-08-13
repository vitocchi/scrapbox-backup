package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() error {
	client, err := newScrapBoxClient()
	if err != nil {
		return err
	}
	id, err := client.getLatestBackupID()
	if err != nil {
		return err
	}
	jsonByte, err := client.getBackupJSON(id)
	if err != nil {
		return err
	}
	return storeBackupToS3(id, jsonByte)
}
func main() {
	lambda.Start(handler)
}
