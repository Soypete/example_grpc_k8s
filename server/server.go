package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	pb "github.com/Soypete/example_grpc_k8s/gengo"
	"google.golang.org/grpc"

	_ "github.com/jnewmano/grpc-json-proxy/codec"
	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("db_host")
	port     = os.Getenv("db_port")
	user     = os.Getenv("db_user")
	password = os.Getenv("db_password")
	dbname   = os.Getenv("db_name")

	db       *sql.DB
	grpcPort *string
)

// RealEstateServer is the grpc server for getting realestate data.
type RealEstateServer struct{}

func init() {
	var err error
	grpcPort = flag.String("port", "8082", "port that grpc server is exposed on")

	flag.Parse()

	// Setup and test DB
	dbStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	log.Println("setup db", dbStr)
	db, err = sql.Open("postgres", dbStr)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer db.Close()
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRealEstateServer(grpcServer, &RealEstateServer{})
	grpcServer.Serve(lis)
}

// FindHouse is grpc call for parsing housing data.
func (s *RealEstateServer) FindHouse(ctx context.Context, in *pb.Parameters) (*pb.Results, error) {
	data, err := getData(ctx, in.MaxPrice, in.Age, in.NumberOfBedrooms, in.NumberOfBathrooms, in.LotSize, db)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getData(ctx context.Context, maxPrice, age, numBedroom, numBathroom, lotSize int32, database *sql.DB) (results *pb.Results, err error) {
	year := 2019 - age
	rows, err := database.Query(`SELECT lotsizesquarefeet, taxvaluedollarcnt  FROM houses AS h
		WHERE h.taxvaluedollarcnt >= $1 AND h.yearbuilt >= $2 AND h.bedroomcnt >= $3 AND h.bathroomcnt >= $4 AND h.lotsizesquarefeet >= $5`,
		maxPrice, year, numBedroom, numBathroom, lotSize)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var house *pb.House
		if err := rows.Scan(house.Price, house.LotSize); err != nil {
			log.Fatal(err)
		}
		house.Address.Street = "123 apple rd"
		house.Address.Location.City = "Boston"
		house.Address.Location.Zip = "23456"
		house.DaysOnMarket = rand.Int31n(365)

		results.Houses = append(results.Houses, house)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}
	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send results")
	return results, nil
}
