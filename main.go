package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	server := http.Server{
		Addr: ":" + port,
	}
	log.Fatal(server.ListenAndServe())
}
