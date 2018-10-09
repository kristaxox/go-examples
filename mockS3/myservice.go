package mockS3

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// GetObjectAsString returns the content of an s3 object as a string
// it is important here that the s3Client is of type s3iface.S3API
func GetObjectAsString(s3Client s3iface.S3API, key string) (string, error) {
	out, err := s3Client.GetObject(&s3.GetObjectInput{
		Key: aws.String(key),
	})
	if err != nil {
		return "", err
	}

	str, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
