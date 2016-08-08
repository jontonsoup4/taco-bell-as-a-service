package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	api := router.PathPrefix("/api").Subrouter();
	api.HandleFunc("/optimizer/{amount}", Optimizer)
	api.HandleFunc("/menu/{type}", MenuHandler)
	api.HandleFunc("/sort/{type}/{sortby}", SortHandler)
	api.HandleFunc("/value/{type}/{sortby}", ValueHandler)
	return router
}
