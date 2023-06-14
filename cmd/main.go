package main

import (
	"fmt"
	"log"
	"os"
	"pbn"
)

func main() {
	file, err := os.Open("example-analysed.pbn")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting parsing")
	boards := pbn.ParsePBN(file)
	log.Println(boards)
	wFile, err := os.Create("test-board.pbn")
	boards.Serialize(wFile, true)
}
