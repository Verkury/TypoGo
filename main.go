package main

import (
	"fmt"
	_ "fmt"
	"os"
	"path/filepath"
)

func main() {
	var filePath string
	// Check errors 
	if (len(os.Args) > 1) {
		filePath = os.Args[1]
	} else {
		filePath = ""
	}

	if filePath != "" {
		checkFile(filePath) // Checking file
		initialization(filePath) // Starting program
	}

	initialization("") // Starting program without file
}

func checkFile(path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("Ошибка при получении аболютного пути %v\n", err)
		return 
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		fmt.Printf("Файл не существует: %s\n", absPath)
		return 
	}

	path = absPath
}