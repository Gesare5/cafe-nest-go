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

func craftACoffee(coffeeType string) {
	println("creafting coffee........")
}

func handleSelection(choice int, quantity int) {
	println("Handle selection")
	COFFEE_SELECTION_LIST := generateCoffeeList()
	choiceList := strings.Split(COFFEE_SELECTION_LIST[choice-1], " ")
	coffeeType := " "
	if len(choiceList) >= 3 {
		choiceList = []string{choiceList[1], choiceList[2]}
		coffeeType = strings.ToLower(strings.Join(choiceList, " "))
	} else {
		coffeeType = strings.ToLower(choiceList[1])
	}
	println(coffeeType)
	for i := 0; i < quantity; i++ {
		craftACoffee(coffeeType)
	}
	return
}

func handlePrelimAction(action string) {
	COFFEE_TABLE_SELECTION_LIST := generateCoffeeList()
	switch action {
	case "a":
		COFFEE_TABLE_SELECTION_LIST = append(COFFEE_TABLE_SELECTION_LIST, "0: Done!!")
		// looping := true
		for i := 0; i < 30; i++ {
			print("")
			generate_table(COFFEE_TABLE_SELECTION_LIST)
			coffeeChoice := ""
			fmt.Scanln(&coffeeChoice)
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
					fmt.Scanln(&quantity)
					// (strings.TrimSpace(input)
					quantityInt, err := strconv.Atoi(quantity)
					if err == nil {
						handleSelection(coffeeChoiceInt, quantityInt)
						println("")
					} else {
						println("Invalid value!!")
						println("")
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
	fmt.Println("Date", time.Now())
	println("Welcome to TeleBaristas!")
	println("")
	// keepLooping := true
	for i := 0; i < 15; i++ {
		for _, value := range prelimList() {
			println(value)
		}
		prelimAction := "a"
		fmt.Scanln(&prelimAction)
		prelimAction = strings.ToLower(prelimAction)
		if prelimAction == "c" {
			break
		} else {
			handlePrelimAction(prelimAction)
		}
	}
}
