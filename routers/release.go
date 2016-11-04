package routers

import "github.com/gorilla/mux"

// SetReleaseRoutes receives a pointer to a Gorilla mux router object (mux.Router) as an
//argument and returns the pointer of the mux.Router object.
func SetReleaseRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/releases", controllers.CreateRelease).Methods("POST")
	router.HandleFunc("/releases/{id}", controllers.UpdateRelease).Methods("PUT")
	router.HandleFunc("/releases", controllers.GetReleases).Methods("GET")
	router.HandleFunc("/releases/{id}", controllers.GetReleaseById).Methods("GET")
	return router
}
