package main

import (
	"strconv"
	"strings"
)

func GetTotals() map[string]float64 {
	dataList := readFromStore("totals.csv")
	totals := map[string]float64{}
	var err error
	for _, value := range dataList {
		supplies := strings.Split(value, ",")
		totals[supplies[0]], err = strconv.ParseFloat(strings.Trim(supplies[1], " "), 32)
		check(err)
	}
	return totals
}

func GetThresholds() map[string]float64 {
	dataList := readFromStore("thresholds.csv")
	thresholds := map[string]float64{}
	var err error
	for _, value := range dataList {
		supplies := strings.Split(value, ",")
		thresholds[supplies[0]], err = strconv.ParseFloat(strings.Trim(supplies[1], " "), 32)
		check(err)
	}
	return thresholds
}

func generateCoffeeList() []string {
	readList := readFromStore("coffee_items.csv")
	coffeeList := []string{}
	for index, value := range readList {
		if index > 0 {
			innerList := strings.Split(value, ",")
			coffeeItem := strconv.Itoa(index) + ": " + string(innerList[0])
			coffeeList = append(coffeeList, coffeeItem)
		}
	}
	return coffeeList
}
