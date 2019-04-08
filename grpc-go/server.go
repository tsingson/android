package main

import (
	"net"

	log "github.com/tsingson/zaplogger"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/tsingson/android/grpc-go/proto"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Info("Say Hello --------------->")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) Order(ctx context.Context, in *pb.CoffeeRequest) (*pb.CoffeeResponse, error) {
	log.Info("Order --------------->")
	return &pb.CoffeeResponse{
		Price:   560,
		Name:    in.Name,
		Message: "Thank you for ordering " + in.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	pb.RegisterCoffeeServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
