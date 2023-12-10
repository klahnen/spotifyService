package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/klahnen/spotifyService/models"
)

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

// GetTracks ... List all tracks
// @Summary List of all tracks
// @Success 200 {array} models.Track
// @Router /tracks [get]
func (a *App) GetTracks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tracks := []models.Track{}
		models.GetTracks(a.DB, &tracks)
		respondWithJSON(w, http.StatusOK, tracks)
	}
}
