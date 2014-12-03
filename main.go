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
		Page:       rangolib.NewPage(),
		ContentDir: "content",
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
