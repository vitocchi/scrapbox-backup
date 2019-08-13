package main

import (
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

const ENCRYPTED_GITHUB_TOKEN = "AQICAHgOhA6XEO8DuSQDXosJ6s6ROlrTcFIxvwlBqAn6IjPbgQFNZpbEchSj8gi8G/4yQsLQAAAAhzCBhAYJKoZIhvcNAQcGoHcwdQIBADBwBgkqhkiG9w0BBwEwHgYJYIZIAWUDBAEuMBEEDJmLhaNTyHPfV8rdZwIBEIBDET5cmHx/8Up4Z/+R5mjwW/OEj4ja4TLPwQ0/uFVLyGbPFb2aHFjwg15lRv3pCGYTwsdhYMq6wQmQHD2f7DBIxVNM/Q=="

func handler() error {
	githubToken, err := getGitHubToken()
	if err != nil {
		return error
	}
	return nil
}

func getGitHubToken() (string, error) {
	svc := kms.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	data, _ := base64.StdEncoding.DecodeString(ENCRYPTED_GITHUB_TOKEN)
	input := &kms.DecryptInput{
		CiphertextBlob: []byte(data),
	}
	result, err := svc.Decrypt(input)
	if err != nil {
		return "", err
	}
	text, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(result.Plaintext))
	return fmt.Sprintf("%s", text), nil
}

func main() {
	lambda.Start(handler)
}
