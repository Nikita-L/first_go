package main

import (
	"../api"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port       = 8080
	dbName     = "clients"
	dbHost     = "db_centre"
	dbAdmin    = "admin"
	dbPassword = "admin#"
	dbPort     = 5432
)

type server struct{}

func (s *server) SendData(ctx context.Context, in *api.DataMessage) (*api.StatusMessage, error) {
	// database connection
	connStr := fmt.Sprintf("postgres://%s:%d/%s?sslmode=disable&user=%s&password=%s", dbHost, dbPort, dbName, dbAdmin, dbPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO people (name, email, mobile) VALUES ($1, $2, $3)", in.Name, in.Email, in.Mobile)
	if err != nil {
		log.Fatalf("DB write error: %v", err)
		return &api.StatusMessage{Ok: false}, err
	}

	return &api.StatusMessage{Ok: true}, nil
}

func main() {
	// create a listener on TCP port 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
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
