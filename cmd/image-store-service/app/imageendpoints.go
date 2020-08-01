package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sonasingh46/image-store-service/pkg/albums"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Image Upload Endpoint Hit...")
	ss, err := NewStoreService()
	if err != nil {
		log.Print("failed to upload image", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to upload image:"+err.Error())
		return
	}

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		msg := fmt.Sprintf("failed to upload image as error in retrieving file: %s", err.Error())
		log.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, msg)
		return
	}
	defer file.Close()

	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)
	//
	//fmt.Println("File Name",handler.Filename)

	// Create a temporary file within our temp directory that follows
	// a particular naming pattern without conflict.
	path := "/tmp"
	tempFile, err := ioutil.TempFile(path, handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	log.Print("Successfully Uploaded File to server\n")
	log.Println("Uploading to minio store...")
	vars := mux.Vars(r)
	albumName := vars["albumName"]
	image := &albums.Image{
		Path: tempFile.Name(),
		Name: handler.Filename,
	}
	fmt.Println("Image Path", image.Path)
	err = ss.UploadImage(*image, albumName)
	if err != nil {
		log.Print("failed to upload image", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to upload image:"+err.Error())
		return
	}
	msg := fmt.Sprintf("image %s uploaded successfully to album %s", image.Path, albumName)
	log.Println(msg)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, msg)
}

func deleteImage(w http.ResponseWriter, r *http.Request) {
	log.Print("Image delete request received...")
	ss, err := NewStoreService()
	if err != nil {
		log.Print("failed to delete image", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to delete image:"+err.Error())
		return
	}
	vars := mux.Vars(r)
	albumName := vars["albumName"]
	imageName := vars["imageName"]
	err = ss.DeleteImage(imageName, albumName)
	if err != nil {
		log.Print("failed to delete image", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to delete image:"+err.Error())
		return
	}
	msg := fmt.Sprintf("image %s deleted successfully from album %s", imageName, albumName)
	log.Println(msg)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, msg)
}

func listImages(w http.ResponseWriter, r *http.Request) {
	log.Print("Image list request received...")
	ss, err := NewStoreService()
	if err != nil {
		log.Print("failed to list images", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to list images:"+err.Error())
		return
	}
	vars := mux.Vars(r)
	albumName := vars["albumName"]
	imageList, err := ss.ListImages(albumName)
	if err != nil {
		log.Print("failed to list images", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to list images:"+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(imageList)
	if err != nil {
		log.Print("failed to list images", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to list images:"+err.Error())
		return
	}

}

func getImage(w http.ResponseWriter, r *http.Request) {
	log.Print("Image get/download request received...")
	ss, err := NewStoreService()
	if err != nil {
		log.Print("failed to get image", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to get image:"+err.Error())
		return
	}
	vars := mux.Vars(r)
	albumName := vars["albumName"]
	imageName := vars["imageName"]
	image, err := ss.GetImage(imageName, albumName)
	if err != nil {
		log.Print("failed to get images", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to get images:"+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(image)
	if err != nil {
		log.Print("failed to get images", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to get images:"+err.Error())
		return
	}
}
