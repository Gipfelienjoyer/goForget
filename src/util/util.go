package util

import (
	"fmt"
)

func ScanLine(PrintText string, Scan bool) string {
	UserInput := ""
	fmt.Println(PrintText)
	if Scan {
		fmt.Scanln(&UserInput)
	}
	return UserInput
}
