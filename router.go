package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api").Subrouter();
	api.HandleFunc("/optimizer/{amount}", Optimizer)
	api.HandleFunc("/menu/{type}", MenuHandler)
	api.HandleFunc("/sort/{type}/{sortby}", SortHandler)
	return router
}
