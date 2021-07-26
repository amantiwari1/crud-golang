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
	log.Println("started server Running at http://localhost:8800")
	log.Fatal(http.ListenAndServe("localhost:8800", route))

}
