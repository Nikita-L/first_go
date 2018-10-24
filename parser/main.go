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
	"regexp"
)

const (
	link = "centre:8080"
)

func main() {
	// start connection
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(link, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewDataFlowClient(conn)

	// untar data
	err = archiver.Tar.Open("data.tar", "temp")
	if err != nil {
		log.Fatalf("can't extract: %s", err)
	}

	// get data
	csvFile, _ := os.Open("./temp/data/data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// mobile number filter
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	// read, send data
	reader.Read() // skip header
	for {
		line, e := reader.Read()
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal(e)
		}
		cleanMobile := reg.ReplaceAllString(line[3], "")

		if len(cleanMobile) > 10 && len(cleanMobile) < 3 {
			log.Printf("Bad mobile number for user %s, %s", line[1], cleanMobile)
			continue
		}

		response, err := c.SendData(context.Background(), &api.DataMessage{Name: line[1], Email: line[2], Mobile: cleanMobile})

		if err != nil {
			log.Fatalf("Error when calling SendData: %s", err)
		}
		log.Printf("%t", response.Ok)
	}
	PrintMemUsage()
}
