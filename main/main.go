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
	if len(os.Args) < 2 {
		fmt.Println("ERROR: no arguments to print")
		return
	}
	var args string
	for i, arg := range os.Args {
		if i >= 1 {
			args += arg + " "
		}
	}
	var firstLines []int
	for _, char := range args {
		if char < 32 || char > 126 {
			fmt.Println("ERROR: Character out of range")
			return
		} else {
			firstLine := ascii.FirstLine(char)
			firstLines = append(firstLines, firstLine)
		}
	}
	// Read the file
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Split the data into a slice of strings, one for each line.
	lines := strings.Split(string(file), "\n")
	for i := 1; i < 9; i++ {
		for i, line := range firstLines {
			filteredLines := lines[line-1]
			fmt.Print(filteredLines)
			firstLines[i]++
		}
		fmt.Println()
	}
}
