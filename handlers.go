package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	"strings"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	filename := mux.Vars(r)["type"];
	output, err := LoadJSON(w, filename);
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(output);
}

func Optimizer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	vars := mux.Vars(r);
	_, err := strconv.ParseFloat(vars["amount"], 64);
	if err != nil {
		fmt.Fprint(w, err);
		return;
	}
	json.NewEncoder(w).Encode(vars);
}

func SortHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	vars := mux.Vars(r);
	items, err := LoadJSON(w, vars["type"]);
	if err != nil {
		return
	}
	sortByFilter := strings.ToLower(vars["sortby"]);
	if sortByFilter == "" {
		sortByFilter = "calories";
	} else if !stringInSlice(sortByFilter, SortByOptions){
		json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: fmt.Sprintf("No filter by that name. Try: %q", SortByOptions),
		});
		return
	}
	itemsMap := make(map[float64]Item);
	keys := make([]float64, len(items));
	for index, item := range items {
		value := item.Cost/sortBy(sortByFilter, item);
		itemsMap[value] = item;
		keys[index] = value;
	}
	sort.Float64s(keys);
	returnItems := make([]Item, len(keys))
	for index, item := range keys {
		returnItems[index] = itemsMap[item];
	}
	json.NewEncoder(w).Encode(returnItems);
}
