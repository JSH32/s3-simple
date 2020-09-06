package s3simple

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadFile : upload a file to the bucket
func (s Session) UploadFile(file bytes.Buffer, filename string, mimetype string, public bool) error {
	manager := s3manager.NewUploader(s.s3)

	options := &s3manager.UploadInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(file.Bytes()),
		ContentType: aws.String(mimetype),
	}

	if public {
		options.ACL = aws.String("public-read")
	}

	_, err := manager.Upload(options)

	if err != nil {
		return err
	}

	return nil
}

// DeleteFile : delete a file from the bucket
func (s Session) DeleteFile(filename string) error {
	manager := s3manager.NewBatchDelete(s.s3)

	objects := []s3manager.BatchDeleteObject{{
		Object: &s3.DeleteObjectInput{
			Key:    aws.String(filename),
			Bucket: aws.String(s.bucket),
		},
	}}

	if err := manager.Delete(aws.BackgroundContext(), &s3manager.DeleteObjectsIterator{Objects: objects}); err != nil {
		return err
	}

	return nil
}
