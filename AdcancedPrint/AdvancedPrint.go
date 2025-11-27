package adcancedprint

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
	
	for v := range line {
		fmt.Print(v)
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	fmt.Print("\n")
}