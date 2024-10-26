package main

import (
	"log"
	"os"
)

func main() {
	dir := os.Args[1]

	log.Printf("selected directory: %v", dir)
}
