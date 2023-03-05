package server

import (
	"context"
	models "go_grpc_docker/models"
	pb "go_grpc_docker/my_albums"
	"log"
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

func (server *MyAlbumServer) GetAlbums(ctx context.Context, req *pb.EmptyRequest) (*pb.MyAlbums, error) {
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
