package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/oleiade/reflections"
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

func SortHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	items, err := LoadJSON(w, vars["type"]);
	if err != nil {
		return
	}
	propertyName := getPropertyName(strings.ToLower(vars["sortby"]));
	if !stringInSlice(vars["sortby"], SortByOptions){
		json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: fmt.Sprintf("No filter by that name. Try: %q", SortByOptions),
		});
		return
	}
	stringQueries := r.URL.Query();
	reverse := stringQueries.Get("reverse");
	var sorted Items;
	if reverse == "true" {
		sorted = reverseSortProperties(propertyName, items);
	} else {
		sorted = sortProperties(propertyName, items);
	}
	json.NewEncoder(w).Encode(sorted);
}

func ValueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	items, err := LoadJSON(w, vars["type"]);
	if err != nil {
		return
	}
	itemsMap := make(map[float64]Item);
	keys := make([]float64, len(items));
	propertyName := getPropertyName(strings.ToLower(vars["sortby"]));
	for index, item := range items {
		propVal, _ := reflections.GetField(item, propertyName);
		var value float64;
		switch propVal.(type) {
		case int:
			value = item.Cost / float64(propVal.(int));
		case float64:
			value = item.Cost / propVal.(float64);
		}
		itemsMap[value] = item;
		keys[index] = value;
	}
	stringQueries := r.URL.Query();
	reverse := stringQueries.Get("reverse");
	if reverse == "true" {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)));
	} else {
		sort.Float64s(keys);
	}
	returnItems := make([]Item, len(keys))
	for index, item := range keys {
		returnItems[index] = itemsMap[item];
	}
	json.NewEncoder(w).Encode(returnItems);
}
