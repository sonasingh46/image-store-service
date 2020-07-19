package miniostore

import (
	"github.com/minio/minio-go/v6"
	"github.com/sonasingh46/image-store-service/pkg/albums"
	"os"
)

func (ms *MinioStore)DeleteImage(imageName,albumName string) error  {
	err := ms.Client.RemoveObject(albumName,imageName)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MinioStore)UploadImage(image albums.Image,albumName string) error  {
	file, err := os.Open(image.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	n, err := ms.Client.PutObject(albumName,
		fileStat.Name(), file, fileStat.Size(),
		minio.PutObjectOptions{ContentType:"application/octet-stream"})
	if err != nil {
		return err
	}
	return nil
}
