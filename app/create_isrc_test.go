package app_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/klahnen/spotifyService/app"
	"github.com/klahnen/spotifyService/mocks"
	"github.com/klahnen/spotifyService/spotify"
	"github.com/stretchr/testify/assert"
)

var a app.App

func deleteTestDB() {
	err := os.Remove("/tmp/dev.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("DB destroyed")
}

func TestMain(m *testing.M) {
	a.Initialize()

	code := m.Run()

	deleteTestDB()

	os.Exit(code)
}

func executeRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

type MockMusicService struct{}

func (m *MockMusicService) ApiSearchTrackByISCR(isrc string) spotify.SearchResponse {
	var response spotify.SearchResponse
	json.Unmarshal([]byte(mocks.GetTrack_GBAYE0601477()), &response)
	return response
}

func TestCreateISRCEndpoint(t *testing.T) {

	mockClient := MockMusicService{}

	handler := a.CreateISRC(&mockClient)

	var jsonStr = []byte(`{"ISRC":"GBAYE0601477"}`)
	req, _ := http.NewRequest("POST", "/isrc", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req, handler)
	assert.Equal(t, http.StatusCreated, response.Result().StatusCode)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	assert.Equal(t, "Track and Artist successfully added to DB", m["message"])
}
