package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/{amount}", Optimizer)
	router.HandleFunc("/", Tester)
	return router
}
