package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sonasingh46/image-store-service/pkg/albums"
	"github.com/sonasingh46/image-store-service/pkg/decoder"
	"io"
	"log"
	"net/http"
)

func uploadImage(w http.ResponseWriter, r *http.Request){
	log.Print("Image upload request received...")
	ss,err:=NewStoreService()
	if err!=nil{
		log.Print("failed to upload image",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to upload image:"+err.Error())
		return
	}
	image:=albums.Image{}
	err = decoder.DecodeBody(r, image)
	if err != nil {
		log.Print("failed to upload image",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to upload image:"+err.Error())
		return
	}
	vars := mux.Vars(r)
	albumName := vars["albumName"]
	err=ss.UploadImage(image,albumName)
	if err != nil {
		log.Print("failed to upload image",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to upload image:"+err.Error())
		return
	}
	msg:=fmt.Sprintf("image %s uploaded successfully to album %s",image.Path,albumName)
	log.Println(msg)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w,msg)
}

func deleteImage(w http.ResponseWriter, r *http.Request){
	log.Print("Image delete request received...")
	ss,err:=NewStoreService()
	if err!=nil{
		log.Print("failed to delete image",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to delete image:"+err.Error())
		return
	}
	vars := mux.Vars(r)
	albumName := vars["albumName"]
	imageName := vars["imageName"]
	err=ss.DeleteImage(imageName,albumName)
	if err!=nil{
		log.Print("failed to delete image",err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to delete image:"+err.Error())
		return
	}
	msg:=fmt.Sprintf("image %s deleted successfully from album %s",imageName,albumName)
	log.Println(msg)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w,msg)
}
