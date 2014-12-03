package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/stayradiated/rango/rangolib"
)

func main() {

	// setup config file
	viper.SetConfigName("config")
	viper.ReadInConfig()

	// set config defaults
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("AdminDir", "admin")

	// create router
	router := NewRouter(&RouterConfig{
		Handlers: &Handlers{
			Config:     rangolib.NewConfig("config.toml"),
			Dir:        rangolib.NewDir(),
			Page:       rangolib.NewPage(),
			ContentDir: viper.GetString("ContentDir"),
		},
		AdminDir: viper.GetString("AdminDir"),
	})

	// start http server
	log.Fatal(http.ListenAndServe(":8080", router))
}
