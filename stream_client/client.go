package main

import (
	"io"
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
	r, err := c.SayHello(context.Background(), &pb.HelloMsg{Msg: "Before Stream"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got: %v\n", r)

	stream, err := c.SayBiStreamHello(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("Fail to recv, err: %v\n", err)
				close(waitc)
				return
			}
			log.Printf("Got streaming result: %v\n", in)
		}
	}()
	msgs := []*pb.HelloMsg{{"MSG1"}, {"MSG2"}, {"MSG3"}, {"MSG4"}}
	for _, msg := range msgs {
		stream.Send(msg)
		log.Printf("Sent: %v\n", msg)
	}
	stream.CloseSend()
	<-waitc

	r, err = c.SayHello(context.Background(), &pb.HelloMsg{Msg: "After Stream"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got: %v\n", r)
}
