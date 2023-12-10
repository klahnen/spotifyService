package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klahnen/spotifyService/config"
	"github.com/klahnen/spotifyService/driver"
	"github.com/klahnen/spotifyService/models"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	conf   config.Config
}

func (a *App) Initialize() {
	a.conf = config.GetConfig()
	a.DB = driver.ConnectDB(a.conf.DBName)
	a.DB.AutoMigrate(&models.Artist{}, &models.Track{})

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/health", a.Health()).Methods("GET")
	a.Router.HandleFunc("/login", a.Login()).Methods("GET")
	a.Router.HandleFunc("/callback", a.Callback()).Methods("GET")
	a.Router.HandleFunc("/isrc", a.CreateISRC()).Methods("POST") // ToDo change endpoint to track
	a.Router.HandleFunc("/track/{iscr}", a.GetTrackByISRC()).Methods("GET")
	a.Router.HandleFunc("/tracks", a.GetTracks()).Methods("GET")
	a.Router.HandleFunc("/artist/{name}", a.GetArtist()).Methods("GET")
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
