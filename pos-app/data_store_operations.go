package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromStore(filename string) []string {
	// READ FROM FILE
	data, err := os.ReadFile(filename)
	check(err)
	fmt.Print(string(data))

	var readList []string
	return readList
}

func writeToStore(filename string, data []string) error {
	return nil
}

func overwriteStore(filename string, data []string) error {
	return nil
}
