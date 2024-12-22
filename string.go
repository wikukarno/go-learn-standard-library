package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"
	fmt.Println(strings.ToUpper(text)) // HELLO, WORLD!
	fmt.Println(strings.ToLower(text)) // hello, world!

	fmt.Println(strings.Replace(text, "World", "Golang", 1)) // Hello, Golang! (1 is the number of replacements)
	fmt.Println(strings.ReplaceAll(text, "World", "Golang")) // Hello, Golang!

	fmt.Println(strings.CutPrefix(text, "Hello, "))  // World!
	fmt.Println(strings.CutSuffix(text, ", World!")) // Hello
	fmt.Println(strings.TrimPrefix(text, "Hello, ")) // World!
}
