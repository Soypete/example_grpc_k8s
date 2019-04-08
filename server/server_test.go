package main

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	pb "github.com/Soypete/example_grpc_k8s/gengo"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	_ "github.com/lib/pq"
)

func TestRealEstateServer_FindHouse(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.Parameters
	}
	tests := []struct {
		name    string
		s       *RealEstateServer
		args    args
		want    *pb.Results
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RealEstateServer{}
			got, err := s.FindHouse(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("RealEstateServer.FindHouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RealEstateServer.FindHouse() = %v, want %v", got, tt.want)
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

func Test_getData(t *testing.T) {
	type args struct {
		context     context.Context
		maxPrice    int32
		age         int32
		numBedroom  int32
		numBathroom int32
		lotSize     int32
		dataBase    *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Results
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getData(tt.args.context, tt.args.maxPrice, tt.args.age, tt.args.numBedroom, tt.args.numBathroom, tt.args.lotSize, tt.args.dataBase)
			if (err != nil) != tt.wantErr {
				t.Errorf("getData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getData() = %v, want %v", got, tt.want)
			}
		})
	}
}
