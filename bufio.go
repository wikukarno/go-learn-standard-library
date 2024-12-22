package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Open file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Failed to open file")
		return
	}

	// Read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Baris:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file")
		return
	}
}
