package main

import (
	"fmt"
	"strconv"
)

func main() {
	numberStr := "100"
	number, err := strconv.Atoi(numberStr) // Convert string to int
	if err == nil {
		fmt.Println(number + 100) // 200
	}

	floatStr := "3.8"
	float, err := strconv.ParseFloat(floatStr, 64) // Convert string to float64

	fmt.Println(float)

	if err == nil {
		fmt.Println(float * 1.2) // 5
	}

	// Convert int to string
	number2 := 200
	numberStr2 := strconv.Itoa(number2)
	fmt.Println("Number to string:", numberStr2) // 200

}
