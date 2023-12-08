package models

import (
	"errors"

	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name  string
	ISRCs []ISRC
}

func (a *Artist) getArtist(db *gorm.DB) error {
	return errors.New("not implemented")
}

func (a *Artist) createArtist(db *gorm.DB, name string) error {
	return errors.New("not implemented")
}

type ISRC struct {
	gorm.Model
	SpotifyImageURI string
	Title           string
	ArtistID        uint
}
