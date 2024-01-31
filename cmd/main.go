package main

import (
	"asciiart/errors"
	"asciiart/pkg/asciiartutil"
	"asciiart/pkg/options"
	"flag"
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	reverse := flag.Bool("reverse", false, "determines if to be printed in reverse")
	output := flag.String("output", "output.txt", "assign output flag")
	align := flag.String("align", "left", "aligns text")
	color := flag.String("color", "white", "assign color")

	flag.Parse()

	err := errorhandler.CheckFormat(args)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	userInput := new(asciiartutil.AsciiArt)

	userInput.Reverse = *reverse
	userInput.OutputFile = *output
	userInput.Alignment = *align
	userInput.Color.AnsiCode = *color
	userInput.Color.LettersToColor = options.SetLetters(args)
	userInput.Font = options.SetFont(args)

	asciiartutil.InitAsciiArt(userInput)
	
}
