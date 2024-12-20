package main

import (
	"lat.sh/backend"
	"log"
	"net/http"
	"time"
)

const Address = ":8080"

const ReadTimeout = 10 * time.Second

const WriteTimeout = 10 * time.Second

func main() {
	http.HandleFunc("/health", backend.Health)
	log.Fatal(http.ListenAndServe(Address, nil))
}
