package main

import (
	"fmt"
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

func handle_prelim_action(prelimAction string) {
	println("yooooooooo!")
	println(prelimAction)
}

func main() {
	println("Good day!")
	fmt.Print("Date", time.Now())
	println("Welcome to TeleBaristas!")
	println("")
	keep_looping := true
	for keep_looping == true {
		for _, value := range prelimList() {
			println(value)
		}
		prelimAction := ""
		fmt.Scanf("%s", &prelimAction)
		prelimAction = strings.ToLower(prelimAction)
		if prelimAction == "c" {
			break
		} else {
			handle_prelim_action(prelimAction)
		}
	}
}
