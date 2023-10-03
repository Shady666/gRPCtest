package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "puppy/gRPCtest/ping_service_proto" // Путь к сгенерированному пакету

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "ok",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingServiceServer(s, &server{})
	fmt.Println("Server started at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
