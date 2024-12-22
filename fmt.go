package main

import "fmt"

func main() {
	firstName := "Wiku"
	middleName := "Karno"
	lastName := "Prasetyagama"
	age := 24

	fmt.Println("Hai ", firstName, middleName, lastName)
	fmt.Printf("Hai %s %s %s\n", firstName, middleName, lastName)
	fmt.Printf("Hai, umur saya sekarang adalah %d tahun\n", age)
	fmt.Print("Hai ", firstName, " ", middleName, " ", lastName, "\n")
}
