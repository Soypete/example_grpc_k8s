package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/soypete/example_grpc_k8s/gengo"
	"google.golang.org/grpc"
)

type RealEstateServer struct{}

func (s *RealEstateServer) FindHouse(ctx context.Context, in *pb.Parameters) (*pb.Results, error) {
	return &pb.Results{}, nil
}

func main() {
	port := flag.String("port", "8082", "port that grpc server is exposed on")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRealEstateServer(grpcServer, &RealEstateServer{})
	grpcServer.Serve(lis)
}
