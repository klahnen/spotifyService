package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klahnen/spotifyService/models"
)

// GetArtist ... Get an artist by name
// @Summary Get an artist
// @Param name path string true "Name"
// @Success 200 {object} models.Artist
// @Failure 404 {object} object
// @Router /artist/{name} [get]
func (a *App) GetArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		name := vars["name"]

		log.Println(name)

		artist := models.Artist{Name: name}
		if err := artist.GetArtist(a.DB); err != nil {
			log.Println(err)
			respondWithError(w, http.StatusNotFound, "Artist not found")
			return
		}

		respondWithJSON(w, http.StatusOK, artist)
	}
}
