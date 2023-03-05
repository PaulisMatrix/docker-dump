package client

import (
	"context"
	pb "go_grpc_docker/my_albums"
	"log"
	"time"

	"google.golang.org/grpc"
)

func GetNewClient(url string, opts ...grpc.DialOption) (pb.AlbumServiceClient, context.Context, error) {
	log.Println("Client establishing a connection to the server!!")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(ctx, url, opts...)
	if err != nil {
		return nil, nil, err
	}
	return pb.NewAlbumServiceClient(conn), ctx, nil
}
