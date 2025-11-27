package main

import (
	"fmt"
	"os"
)

const VERSION = "0.0.4"

func main() {
	//cli args
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("c", VERSION)
		fmt.Println("A build tool for c (like cargo for rust).")
		fmt.Println()
		fmt.Println("On Usage:")
		fmt.Println(" c help")
	} else if len(args) == 1 {
		if args[0] == "version" || args[0] == "-v" || args[0] == "--version" {
			fmt.Println("c", VERSION)
		}
	}
}
