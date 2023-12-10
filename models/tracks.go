package models

import (
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	ISRC            string
	SpotifyImageURI string
	Title           string
	ArtistID        uint
}

func (t *Track) GetTrack(db *gorm.DB) error {
	var err error
	result := db.First(&t, "isrc = ?", t.ISRC)
	if result.Error != nil {
		err = result.Error
	}
	return err
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
