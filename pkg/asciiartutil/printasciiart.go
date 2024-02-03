package asciiartutil

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )


// Initializes AsciiArt by passing the string argument from main.go
func InitAsciiArt(input AsciiArt) {

	// // Splits the input string into a slice called lines 
	// lines := strings.Split(input, "\\n")

	// // Every line is followed by a single new line
	// // Prints the Ascii Art on the new line for every given line
	// // If the last line is a blank line it will not print anything
	// for index, line := range lines {
	// 	if line == "" && index != len(lines) - 1 {
	// 			fmt.Println()
	// 	} else if line != "" {
	// 		PrintAsciiArt(line, font)
	// 	}
	// }
	// return
}

// func PrintAsciiArt(line, font string){

// 	// Creating a two-dimensional set with empty values for each line 
// 	// in the ascii art of each character in the given line 
// 	asciiLines := make([][]string, 9)
// 	for i := 0; i < 9; i++ {
// 		asciiLines[i] = make([]string, len(line))
// 	}

// 	// Calls a function to read the lines from the font file and then prints
// 	// them out one by one adding a new line at the end of each row
// 	for i := 0; i < 9; i++{
// 		asciiLines[i] = GetAsciiLines(line, font, i + 1)
// 		for j := 0; j < len(line); j++{
// 			fmt.Print(asciiLines[i][j])
// 		}
// 		if i != 8 {
// 			fmt.Println()
// 		}
// 	}
// }

// func GetAsciiLines(phrase, font string, n int) []string{
	
// 	// Reads the file and parses the data as a slice of bytes
// 	file, err := os.ReadFile(font)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Makes the file easier to work with by parsing it into a slice 
// 	// of strings separated by the new lines as they appear on the file
// 	fileLines := strings.Split(string(file), "\n")

// 	var asciiLineRows []string

// 	// Runs an algorithm to get the desired lines of each ascii art from the file
// 	for _, char := range phrase {
// 		lineToGet := (int(char) - 32) * 9 + n
// 		asciiLineRows = append(asciiLineRows, fileLines[lineToGet]) 
// 	}

// 	return asciiLineRows 
// }