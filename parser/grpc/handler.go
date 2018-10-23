package grpc

import (
	"golang.org/x/net/context"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

func (s *Server) SendData(ctx context.Context, in *DataMessage) (*StatusMessage, error) {
	log.Printf("Receive message %s", in.type_)
	return &StatusMessage{True}, nil
}
