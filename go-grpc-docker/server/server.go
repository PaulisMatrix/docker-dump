package main

import (
	"context"
	models "go_grpc_docker/models"
	pb "go_grpc_docker/my_albums"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MyAlbumServer struct {
	pb.UnimplementedAlbumServiceServer
}

func ConvertToProto(albums models.Album) *pb.Album {
	r := &pb.Album{
		Id:     albums.ID,
		Artist: albums.Artist,
		Title:  albums.Title,
		Price:  albums.Price,
	}
	return r
}

func (server *MyAlbumServer) GetAlbums(ctx context.Context, e *emptypb.Empty) (*pb.MyAlbums, error) {
	albums := make([]*pb.Album, len(models.Albums))

	for i, a := range models.Albums {
		albums[i] = ConvertToProto(a)
	}

	return &pb.MyAlbums{Albums: albums}, nil
}

func (server *MyAlbumServer) GetAlbumByID(ctx context.Context, req *pb.GetByIDRequest) (*pb.Album, error) {
	albumID := req.GetId()
	log.Println("Got request for album ID:", albumID)

	for _, albums := range models.Albums {
		if albums.ID == albumID {
			return ConvertToProto(albums), nil
		}
	}
	return nil, nil
}

func (server *MyAlbumServer) AddAlbum(ctx context.Context, req *pb.Album) (*pb.MyAlbums, error) {
	var newAlbum models.Album

	newAlbum.ID = req.GetId()
	newAlbum.Artist = req.GetArtist()
	newAlbum.Title = req.GetTitle()
	newAlbum.Price = req.GetPrice()

	newAlbums := make([]*pb.Album, len(models.Albums)+1)
	models.Albums = append(models.Albums, newAlbum)

	for i, a := range models.Albums {
		newAlbums[i] = ConvertToProto(a)
	}

	return &pb.MyAlbums{Albums: newAlbums}, nil
}

func (server *MyAlbumServer) UpdateAlbum(ctx context.Context, req *pb.Album) (*pb.MyAlbums, error) {
	albumID := req.GetId()
	updatedAlbums := make([]*pb.Album, len(models.Albums))

	for idx := range models.Albums {
		a := &models.Albums[idx]
		if a.ID == albumID {
			a.ID = req.GetId()
			a.Artist = req.GetArtist()
			a.Title = req.GetTitle()
			a.Price = req.GetPrice()

			for i, albums := range models.Albums {
				updatedAlbums[i] = ConvertToProto(albums)
			}

			return &pb.MyAlbums{Albums: updatedAlbums}, nil
		}
	}

	return nil, nil
}

func (server *MyAlbumServer) DeleteAlbum(ctx context.Context, req *pb.Album) (*pb.MyAlbums, error) {
	albumID := req.GetId()
	updatedAlbums := make([]*pb.Album, len(models.Albums)-1)

	for idx, a := range models.Albums {
		if a.ID == albumID {
			//delete this album
			models.Albums = append(models.Albums[:idx], models.Albums[idx+1:]...)

			for i, albums := range models.Albums {
				updatedAlbums[i] = ConvertToProto(albums)
			}

			return &pb.MyAlbums{Albums: updatedAlbums}, nil
		}
	}

	return nil, nil
}

func main() {

	log.Println("grpc server starting on: 4772")
	lis, err := net.Listen("tcp", ":4772")
	if err != nil {
		log.Fatal("Failed to Listen on", err)
	}

	// in case client cancels the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcServer := grpc.NewServer()

	pb.RegisterAlbumServiceServer(grpcServer, &MyAlbumServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			grpcServer.GracefulStop()

			<-ctx.Done()
		}
	}()

}
