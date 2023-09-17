package main

import (
	"asciiart/pkg/asciiartutil"
	"asciiart/pkg/color"
	"asciiart/pkg/font"
	"asciiart/pkg/align"
	"asciiart/pkg/output"
	"asciiart/errors/"
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

	err := errorhandler.CheckFormat(os.Args)

	var userInput AsciiArt
	userInput.font = font.SetFont(args)
	userInput.lines = asciiartutil.GetAsciiLines()
	userInput.alignment = align.SetAlignment()
	userInput.outputFile = output.SetOutput()
	userInput.color.ansiCode = color.SetColor()
	userInput.color.lettersToColor = color.SetLetters()

	asciiartutil.PrintAsciiArt(userInput)
	
}
