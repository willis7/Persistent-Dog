package routers

import (
	"github.com/gorilla/mux"
	"github.com/willis7/Persistent-Dog/handlers"
)

// SetOpsRoutes receives a pointer to a Gorilla mux router object (mux.Router) as an
// argument and returns the pointer of the mux.Router object.
func SetOpsRoutes(router *mux.Router) *mux.Router{
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")
	return router
}
