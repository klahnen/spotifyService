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
	"github.com/stretchr/testify/assert"
)

var a app.App

func deleteTestDB() {
	currentPath, _ := os.Getwd()
	err := os.Remove(currentPath + "/dev.db")
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

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestCreateISRCEndpoint1(t *testing.T) {
	var jsonStr = []byte(`{"ISRC":"GBAYE0601477"}`)
	req, _ := http.NewRequest("POST", "/isrc", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	assert.Equal(t, http.StatusCreated, response.Result().StatusCode)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	assert.Equal(t, "Track and Artist successfully added to DB", m["message"])
}
