package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	*minio.Client
}

type Option struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

func New(opt Option) (*Client, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(opt.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(opt.AccessKeyID, opt.SecretAccessKey, ""),
		Secure: opt.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	client := &Client{}
	client.Client = minioClient

	return client, nil
}

func (client *Client) MakeBucket(ctx context.Context, bucketName string, location string) error {

	err := client.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.Client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return fmt.Errorf("We already own %s\n", bucketName)
		} else {
			return err
		}
	}

	return nil
}

type PutOption struct {
	BucketName  string
	ObjectName  string
	Reader      io.Reader
	ObjectSize  int64
	ContentType string
}

func (client *Client) PutObject(ctx context.Context, opt PutOption) (minio.UploadInfo, error) {

	n, err := client.Client.PutObject(ctx, opt.BucketName, opt.ObjectName, opt.Reader, opt.ObjectSize, minio.PutObjectOptions{ContentType: opt.ContentType})
	if err != nil {
		return n, err
	}

	return n, nil
}

type GetOption struct {
	BucketName string
	ObjectName string
}

func (client *Client) GetObject(ctx context.Context, opt GetOption) (*minio.Object, error) {
	return client.Client.GetObject(ctx, opt.BucketName, opt.ObjectName, minio.GetObjectOptions{})
}
