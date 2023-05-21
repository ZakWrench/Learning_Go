package main

import (
	"log"
	"os"
)

func main() {
	// empty file Creation
	// Create() function Using
	myfile, es := os.Create("text.txt")
	if es != nil {
		log.Fatal(es)
	}
	log.Println(myfile)
	myfile.Close()
}
