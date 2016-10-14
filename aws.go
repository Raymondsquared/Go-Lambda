package main

import (
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func readFromS3() {
	reg := "ap-southeast-2"
	bucket := "rr-go"
	key := "jac.json"

	svc := s3.New(session.New(&aws.Config{Region: &reg}))

	out, errSvcGetObject := svc.GetObject(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if errSvcGetObject != nil {
		log.Fatal(errSvcGetObject)
	}

	bytes, errIoutilReadAll := ioutil.ReadAll(out.Body)
	if errIoutilReadAll != nil {
		log.Fatal(errIoutilReadAll)
	}

	log.Print(string(bytes))
}
