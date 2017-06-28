package main

import (
	"log"

	pb "github.com/myyang/grpc_exp/grpc_exp"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50511", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloMsg{Msg: "Garfield sent"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got: %v\n", r)
}
