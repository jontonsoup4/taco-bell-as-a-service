package main

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"errors"
)

func LoadJSON(w http.ResponseWriter, filename string) (Items, error) {
	allowed := []string{"all", "drinks", "food", "sauces"};
	if !stringInSlice(filename, allowed){
		json.NewEncoder(w).Encode(jsonErr{
			Code:http.StatusNotFound,
			Text: "'" + filename + "' not found. Try available paths: all, drinks, food, sauces.",
		});
		return Items{}, errors.New("NotFound");
	}
	raw, err := ioutil.ReadFile("./menu/" + filename + ".json")
	if err != nil {
		json.NewEncoder(w).Encode(jsonErr{
			Code:http.StatusInternalServerError,
			Text: err.Error(),
		});
		return Items{}, errors.New("InternalServerError");
	}
	var output Items;
	json.Unmarshal(raw, &output);
	return output, nil;
}