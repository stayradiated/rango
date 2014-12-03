package main

import (
	"log"
	"net/http"

	"github.com/stayradiated/rango/rangolib"
)

func main() {
	router := NewRouter(&Handlers{
		Config:     rangolib.NewConfig("config.toml"),
		Dir:        rangolib.NewDir(),
		ContentDir: "content",
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
