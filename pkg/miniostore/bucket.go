package miniostore

import (
	"github.com/pkg/errors"
	"github.com/sonasingh46/image-store-service/pkg/albums"
	"log"
)

func (ms *MinioStore)CreateAlbum(album albums.Album) error  {
	location := "us-east-1"
	bucketName:=album.Name
	err := ms.Client.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := ms.Client.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket %s already exist",bucketName)
			return nil
		}
		return errors.Errorf("failed to created bucket %s:{%s}",bucketName,err.Error())
	}
	log.Printf("Successfully created album %s\n", bucketName)
	err=ms.Client.SetBucketNotification(bucketName,ms.GetNotificationPolicy())
	if err!=nil{
		return errors.Errorf("failed to set notification " +
			"on album %s:%s",album.Name,err.Error())
	}
	return nil
}

func (ms *MinioStore)DeleteAlbum(albumName string) error  {
	err := ms.Client.RemoveBucket(albumName)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MinioStore)ListAlbums() (*albums.AlbumList,error)  {
	bInfo,err := ms.Client.ListBuckets()
	if err != nil {
		return nil,err
	}
	albumList:=&albums.AlbumList{
		Items:make([]albums.Album,len(bInfo)),
	}

	for i:=0;i<len(albumList.Items);i++{
		albumList.Items[i].Name=bInfo[i].Name
	}
	return albumList,nil
}
