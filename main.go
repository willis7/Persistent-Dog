package main

import (
	"github.com/willis7/relenv/common"
	"github.com/willis7/relenv/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
)

func main() {
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig[Server],
		Handler: n,
	}
	log.Println("Listening...")
	log.Fatal(server.ListenAndServe())
}
