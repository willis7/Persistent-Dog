package handlers

import (
	"net/http"
	"github.com/willis7/Persistent-Dog/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/willis7/Persistent-Dog/data"
	"github.com/willis7/Persistent-Dog/common"
	"gopkg.in/mgo.v2"
)

var releaseStore = []models.Release{}

// CreateRelease is handler for HTTP Post - "/releases/"
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

// GetReleases is a handler for HTTP Get - "/releases/"
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

// GetReleaseById is a handler for HTTP Get - "/releases/{id}"
// Returns a single Release document sourced by id
func GetReleaseById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := handlers.NewContext()
	defer context.Close()

	c := context.DbCollection("releases")
	repo := &data.ReleaseRepository{c}
	release, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				http.StatusInternalServerError,
			)
			return
		}
	}

	if j, err := json.Marshal(release); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

// UpdateRelease is a handler for HTTP Put - "/releases/{id}"
// Updates a single Release document sourced by id
func UpdateRelease(w http.ResponseWriter, r *http.Request) {

}
