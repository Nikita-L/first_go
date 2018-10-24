package main

import (
	"../api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct{}

func (s *server) SendData(ctx context.Context, in *api.DataMessage) (*api.StatusMessage, error) {
	//log.Printf(in.Name)
	return &api.StatusMessage{Ok: true}, nil
}

func main() {
	// create a listener on TCP port 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := grpc.NewServer()
	api.RegisterDataFlowServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
