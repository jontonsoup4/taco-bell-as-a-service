package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"errors"
)

func LoadJSON(w http.ResponseWriter, filename string) (Items, error) {
	raw, err := ioutil.ReadFile("./menu/" + filename + ".json");
	if err != nil {
		json.NewEncoder(w).Encode(jsonErr{
			Code:http.StatusNotFound,
			Text: "Invalid endpoint: " + filename,
		});
		return Items{}, errors.New("NotFound");
	}
	var output Items;
	json.Unmarshal(raw, &output);
	return output, nil;
}