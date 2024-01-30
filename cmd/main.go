package main

import (
	"asciiart/errors"
	"asciiart/pkg/asciiartutil"
	"flag"
	"fmt"
	"os"
)

func main() {

	reverse := flag.Bool("reverse", false, "determines if to be printed in reverse")
	output := flag.String("output", "output.txt", "assign output flag")
	align := flag.String("align", "left", "aligns text")
	color := flag.String("color", "white", "assign color")

	flag.Parse()

	letters, text, banner, err := errorhandler.CheckFormat(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userInput := asciiartutil.AsciiArt{
		Text: text,
		Reverse: *reverse,
		OutputFile: *output,
		Alignment: *align,
		Color: asciiartutil.Color {
			AnsiCode: *color,
			LettersToColor: letters,
		},
		Font: banner,
	}


	asciiartutil.InitAsciiArt(userInput)
	
}
