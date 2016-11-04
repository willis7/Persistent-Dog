package models

// Release represents a release in a software context
type Release struct {
	ID          string    `json: "id"`
	Name        string `json: "name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}
