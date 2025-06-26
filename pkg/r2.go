package pkg

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/byterotom/infinity-play/config"
)

type R2 struct {
	client     *s3.Client
	bucketName string
}

func NewR2(env *config.Config) *R2 {

	accessKeyId := env.AccessKeyId
	accessKeySecret := env.AccessKeySecret

	cfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		awsConfig.WithRegion("auto"),
	)
	if err != nil {
		log.Fatalf("R2 credentials error %v", err)
	}

	return &R2{
		client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", env.AccountId))
		}),
		bucketName: env.BucketName,
	}
}

func (r2 *R2) Upload(key string, file io.Reader) error {
	params := &s3.PutObjectInput{
		Bucket: &r2.bucketName,
		Key:    &key,
		Body:   file,
	}
	_, err := r2.client.PutObject(context.TODO(), params)
	return err
}

func (r2 *R2) Delete(prefix string) error {

	objects := []types.ObjectIdentifier{
		{Key: aws.String(fmt.Sprintf("%s/thumbnail", prefix))},
		{Key: aws.String(fmt.Sprintf("%s/game_name.swf", prefix))},
		{Key: aws.String(fmt.Sprintf("%s/gif.gif", prefix))},
	}

	params := &s3.DeleteObjectsInput{
		Bucket: &r2.bucketName,
		Delete: &types.Delete{
			Objects: objects,
			Quiet:   aws.Bool(true),
		},
	}
	_, err := r2.client.DeleteObjects(context.TODO(), params)
	return err
}

func (r2 *R2) Get(key string) (io.ReadCloser, error) {
	params := &s3.GetObjectInput{
		Bucket: &r2.bucketName,
		Key:    &key,
	}

	obj, err := r2.client.GetObject(context.TODO(), params)
	
	return obj.Body, err
}
