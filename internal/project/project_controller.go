package project

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetProjects(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, status, err := getProjects(db)

		respondJSON(w, response, status, err)
	}
}

// Update project
// Post project
// Delete project

func respondJSON(w http.ResponseWriter, response interface{}, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
