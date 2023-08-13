package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	file, err := os.Open("standard.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	line := 1
// 	var firstLine int
// 	for scanner.Scan() {
// 		str := os.Args[1]
// 		for _, char := range str {
// 			firstLine = ascii.FirstLine(char)
// 			if scanner.Text() == "\n" {
// 				line++
// 			}
// 			if line == firstLine {
// 				fmt.Print(scanner.Text())
// 			}
// 		}
// 	}
// 	fmt.Println("first line = ", firstLine)

// 	// str := os.Args[1]
// 	// for _, char := range str {
// 	// 	firstLine, _ := ascii.FirstLine(char)
// 	// fmt.Println(firstLine)
// 	// lastLine := firstLine + 8 // || firstline + lineCount
// }

func main() {
	str := os.Args[1]
	var firstLines []int
	for _, char := range str {
		if char < 32 || char > 126 {
			fmt.Println("ERROR: Character out of range")
			return
		} else {
			firstLine := ascii.FirstLine(char)
			firstLines = append(firstLines, firstLine)
		}
	}
	// Open the file in read mode.
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Split the data into a slice of strings, one for each line.
	lines := strings.Split(string(file), "\n")
	// fmt.Println(lines)

	// Get the lines from 11 to 19.
	for i := 1; i < 8; i++ {
		for i, line := range firstLines {
			filteredLines := lines[line-1]
			fmt.Print(filteredLines)
			firstLines[i]++
		}
		fmt.Println()
	}
}
