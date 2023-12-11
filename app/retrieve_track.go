package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klahnen/spotifyService/models"
)

// GetTrack ... Get one Track by ISCR
// @Summary Shows metadata of a track
// @Success 200 {object} models.Track
// @Router /track/{iscr} [get]
func (a *App) GetTrackByISRC() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		iscr := vars["iscr"]

		t := models.Track{ISRC: iscr}
		if err := t.GetTrack(a.DB); err != nil {
			respondWithError(w, http.StatusNotFound, "Track not found")
			return
		}

		respondWithJSON(w, http.StatusOK, t)
	}
}
