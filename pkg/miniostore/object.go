package miniostore

import (
	"github.com/minio/minio-go/v6"
	"github.com/sonasingh46/image-store-service/pkg/albums"
	"io"
	"log"
	"os"
)

func (ms *MinioStore) DeleteImage(imageName, albumName string) error {
	err := ms.Client.RemoveObject(albumName, imageName)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MinioStore) UploadImage(image albums.Image, albumName string) error {
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
		image.Name, file, fileStat.Size(),
		minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	log.Printf("%d bytes uploaded\n", n)
	return nil
}

func (ms *MinioStore) GetImage(imageName, albumName string) (*albums.StoredImage, error) {
	object, err := ms.Client.GetObject(albumName, imageName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	path := "/tmp/local-file.jpg"
	localFile, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(localFile, object); err != nil {
		return nil, err
	}
	sImage := &albums.StoredImage{
		AlbumName: albumName,
		Name:      imageName,
	}
	log.Printf("Image %s from album %s downloaded at %s", imageName, albumName, path)
	return sImage, nil
}

func (ms *MinioStore) ListImages(albumName string) (*albums.StoredImageList, error) {
	// Create a done channel to control 'ListObjects' go routine.
	doneCh := make(chan struct{})

	// Indicate to our routine to exit cleanly upon return.
	defer close(doneCh)

	isRecursive := true
	objectCh := ms.Client.ListObjects(albumName, "", isRecursive, doneCh)
	sImageList := &albums.StoredImageList{}
	for object := range objectCh {
		if object.Err != nil {
			return nil, object.Err
		}
		item := albums.StoredImage{
			AlbumName: albumName,
			Name:      object.Key,
		}
		sImageList.Items = append(sImageList.Items, item)
	}

	return sImageList, nil
}
