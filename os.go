package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	hostName, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostName)
	}

	// Make a file
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}
	defer file.Close()

	// Write to file
	_, err = file.WriteString("Hello World")
	if err != nil {
		fmt.Println("Failed to write to file")
		return
	}

	fmt.Println("Success create and write to file")

	// Read file
	readFile, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Failed to read file")
		return
	}

	fmt.Println("Read file:", string(readFile))

}
