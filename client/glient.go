package client

import "context"

func (s *server) SearchHouses(ctx context.Context, in *pb.Parameters) (*pb.Results, error) {
	return &pb.Results{}, nil
}
