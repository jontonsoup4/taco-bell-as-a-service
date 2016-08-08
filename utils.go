package main

import (
	"math"
	"strings"
)

var SortByOptions = []string{"cost", "calories", "caloriesfromfat",
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

func getPropertyName(name string) string {
	name = strings.ToLower(name);
	var properties = map[string]string{
		"cost": "Cost",
		"calories": "Calories",
		"caloriesfromfat": "CaloriesFromFat",
		"totalfat": "TotalFat",
		"saturatedfat": "SaturatedFat",
		"transfat": "TransFat",
		"cholesterol": "Cholesterol",
		"sodium": "Sodium",
		"carbohydrates": "Carbohydrates",
		"dietaryfiber": "DietaryFiber",
		"sugars": "Sugars",
		"protein": "Protein",
	};
	return properties[name]
}
