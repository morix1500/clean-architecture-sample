package infrastructure

import (
	"net"
	pb "github.com/morix1500/clean-architecture-sample/proto"
	"github.com/morix1500/clean-architecture-sample/interfaces/controllers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50021"
)

type Blog struct{}

func (Blog) Insert(ctx context.Context, in *pb.InsertRequest) (*pb.InsertResponse, error) {
	blogController := controllers.NewBlogController(NewSqlHandler())
	res, err := blogController.Insert(in)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (Blog) Select(ctx context.Context, in *pb.SelectRequest) (*pb.SelectResponse, error) {
	blogController := controllers.NewBlogController(NewSqlHandler())
	res, err := blogController.Select(in)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (b Blog) Run() {
	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, b)
	reflection.Register(s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
