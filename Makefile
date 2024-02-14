.PHONY: all build-server run-server build-client run-client clean

PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(GOPATH)/bin/protoc-gen-go-grpc

all: build-server build-client

$(PROTOC_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go

$(PROTOC_GEN_GO_GRPC):
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

build-server: $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) api/example.proto
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative  api/example.proto
	go build -o server/server server/*.go

run-server: build-server
	./server/server

build-client: $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) api/example.proto
	protoc --go_out=client --go-grpc_out=client api/example.proto
	go build -o client/client client/*.go

run-client: build-client
	./client/client

clean:
	rm -f server/server client/client server/*.pb.go client/*.pb.go
