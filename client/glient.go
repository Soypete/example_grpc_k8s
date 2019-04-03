package client

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/soypete/example_grpc_k8s/gengo"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) FindHouse(ctx context.Context, in *pb.Parameters) (*pb.Results, error) {
	return &pb.Results{}, nil
}

func main() {
	serverAddr := flag.String("addr", "8082", "port to connect to grpc server.")
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.RealEstateClient(conn)

	feature, err := client.FindHouse(context.Background(), &pb.Parameters{})
	if err != nil {
		panic(err)
	}
	fmt.Println(feature)
}
