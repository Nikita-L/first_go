package main

import (
	"google.golang.org/grpc"
	"log"

	"../api"
	"bufio"
	"context"
	"encoding/csv"
	"github.com/mholt/archiver"
	"io"
	"os"
)

const (
	link = "centre:8080"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(link, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewDataFlowClient(conn)

	err = archiver.Tar.Open("data.tar", "temp")
	if err != nil {
		log.Fatalf("can't extract: %s", err)
	}

	csvFile, _ := os.Open("./temp/data/data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for i := 1; ; i++ { // skip header
		line, e := reader.Read()
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal(e)
		}

		response, err := c.SendData(context.Background(), &api.DataMessage{Name: line[1], Email: line[2], Mobile: line[3]})

		if err != nil {
			log.Fatalf("Error when calling SendData: %s", err)
		}
		log.Printf("%t: %d", response.Ok, i)
	}
	PrintMemUsage()
}
