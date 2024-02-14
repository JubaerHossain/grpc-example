package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"

    pb "github.com/JubaerHossain/grpc-example/api"
)

type Server struct{
    pb.UnimplementedExampleServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    log.Printf("Received: %v", in.GetName())
    return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    server := &Server{} // Create an instance of the Server struct
    s := grpc.NewServer()
    pb.RegisterExampleServiceServer(s, server) // Register the server
    log.Println("Server listening on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
