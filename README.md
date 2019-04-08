# Example of gRPC api for Kubernetes

This is an example of a simple gRPC server that is meant to be deployed to kubernetes. The purpose of this example it to display the many open source tools that have been made for creating, testing, and deploying gRPC apis. Tools include prototool, grpcurl, grpc-json-proxy, and the grpc probe. 

This example has been written in go and required GO 1.11.

## Setup Environment
To install packages and export env variables.
```bash
bash bin/environment_setup.sh
go build ./...
docker-compose up -d
```

## RUN
To generate grpc files: 

### option 1:
```bash
prototool lint
prototool config init
prototool generate
```
### option 2:
```
protoc -I=proto/ --go_out=gen/go/ proto/order.proto
protoc -I proto/ proto/real_estate.proto --go_out=plugins=grpc:gengo
```

### start server
```go 
go run server/server.go
```

### run client
```go
go run client/client.go
```

## Data
This example uses public housing data that is not included with the repo.

## Know Bugs
Docker compose is not using init script to fill table with data. 

## TODO:
- finish tests
- reflection
- json manual tests
- k8s deployment
- health probe

