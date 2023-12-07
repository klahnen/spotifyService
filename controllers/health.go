package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func (c Controller) Health(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		err := db.Ping()
		if err != nil {
			json.NewEncoder(w).Encode("Can't ping DB")
		} else {
			json.NewEncoder(w).Encode("ok")
		}
	}
}
