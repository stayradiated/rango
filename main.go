package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	viper.SetDefault("AssetsDir", "static/assets")

	// make sure content dir exists
	contentDir := viper.GetString("ContentDir")
	_, err := os.Stat(contentDir)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(contentDir, 0755)
	}

	// create router
	router := NewRouter(&RouterConfig{
		Handlers: &Handlers{
			Config:     rangolib.NewConfig("config.toml"),
			Dir:        rangolib.NewDir(),
			Page:       rangolib.NewPage(),
			ContentDir: contentDir,
		},
		AdminDir:  viper.GetString("AdminDir"),
		AssetsDir: viper.GetString("AssetsDir"),
	})

	// start http server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
