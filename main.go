package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joshbrusa/go-auth/api"
)

func main() {
	port := flag.Int("port", 80, "The port the server will listen on.")
	flag.Parse()

	server := api.NewServer(*port)
	fmt.Println("Server listening on port:", *port)
	log.Fatal(server.Start())
}