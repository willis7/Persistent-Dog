package routers

import "github.com/gorilla/mux"

// InitRoutes initializes all routes for the RESTful API
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the Release entity
	router = SetReleaseRoutes(router)
	router = SetOpsRoutes(router)
	return router
}
