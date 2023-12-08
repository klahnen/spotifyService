package models_test

import (
	"log"
	"os"
	"testing"

	"github.com/klahnen/spotifyService/driver"
	"github.com/klahnen/spotifyService/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

func getTestDB() *gorm.DB {
	log.Print("creating DB...")
	return driver.ConnectDB("test.db")
}
func deleteTestDB() {
	currentPath, _ := os.Getwd()
	err := os.Remove(currentPath + "/test.db")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("DB destroyed")
}

func TestMain(m *testing.M) {
	db = getTestDB()

	code := m.Run()

	deleteTestDB()

	os.Exit(code)
}
func TestArtistCreation(t *testing.T) {
	db.AutoMigrate(&models.Artist{})

	createdArtist := models.Artist{Name: "Artist 1"}
	db.Create(&createdArtist)

	artist := models.Artist{}
	db.First(&artist, "id = ?", createdArtist.ID)

	assert.Equal(t, artist.Name, "Artist 1")
}

func TestISRCCreation(t *testing.T) {

	db.AutoMigrate(&models.Artist{}, &models.ISRC{})

	artist := models.Artist{Name: "some artist name"}
	db.Create(&artist)

	isrc := models.ISRC{Title: "my title", SpotifyImageURI: "some-uri", ArtistID: artist.ID}
	db.Create(&isrc)
	db.First(&artist, "id = ?", artist.ID)

	db.Model(&artist).Association("ISRCs").Find(&artist.ISRCs)

	assert.Equal(t, isrc.ArtistID, artist.ID)
	assert.Equal(t, 1, len(artist.ISRCs))
	assert.Equal(t, "some-uri", artist.ISRCs[0].SpotifyImageURI)
}
