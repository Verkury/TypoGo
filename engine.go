package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	_ "sync"
	_ "time"

	. "github.com/Verkury/TypoGo/Prints"

	"github.com/eiannone/keyboard"
	"golang.org/x/term"
)

const (
        Reset = "\033[0m"
        Red   = "\033[31m"
        Green  = "\033[32m"
        Blue  = "\033[34m"
        White = "\033[37m"
    )


func clear() {
	fmt.Print("\033c")
}

func colorirLines(lines []string, userText string) []string {
    coloredLines := make([]string, len(lines))
    userTextIndex := 0
    
    for i, line := range lines {
        coloredLine := ""
        
        for _, char := range line {
            if userTextIndex < len(userText) {
                if userText[userTextIndex] == byte(char) {
                    coloredLine += Green + string(char) + Reset
                } else {
                    coloredLine += Red + string(char) + Reset
                }
                userTextIndex++
            } else {
                coloredLine += string(char)
            }
        }
        
        coloredLines[i] = coloredLine
    }
    
    return coloredLines
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
	key := fmt.Sprintf("%d", rand.Intn(21))

	return texts[key]
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

func scanKeys(output chan string) { // output - keys
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

func makeKeyboard(currentKey string) []string {
    keys := [][]string{
        {"[~] ", "[1] ", "[2] ", "[3] ", "[4] ", "[5] ", "[6] ", "[7] ", "[8] ", "[9] ", "[0] ", "[-] ", "[=] ", "[<-- ]"},
        {"[Tab] ", "[q] ", "[w] ", "[e] ", "[r] ", "[t] ", "[y] ", "[u] ", "[i] ", "[o] ", "[p] ", "[{] ", "[}] ", "[\\ ]"},
        {"[Caps] ", "[a] ", "[s] ", "[d] ", "[f] ", "[g] ", "[h] ", "[j] ", "[k] ", "[l] ", "[;] ", "['] ", "[Enter]"},
        {"[Shift] ", "[z] ", "[x] ", "[c] ", "[v] ", "[b] ", "[n] ", "[m] ", "[,] ", "[.] ", "[/] ", "[  Shift ]"},
        {"[Ctrl] ", "[Win] ", "[Alt] ", "[    Space   ] ", "[Alt] ", "[Fn] ", "[Menu] ", "[Ctrl]"},
    }

    keysForFind := [][]string{
        {"`", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", "=", "backspace"},
        {"tab", "q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "[", "]", "\\"},
        {"caps", "a", "s", "d", "f", "g", "h", "j", "k", "l", ";", "'", "enter"},
        {"shift", "z", "x", "c", "v", "b", "n", "m", ",", ".", "/", "shift"},
        {"ctrl", "win", "alt", " ", "alt", "fn", "menu", "ctrl"},
    }

    keysForFindCaps := [][]string{
        {"`", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", "=", "backspace"},
        {"tab", "Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "[", "]", "\\"},
        {"caps", "A", "S", "D", "F", "G", "H", "J", "K", "L", ":", "\"", "enter"},
        {"shift", "Z", "X", "C", "V", "B", "N", "M", ",", ".", "/", "shift"},
        {"ctrl", "win", "alt", " ", "alt", "fn", "menu", "ctrl"},
    }

	keybrd := make([]string, 5)
	shift := false
	for i := 0; i < len(keys); i++ {
        line := ""
        for j := 0; j < len(keys[i]); j++ {
            found := false

            if shift && keysForFind[i][j] == "shift" {
				line += Blue + keys[i][j] + Reset
				shift = false
				continue
			}

            if j < len(keysForFind[i]) && currentKey == keysForFind[i][j] {
                line += Blue + keys[i][j] + Reset
                found = true
            }
            
            if !found && j < len(keysForFindCaps[i]) && currentKey == keysForFindCaps[i][j] {
                line += Blue + keys[i][j] + Reset
                found = true
				shift = true
            }
            
            if !found && i == 0 && j == 13 && currentKey == "backspace" {
                line += Red + keys[i][j] + Reset
                found = true
            }
            
            if !found {
                line += White + keys[i][j] + Reset
            }
        }
        keybrd = append(keybrd, line)
    }
    
    return keybrd
}

func deliteEnters(text string) string {
	return strings.ReplaceAll(text, "\n", " ")
}

func getTerminalSize() (int, int){
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80, 14
	}
	return width, height
}

func finish(width int) {
	clear()
	fmt.Print("\n\n")
	PrintLineCenter("Thanks for palying!", 56, width)
}

func intro() {

}


func Start(text string) {
	width, height := getTerminalSize()
	lines := SplitText(text, width, 0.65)
	userText := ""

	key := make(chan string)

	go scanKeys(key)

	for {
		clear()
		ColorirLines := colorirLines(lines, userText)
		PrintLinesCenter(ColorirLines, 0, width)
		fmt.Print("\n")
		splitedUserText := SplitText(userText + "|", width, 0.65)
		PrintLinesCenter(splitedUserText, 0, width)
		if len(userText) != len(deliteEnters(text)) {
			if strings.HasPrefix(deliteEnters(text), userText) {
				PrintLinesCenter(makeKeyboard(string(text[len(userText)])), 0, width)
			} else {
				PrintLinesCenter(makeKeyboard("backspace"), 0, width)
			}
		} else {
			finish(width)
			break
		}
		k := <- key
		if (k == "backspace") {
			if len(userText) != 0 {
				userText = userText[:len(userText)-1]
			}
			continue
		} else if k != ""  {
			userText += k
		}
		if (k == "exit") {
			break
		}
	}
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