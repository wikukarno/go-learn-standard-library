package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = fmt.Errorf("Data not found")

func findData(id int) error {

	if id == 0 {
		return ErrNotFound
	}

	return nil

}

func main() {
	name := "Wiku"

	// Error
	err := fmt.Errorf("hai %s, Your data is not found", name)
	fmt.Println(err.Error())

	// Error with variable
	errFindData := findData(0)
	if errors.Is(errFindData, ErrNotFound) {
		fmt.Println("Data not found")
	} else {
		fmt.Println("Data found")
	}
}
