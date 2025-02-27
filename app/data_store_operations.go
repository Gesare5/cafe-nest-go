package main

import (
	"fmt"
	"os"
	"strings"
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
	dataList := strings.Split(string(data), "\n")
	return dataList
}

func saveToStore(filename string, data string) {
	// ADD/APPEND DATA TO STORE
	fmt.Println("Opening file")
	oldData, err := os.ReadFile(filename)
	dataString := string(oldData)

	fmt.Println("Writing file")
	file, err := os.Create(filename)
	check(err)

	_, err = file.WriteString(dataString + data + "\n")
	check(err)
	fmt.Printf("File name: %s \n", file.Name())
}

func overwriteStore(filename string, data string) error {
	// POPULATE DATA STORE WITH ALL NEW DATA
	fmt.Println("Writing file")
	file, err := os.Create(filename)
	check(err)

	_, err = file.WriteString(data + "\n")
	check(err)
	fmt.Printf("File name: %s \n", file.Name())
	return nil
}
