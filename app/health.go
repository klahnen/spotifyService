package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *App) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Println(a.DB.DB())
		json.NewEncoder(w).Encode("ok")
	}
}
