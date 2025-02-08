package main

import (
	"fmt"
	"os"
)

func printColored(value string, bold bool, color_code int) {
	bold_val := 0
	if bold {
		bold_val = 1
	}

	fmt.Fprintf(os.Stdout, "\033[%d;%dm%s\033[0m", bold_val, color_code, value)
}

func Red(value string, bold bool) {
	printColored(value, bold, 31)
}

func Yellow(value string, bold bool) {
	printColored(value, bold, 33)
}

func Green(value string, bold bool) {
	printColored(value, bold, 32)
}
