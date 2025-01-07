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
	for _, value := range readList {
		innerList := strings.Split(value, ",")
		var err error
		if innerList[0] == coffeeType {
			coffee["milk"], err = strconv.ParseFloat(innerList[1], 32)
			coffee["coffee"], err = strconv.ParseFloat(innerList[2], 32)
			coffee["sugar"], err = strconv.ParseFloat(innerList[3], 32)
			coffee["vanilla"], err = strconv.ParseFloat(innerList[4], 32)
			coffee["cocoa"], err = strconv.ParseFloat(innerList[5], 32)
			// TODO: refactor this
			check(err)
		}
	}

	// subtract quantities
	newTotals := subtractUsedQuantititesFromTotal(coffee, coffeeType)

	// update inventory
	dataString := "coffee," + fmt.Sprintf("%f", newTotals["coffee"]) + "\n" +
		"milk," + fmt.Sprintf("%f", newTotals["milk"]) + "\n" +
		"sugar," + fmt.Sprintf("%f", newTotals["sugar"]) + "\n" +
		"vanilla," + fmt.Sprintf("%f", newTotals["vanilla"]) + "\n" +
		"cocoa," + fmt.Sprintf("%f", newTotals["cocoa"]) + "\n"
	//TODO: refactor this
	saveToStore("inventory.csv", dataString)

	// Print Report - Much Later
}
