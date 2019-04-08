package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/soypete/example_grpc_k8s/gengo"
)

func Test_server_FindHouse(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.Parameters
	}
	tests := []struct {
		name    string
		s       *server
		args    args
		want    *pb.Results
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{}
			got, err := s.FindHouse(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.FindHouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.FindHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
