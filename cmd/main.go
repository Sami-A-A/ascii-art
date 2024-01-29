package main

import (
	"asciiart/errors"
	"asciiart/pkg/asciiartutil"
	"asciiart/pkg/options"
	"flag"
	"fmt"
	"os"
)

type Color struct {
	ansiCode string
	lettersToColor string 
}

type AsciiArt struct {
	font string
	lines [][]string
	alignment string
	outputFile string
	color Color
}

func main() {

	align := flag.String("align", "left", "aligns text")
	color := flag.String("color", "white", "assign color")
	output := flag.String("output", "output.txt", "assign output flag")

	flag.Parse()

	err := errorhandler.CheckFormat(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var userInput AsciiArt
	
	userInput.font = options.SetFont(args)
	userInput.lines = asciiartutil.GetAsciiLines()
	userInput.alignment = options.SetAlignment()
	userInput.outputFile = options.SetOutput()
	userInput.color.ansiCode = options.SetColor()
	userInput.color.lettersToColor = options.SetLetters()

	asciiartutil.InitAsciiArt(userInput)
	
}
