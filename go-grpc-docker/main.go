package main

import (
	"context"
	"encoding/json"
	"go_grpc_docker/client"
	models "go_grpc_docker/models"
	pb "go_grpc_docker/my_albums"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const address string = ":4772"

type serverHandler struct {
	client  pb.AlbumServiceClient
	context context.Context
}

var (
	listAlbumRe = regexp.MustCompile(`^\/albums\/$`)
	getAlbumRe  = regexp.MustCompile(`^\/albums\/(\d+)$`)
)

func (s *serverHandler) getAlbums(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our get albums endpoint")

	w.Header().Set("Content-Type", "application/json")

	myAlbums, err := s.client.GetAlbums(s.context, &emptypb.Empty{})
	if err != nil {
		log.Fatal("Error in getting response from server", err)
	}
	json.NewEncoder(w).Encode(myAlbums.GetAlbums())

}
func (s *serverHandler) getAlbumByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our get album by id endpoint")

	//parse the url
	id := strings.TrimPrefix(r.URL.Path, "/albums/")

	w.Header().Set("Content-Type", "application/json")
	/*
		for _, a := range models.Albums {
			if a.ID == id {
				json.NewEncoder(w).Encode(a)
				return
			}
		}
	*/

	album, err := s.client.GetAlbumByID(s.context, &pb.GetByIDRequest{Id: id})
	if err != nil {
		log.Fatal("Error in getting response from server", err)
	}
	if reflect.ValueOf(album).IsZero() {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID or ID not present.")
	}
	var album_ models.Album
	album_.ID = album.GetId()
	album_.Artist = album.GetArtist()
	album_.Title = album.GetTitle()
	album_.Price = album.GetPrice()

	json.NewEncoder(w).Encode(album_)
}

func (s *serverHandler) addAlbum(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our add albums endpoint")
	var newAlbum models.Album

	//Unmarshal json to struct
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newAlbum)
	/*
		for _, a := range models.Albums {
			// if already present, don't add.
			if a.ID == newAlbum.ID {
				respondWithError(w, http.StatusBadRequest, "Album with that ID already present.")
				return
			}
		}

		models.Albums = append(models.Albums, newAlbum)
	*/

	w.Header().Set("Content-Type", "application/json")
	albums, err := s.client.AddAlbum(s.context, &pb.Album{Id: newAlbum.ID, Artist: newAlbum.Artist, Title: newAlbum.Title, Price: newAlbum.Price})
	if err != nil {
		log.Fatal("Error in getting response from server", err)
	}

	//Marshal it back
	json.NewEncoder(w).Encode(albums.GetAlbums())
}

func (s *serverHandler) deleteAlbum(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our delete a album by id endpoint")

	//parse the url
	id := strings.TrimPrefix(r.URL.Path, "/albums/")

	/*
		for index, a := range models.Albums {
			// if already present, delete that album
			if a.ID == id {
				models.Albums = append(models.Albums[:index], models.Albums[index+1:]...)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(models.Albums)
				return
			}
		}
	*/
	myAlbums, err := s.client.DeleteAlbum(s.context, &pb.Album{Id: id})
	if err != nil {
		log.Fatal("Error in getting response from server", err)
	}

	if reflect.ValueOf(myAlbums).IsZero() {
		respondWithError(w, http.StatusBadRequest, "Album with that ID not present to yeet.")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myAlbums.GetAlbums())

}

func (s *serverHandler) updateAlbum(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit our update a album by id endpoint")

	var updatedAlbum models.Album

	//Unmarshal json to struct
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updatedAlbum)

	id := strings.TrimPrefix(r.URL.Path, "/albums/")

	/*

		for index := range models.Albums {
			// if already present, delete that album
			a := &models.Albums[index]
			if a.ID == id {
				//replace in place, instead of appending at the last.
				a.ID = updatedAlbum.ID
				a.Title = updatedAlbum.Title
				a.Artist = updatedAlbum.Artist
				a.Price = updatedAlbum.Price

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(a)
				return
			}
		}
	*/
	myAlbums, err := s.client.UpdateAlbum(s.context, &pb.Album{Id: id})
	if err != nil {
		log.Fatal("Error in getting response from the server", err)
	}

	if reflect.ValueOf(myAlbums).IsZero() {
		respondWithError(w, http.StatusBadRequest, "Album with that ID not present to update.")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myAlbums.GetAlbums())

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := make(map[string]string)
	resp["message"] = message
	payload, _ := json.Marshal(resp)
	w.Write(payload)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello There!!"))
}

func (s *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && listAlbumRe.MatchString(r.URL.Path):
		s.getAlbums(w, r)
		return
	case r.Method == http.MethodGet && getAlbumRe.MatchString(r.URL.Path):
		s.getAlbumByID(w, r)
		return
	case r.Method == http.MethodPost:
		s.addAlbum(w, r)
		return
	case r.Method == http.MethodDelete:
		s.deleteAlbum(w, r)
	case r.Method == http.MethodPut:
		s.updateAlbum(w, r)
		return

	}
}

// Route declaration
func router() *http.ServeMux {

	router := http.NewServeMux()
	router.HandleFunc("/ping", healthcheck)

	grpc_client, ctx, err := client.GetNewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Not able to initialize the client", err)
	}
	router.Handle("/albums/", &serverHandler{client: grpc_client, context: ctx})
	return router
}

// Initiate web server
func main() {
	/*
		go func() {
			log.Println("grpc server starting on: 4772")
			lis, err := net.Listen("tcp", ":4772")
			if err != nil {
				log.Fatal("Failed to Listen on", err)
			}

			grpcServer := grpc.NewServer()

			pb.RegisterAlbumServiceServer(grpcServer, &server.MyAlbumServer{})

			if err := grpcServer.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()
	*/
	time.Sleep(time.Second) //wait for server to start.
	router := router()
	log.Println("Web server starting on: 3000")

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
