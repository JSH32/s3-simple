package s3simple

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Session : new S3 session
type Session struct {
	s3     *session.Session
	bucket string
}

// New : create a new S3 session
func New(config Config) (*Session, error) {
	s, err := session.NewSession(&aws.Config{
		Region:   aws.String(config.Region),
		Endpoint: aws.String(config.Endpoint),
		Credentials: credentials.NewStaticCredentials(
			config.Credentials.Accesskey,
			config.Credentials.Secretkey,
			"",
		),
	})
	if err != nil {
		return nil, err
	}
	return &Session{
		s3:     s,
		bucket: config.Bucket,
	}, err
}
