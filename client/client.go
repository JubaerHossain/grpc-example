package main

import (
    "context"
    "log"

    "google.golang.org/grpc"

    pb "github.com/JubaerHossain/grpc-example/client"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewExampleServiceClient(conn)

    resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Response: %s", resp.Message)
}
