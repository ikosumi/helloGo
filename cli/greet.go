package main

import (
	"fmt"
	"os"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m\n"
	WarningColor = "\033[1;33m%s\033[0m\n"
	ErrorColor   = "\033[1;31m%s\033[0m\n"
	DebugColor   = "\033[0;36m%s\033[0m\n"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	if len(argsWithProg) < 4 {
		fmt.Printf(ErrorColor, "Need at least 3 parameters.")
		return
	}

	thirdArg := argsWithProg[3]

	fmt.Printf(InfoColor, argsWithProg)
	fmt.Printf(DebugColor, argsWithoutProg)
	fmt.Printf(WarningColor, thirdArg)
}
