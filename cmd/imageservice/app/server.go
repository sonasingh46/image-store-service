package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)




func StartImageService()  {
	r := mux.NewRouter()
	// Endpoint for checking service liveness
	r.HandleFunc("/healthz", HealthCheckHandler)

	//// Endpoint for listing all the albums.
	//r.HandleFunc("/albums", listAllAlbums).Methods("GET")

	// Endpoint for creating an album
	r.HandleFunc("/albums", createAlbum).Methods("POST")
	// Endpoint for deleting abum
	r.HandleFunc("/albums/{albumName}", deleteAlbum).Methods("DELETE")
	//
	//// Endpoint to list all images in an album
	//r.HandleFunc("/albums/{albumName}/images", listImage).Methods("GET")
	// Endpoint to create an image in an album
	r.HandleFunc("/albums/{albumName}/images", uploadImage).Methods("POST")
	r.HandleFunc("/albums/{albumName}/images/{imageName}", deleteImage).Methods("DELETE")

	log.Print("Image service started...")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
