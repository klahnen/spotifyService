package models

import (
	"gorm.io/gorm"
)

// Track model info
// @Description Track information
// @Description ISRC, image and artistID
type Track struct {
	gorm.Model
	ISRC            string
	SpotifyImageURI string
	Title           string
	ArtistID        uint
}

func (t *Track) GetTrack(db *gorm.DB) error {
	result := db.First(&t, "isrc = ?", t.ISRC)
	return result.Error
}

func GetTracks(db *gorm.DB, tracks *[]Track) error {
	db.Find(tracks)
	return nil
}

type ISRCSuccess struct {
	Message string `json:"message"`
}

type CreateISRCRequest struct {
	ISRC string
}
