package main

import (
	"log"
	"net/http"

	"github.com/stayradiated/rango/server"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", server.Start()))
}
