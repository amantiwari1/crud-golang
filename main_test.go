package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amantiwari1/crud-golang/pkg/routes"
	"github.com/gorilla/mux"
)

func Test_getUser(t *testing.T) {
	r := mux.NewRouter()
	routes.RegisterWeatherRoutes(r)

	http.Handle("/", r)

	ts := httptest.NewServer(r)
	defer ts.Close()
	res, err := http.Get(ts.URL + "/user")

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}
