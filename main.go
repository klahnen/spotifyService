package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klahnen/spotifyService/config"
	"github.com/klahnen/spotifyService/controllers"
	"github.com/klahnen/spotifyService/driver"
)

func main() {
	conf := config.GetConfig()

	db := driver.ConnectDB()
	controller := controllers.Controller{}

	r := mux.NewRouter()
	r.HandleFunc("/health", controller.Health(db)).Methods("GET")
	r.HandleFunc("/redirect-URI", controller.Callback(conf))
	// creation of the track
	http.Handle("/", r)

	log.Printf("Server up on port %v....\n", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.Port), r))
}
