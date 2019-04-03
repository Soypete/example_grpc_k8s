package client

import (
	"context"

	pb "github.com/soypete/example_grpc_k8s/proto"
)

type server struct{}

func (s *server) SearchHouses(ctx context.Context, in *pb.Parameters) (*pb.Results, error) {
	return &pb.Results{}, nil
}
