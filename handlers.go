package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

func DisplayAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	raw, err := ioutil.ReadFile("./menu/items.json")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	var output Items;
	json.Unmarshal(raw, &output)
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

func Tester(w http.ResponseWriter, r *http.Request)  {
	taco := Item{
		Name: "Taco",
		Cost: 1.59,
		Calories: 500,
	}
	burrito := Item{
		Name: "Burrito",
		Cost: 2.49,
		Calories: 800,
	}
	pizza := Item{
		Name: "Pizza",
		Cost: 2.49,
		Calories: 2000,
	}
	items := Items{taco, burrito, pizza};
	itemsMap := make(map[float64]Item);
	keys := make([]float64, len(items));
	for index, item := range items {
		value := item.Cost/float64(item.Calories);
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
