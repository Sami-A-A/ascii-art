package ascii

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt() {
	if len(os.Args) < 2 {
		fmt.Println("ERROR: no arguments to print")
		return
	}
	// var args string
	// for i, arg := range os.Args {
	// 	if i >= 1 {
	// 		args += arg + " "
	// 	}
	// }
	
	phrases := strings.Split(os.Args[1], "\\n")
	fmt.Println(phrases)
		
	for _, phrase := range phrases{
		var firstLines []int
		for _, char := range phrase {
			if char < 32 || char > 126 {
				fmt.Println("ERROR: Character out of range")
				return
			} else {
				firstLine := FirstLine(char)
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
