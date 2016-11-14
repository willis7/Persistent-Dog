package routers

import (
	"github.com/gorilla/mux"
	"github.com/willis7/Persistent-Dog/handlers"
)

// SetReleaseRoutes receives a pointer to a Gorilla mux router object (mux.Router) as an
//argument and returns the pointer of the mux.Router object.
func SetReleaseRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/releases", handlers.CreateRelease).Methods("POST")
	router.HandleFunc("/releases/{id}", handlers.UpdateRelease).Methods("PUT")
	router.HandleFunc("/releases", handlers.GetReleases).Methods("GET")
	router.HandleFunc("/releases/{id}", handlers.GetReleaseById).Methods("GET")
	return router
}
