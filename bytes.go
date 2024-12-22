package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Hello, World!")
	fmt.Println(bytes.Contains(data, []byte("World"))) // true

	upper := bytes.ToUpper(data)
	fmt.Println(string(upper)) // HELLO, WORLD!

	lower := bytes.ToLower(data)
	fmt.Println(string(lower)) // hello, world!

	replaceAll := bytes.ReplaceAll(data, []byte("World"), []byte("Golang"))
	fmt.Println(string(replaceAll)) // Hello, Golang!
}
