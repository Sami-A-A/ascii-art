package main

import (
	ascii "ascii/pkg"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Handle error if no arguments are passed
	if len(os.Args) < 2 {
		fmt.Println("ERROR: Insufficient arguments")
		return
	}

	// Create a flag named "output"
	outputFlag := flag.String("output", "output.txt", "assign output file")
	flag.Parse()

	// Confirm flag format at input
	if strings.Contains(os.Args[1], "-output") && !strings.HasPrefix(os.Args[1], "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	// Error handling: flag arguments
	// Run only if 1 arg after flag
	if len(flag.Args()) < 1 {
		fmt.Println("ERROR: No arguments to print")
		return
	} else if len(flag.Args()) > 1 {
		fmt.Println("ERROR: Place all text in a single argument")
		return
	} else if len(flag.Args()[0]) < 1 {
		return
	}

	// Convert input to art and output to file named <*outputFlag>
	art := ascii.AsciiArt(flag.Args()[0])
	ascii.Output(art, outputFlag)
}
