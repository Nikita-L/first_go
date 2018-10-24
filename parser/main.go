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

//func processLine(line []string, conn *grpc.ClientConn, wg sync.WaitGroup) {
//	log.Print("processLine")
//
//	c := api.NewDataFlowClient(conn)
//
//	response, err := c.SendData(context.Background(), &api.DataMessage{Name: line[1], Email: line[2], Mobile: line[3]})
//	log.Print("we have response")
//
//	if err != nil {
//		log.Fatalf("Error when calling SendData: %s", err)
//	}
//	log.Printf("%t", response.Ok)
//
//	wg.Done()
//}
//
//
//func process(reader *csv.Reader, conn *grpc.ClientConn) {
//	var wg sync.WaitGroup
//
//	for i := 1 ; ; i++ {  // i = 1 - skip header
//		line, e := reader.Read()
//		if e == io.EOF {
//			return
//		} else if e != nil {
//			log.Fatal(e)
//		}
//
//		wg.Add(1)
//		go processLine(line, conn, wg)
//	}
//
//
//	wg.Wait()
//}

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

	//process(reader, conn)

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
