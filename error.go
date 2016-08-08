package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func notFound(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(jsonErr{
		Code: http.StatusNotFound,
		Text: fmt.Sprintf("No endpoint found for '%s'", r.URL),
	});
}
