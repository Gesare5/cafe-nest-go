package main

import "strconv"

func generateTable(lst []string) {
	for _, value := range lst {
		println(value)
	}
	println("")
}

func validateNumberInput(input string) error {
	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}
	return nil
}
