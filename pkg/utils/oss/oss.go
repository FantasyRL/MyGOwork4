package aliyunoss

import (
	"bibi/pkg/conf"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func OSSBucketCreate() (*oss.Bucket, error) {
	endpoint := conf.OSS.EndPoint
	accessKeyId := conf.OSS.AccessKeyId
	accessKeySecret := conf.OSS.AccessKeySecret
	bucketName := conf.OSS.BucketName
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
