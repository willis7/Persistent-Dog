package handlers

import (
	"net/http"
	"github.com/willis7/Persistent-Dog/models"
	"encoding/json"
)

var releaseStore = []models.Release{}

// Handler for HTTP Post - "/releases/"
// Creates a new Release document
func CreateRelease(w http.ResponseWriter, r *http.Request) {
	var release models.Release

	// Decode the incoming Release json
	err := json.NewDecoder(r.Body).Decode(&release)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Validate the Release entity
	err = models.Validate(release, releaseStore)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Insert release into Release Store
	releaseStore = append(releaseStore, release)
	w.WriteHeader(http.StatusCreated)
}

// Handler for HTTP Get - "/releases/"
// Returns a list of Release documents
func GetReleases(w http.ResponseWriter, r *http.Request) {
	releases, err := json.Marshal(releaseStore)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(releases)
}

// Handler for HTTP Get - "/releases/{id}"
// Returns a single Release document by id
func GetReleaseById(w http.ResponseWriter, r *http.Request) {

}

func UpdateRelease(w http.ResponseWriter, r *http.Request) {

}
