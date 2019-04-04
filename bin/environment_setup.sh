#bin/bash
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc

export db_name="db"
export db_user="example"
export db_password="example"
export db_host="localhost"
export db_port="5432"
