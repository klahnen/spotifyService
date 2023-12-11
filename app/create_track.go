package app

import (
	"encoding/json"
	"net/http"

	"github.com/klahnen/spotifyService/models"
	"github.com/klahnen/spotifyService/spotify"
)

type MusicService interface {
	ApiSearchTrackByISCR(iscr string) spotify.SearchResponse
}

// CreateTrack ... Creates an ISRC
// @Summary From an ISRC executes a search in Spotify to pull data from artists and tracks
// @Param name body object true "ISRC"
// @Success 200 {object} models.CreateISRCRequest
// @Router /track [post]
func (a *App) CreateTrack(client MusicService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createRequest models.CreateISRCRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&createRequest); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		defer r.Body.Close()

		data := client.ApiSearchTrackByISCR(createRequest.ISRC)
		if len(data.Tracks.Items) == 0 {
			respondWithError(w, http.StatusBadRequest, "No data to process")
			return
		}

		artist := a.getArtistFromData(data)

		if err := artist.CreateArtist(a.DB); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusCreated, models.ISRCSuccess{Message: "Track and Artist successfully added to DB"})
	}
}

func (a *App) getArtistFromData(data spotify.SearchResponse) models.Artist {
	itemIndex := a.getMostPopularItemIndex(data)

	return models.Artist{
		Name: data.Tracks.Items[itemIndex].Artists[0].Name,
		Tracks: []models.Track{{
			Title:           data.Tracks.Items[itemIndex].Name,
			SpotifyImageURI: data.Tracks.Items[itemIndex].Album.Images[0].URL,
			ISRC:            data.Tracks.Items[itemIndex].ExternalIds.Isrc,
		}},
	}

}

func (a *App) getMostPopularItemIndex(data spotify.SearchResponse) int {
	mostPopularIndex := 0

	for i := range data.Tracks.Items {
		if data.Tracks.Items[i].Popularity > data.Tracks.Items[mostPopularIndex].Popularity {
			mostPopularIndex = i
		}
	}

	return mostPopularIndex
}
