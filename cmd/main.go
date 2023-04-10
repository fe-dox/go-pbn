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
	err = boards.Boards[0].Serialize(wFile, true)
	if err != nil {
		log.Fatal(err)
	}

}
