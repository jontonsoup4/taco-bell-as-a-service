package main

import (
	"math"
)

var SortByOptions = []string{"calories", "caloriesfromfat",
"totalfat", "saturatedfat", "transfat", "cholesterol", "sodium", "totalcarbohydrates", "dietaryfiber",
"sugars", "protein"}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func sortBy(s string, item Item) float64 {
	switch s {
	case "calories":
		return float64(item.Calories);
	case "caloriesfromfat":
		return float64(item.CaloriesFromFat);
	case "totalfat":
		return float64(item.TotalFat);
	case "saturatedfat":
		return float64(item.SaturatedFat);
	case "transfat":
		return float64(item.TransFat);
	case "cholesterol":
		return float64(item.Cholesterol);
	case "sodium":
		return float64(item.Sodium);
	case "totalcarbohydrates":
		return float64(item.TotalCarbohydrates);
	case "dietaryfiber":
		return float64(item.DietaryFiber);
	case "sugars":
		return float64(item.Sugars);
	case "protein":
		return float64(item.Protein);
	}
	return float64(0)
}