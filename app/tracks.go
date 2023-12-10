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

		var err error

		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid ISCR")
			return
		}

		t := models.Track{ISRC: iscr}
		if err := t.GetTrack(a.DB); err != nil {
			respondWithError(w, http.StatusNotFound, "Track not found")
			return
		}

		respondWithJSON(w, http.StatusOK, t)

	}
}
