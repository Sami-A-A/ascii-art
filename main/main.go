package main

import (
	// ascii "ascii/pkg/printAscii"
	ascii "ascii/pkg/output"
	"fmt"
	"os"
)

func main() {
	// Error handling os arguments
	// Run only if 1 arg
	if len(os.Args) < 2 {
		fmt.Println("ERROR: No arguments to print")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("ERROR: Place all text in a single argument")
		return
	} else if len(os.Args[1]) < 1 {
		return
	}
	// ascii.AsciiArt()
	art := ascii.Output()
	fmt.Print(art)
}
