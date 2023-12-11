package app

import (
	"net/http"

	"github.com/klahnen/spotifyService/models"
)

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
