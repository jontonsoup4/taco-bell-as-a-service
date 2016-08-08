package main

import (
	"github.com/oleiade/reflections"
	"github.com/bradfitz/slice"
)

func sortProperties(field string, items Items) Items {
	slice.Sort(items, func(i, j int) bool {
		iValue, _ := reflections.GetField(items[i], field)
		jValue, _ := reflections.GetField(items[j], field)
		switch iValue.(type) {
		case int:
			return iValue.(int) > jValue.(int);
		case float64:
			return iValue.(float64) > jValue.(float64);
		}
		return false
	})
	return items
}

func reverseSortProperties(field string, items Items) Items {
	slice.Sort(items, func(i, j int) bool {
		iValue, _ := reflections.GetField(items[i], field)
		jValue, _ := reflections.GetField(items[j], field)
		switch iValue.(type) {
		case int:
			return iValue.(int) < jValue.(int);
		case float64:
			return iValue.(float64) < jValue.(float64);
		}
		return false
	})
	return items
}
