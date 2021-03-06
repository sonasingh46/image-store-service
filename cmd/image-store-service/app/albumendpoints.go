package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sonasingh46/image-store-service/pkg/albums"
	"github.com/sonasingh46/image-store-service/pkg/decoder"
	"io"
	"log"
	"net/http"
)

func createAlbum(w http.ResponseWriter, r *http.Request) {
	log.Printf("Album create request received")
	ss, err := NewStoreService()
	if err != nil {
		log.Print("failed to create album", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to create album:"+err.Error())
		return
	}

	album := &albums.Album{}
	err = decoder.DecodeBody(r, album)
	if err != nil {
		log.Print("failed to create album", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to create album:"+err.Error())
		return
	}
	err = ss.CreateAlbum(*album)
	if err != nil {
		log.Print("failed to create album", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to create album:"+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Album created successfully")
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	log.Printf("Album delete request received")
	ss, err := NewStoreService()
	if err != nil {
		log.Print("failed to delete album", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to delete album:"+err.Error())
		return
	}

	vars := mux.Vars(r)
	albumName := vars["albumName"]
	if err != nil {
		log.Print("failed to create album", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to create album:"+err.Error())
		return
	}
	err = ss.DeleteAlbum(albumName)
	if err != nil {
		log.Print("failed to delete album", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to delete album:"+err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Album deleted successfully")
}

func listAlbums(w http.ResponseWriter, r *http.Request) {
	log.Printf("Album list request received")
	ss, err := NewStoreService()
	if err != nil {
		log.Print("failed to list albums", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to list album:"+err.Error())
		return
	}

	albumList, err := ss.ListAlbums()
	if err != nil {
		log.Print("failed to list albums", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to list album:"+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(albumList)
	if err != nil {
		log.Print("failed to list albums", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "failed to delete albums:"+err.Error())
		return
	}
}
