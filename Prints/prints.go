package Prints

import (
	"fmt"
	"time"
)

func PrintLine(line string, delay int) {
	for v := range line {
		fmt.Print(v)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	fmt.Print("\n")
}

func PrintLineCenter(line string, delay int, width int) {
	indent := (width - len(line)) / 2

	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}
	
	PrintLine(line, delay)
}

func PrintLinesCenter(lines []string, delay int, width int) {
	for _, line := range lines {
		PrintLineCenter(line, delay, width)
	}
}

func SplitText(text string, width, procent int) []string { // Spliting text to send in print func
	length := width * procent
	
	runes := []rune(text)
	lines := make([]string, len(text) / length + 3)

	var line string
	for _, value := range runes {
		if len(line) <= length  && value != ' ' {
			line += string(value)
		} else {
			lines = append(lines, line)
			line = ""
		}
	} 
	return lines
}