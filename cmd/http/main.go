package main

import (
	"log"
	"os"

	"github.com/mjudeikis/go-test-app/pkg/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s port", os.Args[0])
	}
	server.Start(os.Args[1])
}
