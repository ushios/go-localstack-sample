package sample

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const S3Endpoint = "http://localstack:4572"

func S3PutObject(bucketName, key string, rs io.ReadSeeker) *s3.PutObjectOutput {
	s := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(endpoints.UsWest2RegionID),
		Endpoint:         aws.String(S3Endpoint),
	}))

	c := s3.New(s, &aws.Config{})

	p := s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		ACL:    aws.String("public-read"),
		Body:   rs,
	}

	r, err := c.PutObject(&p)
	if err != nil {
		panic(err)
	}

	return r
}
