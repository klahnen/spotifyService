package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name   string
	Tracks []Track
}

func (a *Artist) CreateArtist(db *gorm.DB) error {
	// If ISRC already in DB skip
	dbTrack := Track{ISRC: a.Tracks[0].ISRC}
	db.Where("isrc = ?", a.Tracks[0].ISRC).First(&dbTrack)
	log.Println("dbtrack", dbTrack)
	if dbTrack.ID > 0 {
		return nil
	}

	// If already an artist with the same name, get existent Artist and just create Track
	// use FirstOrCreate in Gorm

	dbArtist := Artist{}
	db.Where("name = ?", a.Name).First(&dbArtist)

	if dbArtist.ID != 0 {
		a.ID = dbArtist.ID
		db.Save(&a)
		return nil
	}

	result := db.Create(&a)
	log.Println(a)
	return result.Error
}

func (a *Artist) GetArtist(db *gorm.DB) error {
	result := db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", a.Name)).Preload("Tracks").First(&a)
	return result.Error
}
