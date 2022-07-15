package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Omar-Khawaja/grpc-testing/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Omar"})
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
