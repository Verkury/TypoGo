package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	_ "sync"
	_ "time"

	. "github.com/Verkury/TypoGo/Prints"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)



func colorirLines(lines []string, userText string) []string {
	return []string{}
}


func Start(text string) {
	width, height := getTerminalSize()
	lines := SplitText(text, width, 0.65)
	userText := ""

	key := make(chan string)

	go scanKeys(key)

	for {
		k := <- key
		if k != "" {
			userText += k
		}
		if (k == "exit") {
			break
		}
		fmt.Println(userText, k)

	}

	_ = lines
	_ = height
}

func initialization(absPath string) { 
	if absPath == "" {
		text := getText()
		Start(text)
	} else {
		file, err := os.ReadFile(absPath)
		if err != nil {
			fmt.Printf("Ошибка при чтении файла %v\n", err)
			return
		}
		Start(string(file))
	}
}

func getText() string { // Parsing text from json
	data, err := os.ReadFile("Texts.json")
	if err != nil {
		log.Fatal(err)
	}

	var texts map[string]string
	err = json.Unmarshal(data, &texts)
	if err != nil {
		log.Fatal(err)
	}

	return texts[string(rand.Intn(21))]
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
	case keyboard.KeyEsc, keyboard.KeyCtrlC, keyboard.KeyCtrlD:
		return "exit"
	default:
		if key == 0 && char != 0 {
			return string(char)
		}
		return ""
	}
}

func scanKeys(output chan string) { // input - work status, output - keys
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		keyStr := keyToString(key, char)

		if keyStr != "" {
			output <- keyStr
		}
		if (keyStr == "exit") {
			break
		}
	}
}