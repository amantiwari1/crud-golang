package routes

import (
	"github.com/amantiwari1/crud-golang/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterWeatherRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser).Methods("DELETE")
}
