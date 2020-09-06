package s3simple

// Config : basic config for your S3 bucket
type Config struct {
	Region      string
	Endpoint    string
	Bucket      string
	Credentials Credentials
}

// Credentials : credentials to access your bucket
type Credentials struct {
	Accesskey string
	Secretkey string
}
