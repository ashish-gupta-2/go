# go

## To generate protobug files from protos 
1. sudo apt install protobuf-compiler
2. sudo apt install protoc-gen-go
3. protoc -I=protos --go_out=pb protos/employee.proto
4. protoc -I=protos \
    --go_out=pb \
    --go-grpc_out=pb \
    protos/employee.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc  
5. protoc -I. -I/<path>/go-practice/go/protos/google/api \
    --proto_path=protos \
    --grpc-gateway_out=pb \
    protos/employee.proto 
    
## To run docker postgres 
1. sudo docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v /data:/var/lib/postgresql/data -d postgres

## Notes:
  --plugin=$(go env GOPATH)/bin/protoc-gen-grpc-gateway

   export PATH=$PATH:/usr/bin/protoc-gen-grpc-gateway

   export PATH="$PATH:$(go env GOPATH)/bin"


go install github.com/golang/mock/mockgen
go install  github.com/onsi/ginkgo/v2/ginkgo
go install github.com/golangci/golangci-lint/cmd/golangci-lint


References: 
https://grpc.io/docs/languages/go/quickstart/