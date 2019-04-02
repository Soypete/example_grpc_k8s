package main

import (
	"context"
	"fmt"

	pb "github.com/soypete/example_grpc_k8s/gen/go"
)

func main() {
	fmt.Println("vim-go")
}

type server struct{}

func (s *server) SearchHouses(ctx context.Context, in *pb.Parameters) (*pb.Results, error) {
	return &pb.Results{}, nil
}
