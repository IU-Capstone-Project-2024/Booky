package s3

import (
	"booky-back/internal/config"
	"booky-back/internal/pkg/logger"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	client *s3.Client
)

func getClient(c *config.StorageConfig) *s3.Client {
	if client == nil {
		return newClient(c)
	}

	return client
}

func newClient(c *config.StorageConfig) *s3.Client {
	cfg, err := awsConfig.LoadDefaultConfig(context.Background())
	if err != nil {
		logger.Fatalf("unable to load SDK config, %v", err)
	}

	logger.Infof("starting s3 client with endpoint: %s", c.S3.Endpoint)

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(c.S3.Endpoint)
	})

	return s3Client
}
