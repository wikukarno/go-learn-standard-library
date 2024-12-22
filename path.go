package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("folder", "sub-folder", "file.txt")
	fmt.Println(path)
	fmt.Println(filepath.Abs(path)) // absolute path

	// Dir and File
	dir := filepath.Dir(path)
	file := filepath.Base(path)

	fmt.Println("Dir:", dir)
	fmt.Println("File:", file)

	// Ext
	ext := filepath.Ext(path)
	fmt.Println("Ext:", ext)
}
