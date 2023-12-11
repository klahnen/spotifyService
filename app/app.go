package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klahnen/spotifyService/config"
	"github.com/klahnen/spotifyService/spotify"

	_ "github.com/klahnen/spotifyService/docs"
	"github.com/klahnen/spotifyService/driver"
	"github.com/klahnen/spotifyService/models"
	httpSwagger "github.com/swaggo/http-swagger"
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
	a.Router.HandleFunc("/track", a.CreateTrack(spotify.GetClient())).Methods("POST")
	a.Router.HandleFunc("/track/{iscr}", a.GetTrackByISRC()).Methods("GET")
	a.Router.HandleFunc("/tracks", a.GetTracks()).Methods("GET")
	a.Router.HandleFunc("/artist/{name}", a.GetArtist()).Methods("GET")
	a.Router.PathPrefix("/docs/swagger.json").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))
	a.Router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://127.0.0.1:8000/docs/swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
