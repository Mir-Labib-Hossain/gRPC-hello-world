package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Mir-Labib-Hossain/grpc-hello-world/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect go gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloWorldServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloWorldRequest{})
	if err != nil {
		log.Fatalf("error calling function SayHello() : %v", err)
	}
	log.Printf("Response from gRPC server: %s", r.GetMessage())
}
