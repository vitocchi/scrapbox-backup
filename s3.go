package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func storeBackupToS3(id int, data []byte) error {
	var BUCKET = os.Getenv("BUCKET")
	key := fmt.Sprintf("%d.json", id)
	svc := s3.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	_, err := svc.PutObject(&s3.PutObjectInput{
		Body:                 bytes.NewReader(data),
		Bucket:               aws.String(BUCKET),
		Key:                  aws.String(key),
		ACL:                  aws.String("private"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}
