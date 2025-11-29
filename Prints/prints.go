package Prints

import (
	"fmt"
	"math"
	"time"
	"strings"
	"regexp"
)

func visibleLen(s string) int {
    re := regexp.MustCompile(`\033\[[0-9;]*[a-zA-Z]`)
    clean := re.ReplaceAllString(s, "")
    return len(clean)
}

func PrintLine(line string, delay int) {
	for _, v := range line {
		fmt.Print(string(v))
		if delay != 0 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
	fmt.Print("\n")
}

func PrintLineCenter(line string, delay int, width int) {
	indent := (width - visibleLen(line)) / 2

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

func SplitText(text string, width int, procent float64) []string  {
	var length int = int(math.Round(float64(width) * procent))
	
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 <= length {
			if currentLine != "" {
				currentLine += " " + word
			} else {
				currentLine = word
			}
		} else {
			if currentLine != "" {
				lines = append(lines, currentLine + " ")
			}
			currentLine = word
		}
	}
	
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	
	return lines
}