package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func (c Controller) Health(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Println(db.DB())
		json.NewEncoder(w).Encode("ok")
	}
}
