package main

import (
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func decryptByKMS(encrypted string) (string, error) {
	data, _ := base64.StdEncoding.DecodeString(encrypted)
	svc := kms.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
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
