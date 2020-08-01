package miniostore

import "github.com/minio/minio-go/v6"

func (c *MinioStore) GetNotificationPolicy() minio.BucketNotification {
	queueArn := minio.NewArn("minio", "sqs", "", "1", "kafka")
	queueConfig := minio.NewNotificationConfig(queueArn)
	queueConfig.AddEvents(minio.ObjectCreatedAll, minio.ObjectRemovedAll, minio.ObjectAccessedAll)
	//queueConfig.AddFilterSuffix(".jpg")
	queueConfig.ID = "1"
	bucketNotification := minio.BucketNotification{}
	bucketNotification.AddQueue(queueConfig)
	bucketNotification.QueueConfigs[0].Queue = queueArn.String()
	return bucketNotification
}
