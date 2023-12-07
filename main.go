package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klahnen/spotifyService/config"
	"github.com/klahnen/spotifyService/controllers"
	"github.com/klahnen/spotifyService/driver"
	"github.com/klahnen/spotifyService/models"
)

func main() {
	conf := config.GetConfig()

	db := driver.ConnectDB()
	db.AutoMigrate(&models.Artist{}, &models.ISRC{})

	controller := controllers.Controller{}

	r := mux.NewRouter()
	r.HandleFunc("/health", controller.Health(db)).Methods("GET")
	r.HandleFunc("/redirect-URI", controller.Callback(conf))

	// creation of the track, get spotify image uri, title, artist name list
	//r.HandleFunc("/isrc", controller.CreateISRC(db).Methods("POST"))

	// get metadata by ISRC (single result)
	// r.HandleFunc("/isrc/{id}", controller.RetrieveISRC().Methods("GET"))
	// get metadata by Artist (multiple results)
	http.Handle("/", r)

	log.Printf("Server up on port %v....\n", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.Port), r))
}
