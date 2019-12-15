package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"net/url"
	"os"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		DisableRestProtocolURICleaning: aws.Bool(true),
		Endpoint:                       aws.String(os.Args[1]),
	}))
	//		LogLevel:                       aws.LogLevel(aws.LogDebugWithHTTPBody),

	u, _ := url.Parse(os.Args[2])
	bucket := "//" + u.Host
	path := u.Path
	path = path[1:]

	fmt.Printf("bucket: %s path: %s\n", bucket, path)
	svc := s3.New(sess)
	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(path),
	}
	resp, err := svc.ListObjects(params)

	if err != nil {
		fmt.Printf("sv.ListObjects err: %s\n", err)
	}
	if resp.Delimiter != nil {
		fmt.Printf("Delimiter: %s\n", *resp.Delimiter)
	}
	if resp.EncodingType != nil {
		fmt.Printf("EncodingType: %v\n", *resp.EncodingType)
	}
	if resp.IsTruncated != nil {
		fmt.Printf("IsTruncated: %v\n", *resp.IsTruncated)
	}
	if resp.Marker != nil {
		fmt.Printf("Marker: %v\n", *resp.Marker)
	}
	if resp.MaxKeys != nil {
		fmt.Printf("MaxKeys: %v\n", *resp.MaxKeys)
	}
	if resp.Name != nil {
		fmt.Printf("Name: %v\n", *resp.Name)
	}
	if resp.NextMarker != nil {
		fmt.Printf("NextMarker: %v\n", *resp.NextMarker)
	}
	if resp.Prefix != nil {
		fmt.Printf("Prefix: %v\n", *resp.Prefix)
	}

	for _, key := range resp.Contents {
		fmt.Printf("Key: %s \n", *key.Key)
		key := *key.Key
		bucket = bucket
		fmt.Printf("bucket: %s Key: %s \n", bucket, key)
		resp, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})

		if err != nil {
			fmt.Printf("svc.GetObject failed err: %s\n", err)
			break
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ioutil.ReadAll failed err: %s\n", err)
			break
		}

		Data := bytes.NewBuffer(b)
		fmt.Printf("data:\n%s\n", Data)
	}
}
