package server

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	pb "github.com/soypete/example_grpc_k8s/gengo"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("db_host")
	port     = os.Getenv("db_port")
	user     = os.Getenv("db_user")
	password = os.Getenv("db_password")
	dbname   = os.Getenv("db_name")
)

type RealEstateServer struct{}

func (s *RealEstateServer) FindHouse(ctx context.Context, in *pb.Parameters) (*pb.Results, error) {
	data, err := getData(in.GetMaxPrice(), in.GetAge(), in.GetNumberOfBedrooms(), in.GetNumberOfBathrooms(), in.GetLotSize(), in.GetLocation().GetCity(), in.GetLocation().GetZip())
	if err != nil {
		return nil, err
	}
	return data, nil
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

func getData(maxPrice, age, numBedroom, numBathroom, lotSize int32, city, zip string) (*pb.Results, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%, port=%, user=%, password=%, dbname=%", host, port, user, password, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var results *pb.Results
	year := 2019 - age

	rows, err := db.Query(`SELECT regionidcity, regionidzip, taxvaluedollarcnt  FROM houses
		WHERE taxvaluedollarcnt >= $1, yearbuilt > $2, bedroomcnt > $3, bathroomcnt > $4, lotsizesquarefeet > $5, regionidcity == $6, regionidzip ==$7`,
		maxPrice, year, numBedroom, numBathroom, lotSize, city, zip)
	for rows.Next() {
		var house *pb.House

		if err := rows.Scan(&house.Address.Location.City, &house.Address.Location.Zip, house.Price); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		house.Address.Street = "123 apple rd"
		house.DaysOnMarket = rand.Int31n(365)

		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
		results.Houses = append(results.Houses, house)
	}
	return results, nil
}
