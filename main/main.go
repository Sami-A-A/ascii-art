package main

import (
	ascii "ascii/pkg"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Handle error if no arguments are passed
	if len(os.Args) < 2 {
		fmt.Println("ERROR: Insufficient arguments")
		return
	}

	// Create output flag and return output file name
	fileName, err := ascii.OutputFlag()
	if err {
		return
	}

	// Error handling; flag arguments: should be a single arg
	err = ascii.FlagArgErrors()
	if err {
		return
	}

	// Convert input to art and print output to file given in output flag
	art := ascii.AsciiArt(flag.Args()[0])
	ascii.PrintOutput(art, fileName)
}
