package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/joshbrusa/go-auth/routes"
)

func main() {
	port := flag.Int("port", 80, "The port the server will listen on.")
	flag.Parse()
	addr := ":" + strconv.Itoa(*port)

	routes.Router()

	fmt.Println("Server listening on port:", *port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

