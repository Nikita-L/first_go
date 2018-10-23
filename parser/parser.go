package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

const (
	address = "centre:8080"
)

func Process() {
	Untar("./temp", strings.NewReader("data.tar"))
	csvFile, _ := os.Open("./temp/data/data.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, e := reader.Read()
		if e == io.EOF {
			break
		} else if e != nil {
			log.Fatal(e)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}
}
