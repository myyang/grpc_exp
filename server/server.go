package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/myyang/grpc_exp/grpc_exp"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloMsg) (*pb.HelloReply, error) {
	log.Printf("SayHello from client: %v\n", in)
	return &pb.HelloReply{Result: fmt.Sprintf("Result: %s (from server)", in.Msg)}, nil
}

func (s *server) SayBiStreamHello(stream pb.HelloService_SayBiStreamHelloServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("SayBiStreamHello from client: %v\n", in)
		stream.Send(&pb.HelloReply{Result: fmt.Sprintf("Stream Result: %s (from server)", in.Msg)})

	}
}

func main() {
	lis, err := net.Listen("tcp", ":50511")
	if err != nil {
		log.Fatal(err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("Serve at port :50511...")
	s.Serve(lis)
}
