package app

import (
	"github.com/sonasingh46/image-store-service/pkg/albums"
	"github.com/sonasingh46/image-store-service/pkg/miniostore"
)

func NewStoreService() (albums.StoreService,error){
	mc:=miniostore.NewMinioStoreConfig()
	ms,err:=miniostore.NewMinioStore(mc)
	if err!=nil{
		return nil,err
	}
	return ms,nil
}