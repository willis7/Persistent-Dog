package handlers

import "github.com/willis7/Persistent-Dog/models"

type ReleaseResource struct {
	Data models.Release `json:"data"`
}
