package mockS3_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"

	"github.com/kristaxox/go-examples/mockS3"
)

// MockS3 implements the GetObject function in order to mock out
// the required S3 calls made by our function
type mockS3Impl struct {
	s3iface.S3API
	InMemoryStore map[string]string
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

// GetObject will return the object contents as stored in the InMemoryStore.
func (m mockS3Impl) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return &s3.GetObjectOutput{
		Body: nopCloser{bytes.NewBufferString(m.InMemoryStore[*input.Key])},
	}, nil
}

// Since the mockS3Impl struct implements the required s3 api calls that our client
// need to make AND is of type s3iface, in our code that actually calls s3 we can
// the s3iface type instead of s3.S3, and mock out s3 call right in our tests.
// The input and output contents can be customizes to suit the amount of fileds
// your actual code uses.
func TestGetObject(t *testing.T) {
	mock := mockS3Impl{
		InMemoryStore: make(map[string]string),
	}

	mock.InMemoryStore["foo"] = "bar"

	myservice := mockS3.Myservice{
		S3Client: mock,
	}

	str, err := myservice.GetObjectAsString("foo")
	if err != nil {
		t.Fail()
	}

	if str != "bar" {
		t.Fail()
	}
}
