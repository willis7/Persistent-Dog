package models

import (
	"errors"
)

// Release represents a release in a software context
type Release struct {
	ID          string    `json: "id"`
	Name        string `json: "name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

func Validate(release Release, releaseStore []Release) error {
	for _, u := range releaseStore {
		if u.Name == release.Name {
			return errors.New("The Name already exists")
		}
	}
	return nil
}
