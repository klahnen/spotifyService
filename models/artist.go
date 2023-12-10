package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Artist struct {
	gorm.Model
	Name   string
	Tracks []Track
}

func (a *Artist) CreateArtist(db *gorm.DB) error {
	first := int64(0)

	result := db.Where("isrc = ?", a.Tracks[0].ISRC).FirstOrCreate(&a.Tracks[0])
	if result.RowsAffected == first {
		return nil
	}

	db.Where("name = ?", a.Name).First(&a)
	db.Save(&a)

	return result.Error
}

func (a *Artist) GetArtist(db *gorm.DB) error {
	result := db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", a.Name)).Preload("Tracks").First(&a)
	return result.Error
}
