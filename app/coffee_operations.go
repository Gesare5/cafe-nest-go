package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

	//  Save to daily report
	reportName := fmt.Sprintf("Daily_Sales_%s.csv", time.Now().Format(time.DateOnly))
	reportDataString := fmt.Sprintf("%s,%.2f,%s", coffeeType, coffee["cost"], time.Now().Format(time.DateTime))
	saveToStore(reportName, reportDataString)
}

func addNewCoffeeItem() {
	println("Add New Coffee Type!")

	// fetch headers in coffee_items.csv
	readList := readFromStore("coffee_items.csv")
	headers := []string{}
	newCoffeeItem := []string{}
	for i, value := range readList {
		if i == 0 {
			headers = strings.Split(value, ",")
		} else {
			break
		}
	}
	input := ""
	for _, header := range headers {
		if header == "coffee_type" {
			println("Name of Coffee? ")
			fmt.Scanln(&input)
		} else {
			fmt.Printf("%s amount: ", header)
			fmt.Scanln(&input)
			if err := validateNumberInput(input); err != nil {
				println("Invalid Input!! Try again.")
				break
			}
		}
		newCoffeeItem = append(newCoffeeItem, input)
	}
	println("")
	if len(newCoffeeItem) == len(headers) {
		newCoffeeItemString := strings.Join(newCoffeeItem, ",")
		saveToStore("coffee_items.csv", newCoffeeItemString)
	}
}

func removeCoffeeItem() {
	println("Remove Coffee Type!")
	println("")
}

func replenishInventory(supplyItem string, amount string) {
	println("Replenish Inventory")
	floatAmount, err := strconv.ParseFloat(amount, 32)
	println(amount)
	println(floatAmount)
	println("")
	check(err)
}

func manageCoffeeItems(choice int) {
	COFFEE_TABLE_SELECTION_LIST := generateCoffeeList()

	if choice == 1 {
		addNewCoffeeItem()
	} else if choice == 2 {
		removeCoffeeItem()
	} else if choice == 4 {
		supplies := []string{"1: Milk", "2: Coffee", "3: Sugar", "4: Cocoa", "5: Vanilla"}
		// Read this from file
		println("Select supply: ")
		generateTable(supplies)
		supply := 0
		fmt.Scanln(&supply)
		println("")
		println("Amount in grams/millilitre: ")
		amount := ""
		fmt.Scanln(&amount)
		println("")
		replenishInventory(supplies[supply-1], amount)
	} else {
		generateTable(COFFEE_TABLE_SELECTION_LIST)
	}
}
