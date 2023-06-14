package main

import (
	"fmt"
	pbn "go-pbn"
	"log"
	"os"
)

func main() {
	file, err := os.Open("example-analysed.go-pbn")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting parsing")
	boards := pbn.ParsePBN(file)
	log.Println(boards)
	wFile, err := os.Create("test-board.go-pbn")
	boards.Serialize(wFile, true)
}
