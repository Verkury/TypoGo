package main

import (
	"fmt"
	"os"
	_ "sync"
	"log"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"

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

func getText() string { // Parsing text from json
	
	return ""
}

func splitText(text string) string { // Spliting text to send in print func

	return ""
}

func getTerminalSize() (int, int){
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80, 14
	}
	return width, height
}

func clear() {
	fmt.Print("\033c")
}

func keyToString(key keyboard.Key, char rune) string {
	switch key {
	case keyboard.KeyBackspace, keyboard.KeyBackspace2, keyboard.KeyDelete:
		return "backspace"
	case keyboard.KeySpace:
		return " "
	case keyboard.KeyEsc:
		return "esc"
	default:
		if key == 0 && char != 0 {
			return string(char)
		}
		return ""
	}
}

func getKeys(input chan bool, output chan string) { // input - work status, output - keys
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		if (<- input) {
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Fatal(err)
			}

			keyStr := keyToString(key, char)

			if keyStr != "" {
				output <- keyStr
			}
		} else {
			break
		}
	}
}