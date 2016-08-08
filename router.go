package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/optimizer/{amount}", Optimizer)
	router.HandleFunc("/api/all", DisplayAll)
	router.HandleFunc("/", Tester)
	return router
}
