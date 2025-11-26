package main

import (
	"fmt"
	"os"
)


func Start(text string) {

}

func initialization(absPath string) {
	if absPath == "" {
		text := getText()
		text = splitText(text)
	} else {
		file, err := os.ReadFile(absPath)
		if err != nil {
			fmt.Printf("Ошибка при чтении файла %v\n", err)
			return
		}
		text := splitText(string(file))
		Start(text)
	}
}

func getText() string {
	
	return ""
}

func splitText(text string) string {

	return ""
}

func getWidth() {

}

