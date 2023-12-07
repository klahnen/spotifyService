package models

import "gorm.io/gorm"

type Artist struct {
	gorm.Model
	ISRCs []ISRC
}

type ISRC struct {
	gorm.Model
	SpotifyImageURI string
	Title           string
	ArtistID        uint
}
