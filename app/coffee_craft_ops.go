package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetTotals() map[string]float64 {
	dataList := readFromStore("totals.csv")
	totals := map[string]float64{}
	var err error
	for _, value := range dataList {
		supplies := strings.Split(value, ",")
		if len(supplies) > 1 {
			totals[supplies[0]], err = strconv.ParseFloat(strings.Trim(supplies[1], " "), 32)
			check(err)
		}
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

func subtractUsedQuantititesFromTotal(coffeeTotals map[string]float64, coffeeType string) map[string]float64 {
	totals := GetTotals()
	totals["milk"] = totals["milk"] - coffeeTotals["milk"]
	totals["sugar"] = totals["sugar"] - coffeeTotals["sugar"]
	totals["coffee"] = totals["coffee"] - coffeeTotals["coffee"]

	if strings.Contains(coffeeType, "vanilla") {
		totals["vanilla"] = totals["vanilla"] - coffeeTotals["vanilla"]
	}
	if strings.Contains(coffeeType, "mocha") {
		totals["cocoa"] = totals["cocoa"] - coffeeTotals["cocoa"]
	}
	return totals
}

func craftACoffee(coffeeType string) {
	println("creafting coffee........")
	// read from store
	readList := readFromStore("coffee_items.csv")
	coffee := map[string]float64{}
	headers := []string{}
	for i, value := range readList {
		var err error
		coffeeIngredientsList := strings.Split(value, ",")
		if i == 0 {
			headers = strings.Split(value, ",")
		}
		if coffeeIngredientsList[0] == coffeeType {
			for index, value := range coffeeIngredientsList {
				if index > 0 {
					coffee[headers[index]], err = strconv.ParseFloat(value, 32)
					check(err)
				}
			}
		}
	}

	// subtract quantities
	newTotals := subtractUsedQuantititesFromTotal(coffee, coffeeType)

	// update inventory
	dataString := ""
	for key, value := range newTotals {
		dataString = dataString + key + "," + fmt.Sprintf("%.2f", value) + "\n"
	}
	saveToStore("totals.csv", dataString)

	// Print Report - Much Later
}
