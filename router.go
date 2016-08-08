package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Tester)
	api := router.PathPrefix("/api").Subrouter();
	api.HandleFunc("/optimizer/{amount}", Optimizer)
	api.HandleFunc("/menu/{type}", MenuHandler)
	return router
}
