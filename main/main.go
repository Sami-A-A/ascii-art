package main

import (
	// "ascii"
	"ascii"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	str := os.Args[1]
	for _, char := range str {
		firstLine, _ := ascii.FirstLine(char)
		fmt.Println(firstLine)
		// lastLine := firstLine + 8 // || firstline + lineCount
	}
}

// create an offset function: if char == 32, nil offset; for a unit increase in char, offset is 8 ...
// firstLine = var*8 ...
// var = char - 32
// if char == 32, firstLine = 1
