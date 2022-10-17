package main

import (
	"context"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/NBsCoder/bookstore/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// bookstore grpc服务

type server struct {
	pb.UnimplementedBookstoreServer
	bs *bookstore
}

// ListShelves 列出书架 RPC服务
func (s *server) ListShelves(ctx context.Context, in *emptypb.Empty) (*pb.ListShelvesResponse, error) {
	sl, err := s.bs.ShelfList(ctx)
	if err == gorm.ErrEmptySlice {
		return &pb.ListShelvesResponse{}, nil
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "query failed!")
	}
	rsl := make([]*pb.Shelf, 0, len(sl))
	for _, sf := range sl {
		rsl = append(rsl, &pb.Shelf{
			Id:    sf.ID,
			Size:  sf.Size,
			Theme: sf.Theme,
		})
	}
	return &pb.ListShelvesResponse{Shelves: rsl}, nil
}

// CreateShelf 创建书架 RPC服务
func (s *server) CreateShelf(ctx context.Context, in *pb.CreateShelfRequest) (*pb.Shelf, error) {

}

// GetShelf 查询书架 RPC服务
func (s *server) GetShelf(ctx context.Context, in *pb.GetShelfRequest) (*pb.Shelf, error) {

}

// DeleteShelf 删除书架 RPC服务
func (s *server) DeleteShelf(ctx context.Context, in *pb.DeleteShelfRequest) (*emptypb.Empty, error) {

}
