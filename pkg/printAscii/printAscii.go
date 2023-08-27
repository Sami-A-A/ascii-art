package ascii

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt() {

	// Error handling os arguments
	// Run only if more than 1 arg
	if len(os.Args) < 2 {
		fmt.Println("ERROR: No arguments to print")
		return
	} else if len(os.Args) != 2 {
		fmt.Println("ERROR: Place all text in a single argument")
		return
	} else if len(os.Args[1]) < 1 {
		return
	} else if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	// Split args by each new line
	phrases := strings.Split(os.Args[1], "\\n")

	// Print ascii for each phrase that was separated by a new line
	for _, phrase := range phrases {
		if len(phrase) == 0 {
			fmt.Println()
			continue
		}
		var firstLines []int
		// Save all first lines of each character into a slice
		for _, char := range phrase {
			if char < 32 || char > 126 {
				fmt.Println("ERROR: Character out of range")
				return
			} else {
				firstLine := getFirstLine(char)
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
			for j, line := range firstLines {
				filteredLines := lines[line-1]
				fmt.Print(filteredLines)
				firstLines[j]++
			}
			fmt.Println()
		}
	}
}

// Algorithm to find first line of each character
func getFirstLine(char rune) int {
	var firstLine int
	diff := int(char) - 32
	firstLine = diff*9 + 2
	return firstLine
}
