package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func prelimList() []string {
	prelimList := []string{
		"A: Order Coffee",
		"B: Manage Coffee Items",
		"C: Exit",
	}
	return prelimList
}

func manageCoffeeItemsList() []string {
	manageCoffeeList := []string{
		"1: Add Coffee Item",
		"2: Delete Coffee Item",
		"3: View Coffee Items",
	}
	return manageCoffeeList
}

func manageCoffeeItems(coffeeChoice int) {
	println("manageee!")
	println(coffeeChoice)
}

func generate_table(lst []string) {
	for _, value := range lst {
		println(value)
	}
}

func handleSelection(coffeeChoice int, quantity int) {
	println("Handle selection")
}

func handlePrelimAction(action string) {
	COFFEE_TABLE_SELECTION_LIST := []string{"0 Done!!", "0 Done!!"}
	switch action {
	case "a":
		COFFEE_TABLE_SELECTION_LIST = append(COFFEE_TABLE_SELECTION_LIST, "3 Done!!")
		looping := true
		for looping == true {
			print("")
			generate_table(COFFEE_TABLE_SELECTION_LIST)
			coffeeChoice := ""
			go fmt.Scanf("%d", &coffeeChoice)
			coffeeChoiceInt, err := strconv.Atoi(coffeeChoice)
			if err != nil {
				coffeeChoiceInt = 0
			}

			if coffeeChoiceInt < len(COFFEE_TABLE_SELECTION_LIST) {
				if coffeeChoiceInt == 0 {
					break
				} else {
					print("How many? ")
					quantity := ""
					fmt.Scanf("%d", &quantity)
					quantityInt, err := strconv.Atoi(quantity)
					if err == nil {
						handleSelection(coffeeChoiceInt, quantityInt)
						print("")
					} else {
						print("Invalid value!!")
						print("")
						break
					}
				}
			} else {
				print("Invalid Choice!!")
				continue
			}
		}
	case "b":
		print("")
		generate_table(manageCoffeeItemsList())
		coffeeItemsChoice := 0
		fmt.Scanf("%d", &coffeeItemsChoice)
		manageCoffeeItems(coffeeItemsChoice)
		print("")

	default:
		return
	}
}

func main() {
	println("Good day!")
	fmt.Print("Date", time.Now())
	println("Welcome to TeleBaristas!")
	println("")
	keepLooping := true
	for keepLooping == true {
		for _, value := range prelimList() {
			println(value)
		}
		prelimAction := ""
		fmt.Scanf("%s", &prelimAction)
		prelimAction = strings.ToLower(prelimAction)
		if prelimAction == "c" {
			break
		} else {
			handlePrelimAction(prelimAction)
		}
	}
}
