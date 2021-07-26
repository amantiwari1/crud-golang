package main

import (
	"log"
	"net/http"

	"github.com/amantiwari1/crud-golang/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	routes.RegisterWeatherRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("9000", route))

}


