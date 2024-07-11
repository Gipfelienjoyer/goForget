package util

import "fmt"

func ScanLine(PrintText string) any {
	var userInput any
	fmt.Println(PrintText)
	scanln, err := fmt.Scanln(&userInput)
	if err != nil {

	}

	return scanln
}
