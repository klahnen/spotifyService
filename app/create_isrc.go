package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/klahnen/spotifyService/models"
)

func (a *App) CreateISRC() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createRequest models.CreateISRCRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&createRequest); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		defer r.Body.Close()

		data := a.apiSearchTrackByISCR(createRequest.ISRC)
		if len(data.Tracks.Items) == 0 {
			respondWithError(w, http.StatusBadRequest, "No data to process")
			return
		}

		var track models.Track
		var artist models.Artist

		itemIndex := a.getMostPopularItemIndex(data)
		a.populateTrackWithData(data, itemIndex, &track)
		a.populateArtistWithData(data, itemIndex, &artist)

		artist.Tracks = []models.Track{track}

		if err := artist.CreateArtist(a.DB); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusCreated, models.ISRCSuccess{Message: "Track and Artist successfully added to DB"})
	}
}

func (a *App) apiSearchTrackByISCR(iscr string) models.SearchResponse {
	var data models.SearchResponse

	url := "https://api.spotify.com/v1/search?type=track&q=isrc%3A" + iscr

	req, _ := http.NewRequest("GET", url, nil)

	token := a.conf.BearerToken
	req.Header.Add("Authorization", "Bearer "+token)

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusOK {
		log.Fatal("renew the token")
		return data
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&data)

	return data

}

func (a *App) populateTrackWithData(data models.SearchResponse, itemIndex int, t *models.Track) {
	t.Title = data.Tracks.Items[itemIndex].Name
	t.SpotifyImageURI = data.Tracks.Items[itemIndex].Album.Images[0].URL
	t.ISRC = data.Tracks.Items[itemIndex].ExternalIds.Isrc
}

func (a *App) populateArtistWithData(data models.SearchResponse, itemIndex int, artist *models.Artist) {
	artist.Name = data.Tracks.Items[itemIndex].Artists[0].Name
}

func (a *App) getMostPopularItemIndex(data models.SearchResponse) int {
	mostPopularIndex := 0

	for i := range data.Tracks.Items {
		if data.Tracks.Items[i].Popularity > data.Tracks.Items[mostPopularIndex].Popularity {
			mostPopularIndex = i
		}
	}

	return mostPopularIndex
}
