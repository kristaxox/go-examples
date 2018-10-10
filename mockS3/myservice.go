package mockS3

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// Myservice is an example serice that uses an s3Client
type Myservice struct {
	S3Client s3iface.S3API
}

// GetObjectAsString returns the content of an s3 object as a string
// it is important here that the s3Client is of type s3iface.S3API
func (m *Myservice) GetObjectAsString(key string) (string, error) {
	out, err := m.S3Client.GetObject(&s3.GetObjectInput{
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
