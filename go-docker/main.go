package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our get albums endpoint")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)

}
func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our get album by id endpoint")

	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	for _, a := range albums {
		if a.ID == id {
			json.NewEncoder(w).Encode(a)
			return
		}
	}
	respondWithError(w, http.StatusBadRequest, "Invalid user ID")
}

func addAlbum(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our add albums endpoint")
	var newAlbum album

	//Unmarshal json to struct
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newAlbum)

	for _, a := range albums {
		// if already present, don't add.
		if a.ID == newAlbum.ID {
			respondWithError(w, http.StatusBadRequest, "Album with that ID already present.")
			return
		}
	}

	albums = append(albums, newAlbum)
	w.Header().Set("Content-Type", "application/json")

	//Marshal it back
	json.NewEncoder(w).Encode(albums)
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our delete a album by id endpoint")

	vars := mux.Vars(r)
	id := vars["id"]

	for index, a := range albums {
		// if already present, delete that album
		if a.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(albums)
			return
		}
	}
	respondWithError(w, http.StatusBadRequest, "Album with that ID not present to yeet.")
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our update a album by id endpoint")

	var updatedAlbum album

	//Unmarshal json to struct
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updatedAlbum)

	vars := mux.Vars(r)
	id := vars["id"]

	for index, a := range albums {
		// if already present, delete that album
		if a.ID == id {
			// first remove that album
			albums = append(albums[:index], albums[index+1:]...)

			//append the album with updated details
			albums = append(albums, updatedAlbum)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(albums)
			return
		}
	}
	respondWithError(w, http.StatusBadRequest, "Album with that ID not present to update.")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello There!!"))
}

// Route declaration
func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", healthcheck)
	router.HandleFunc("/albums", getAlbums).Methods("GET")
	router.HandleFunc("/albums/{id}", getAlbumByID).Methods("GET")
	router.HandleFunc("/albums", addAlbum).Methods("POST")
	router.HandleFunc("/albums/{id}", deleteAlbum).Methods("DELETE")
	router.HandleFunc("/albums/{id}", updateAlbum).Methods("PATCH")
	return router
}

// Initiate web server
func main() {
	router := router()
	log.Println("Listening on :3000")
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
