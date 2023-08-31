package ascii

import (
	"fmt"
	"os"
	"strings"
)

func PrintAsciiArt(inputString string) {

	inputLine := splitInput(inputString)
	if inputLine == ""{
		return
	}

	asciiLines := make([][]string, 9)
	for i := 0; i < 9; i++ {
		asciiLines[i] = make([]string, len(inputLine))
	}

	for i := 0; i < 9; i++{
		asciiLines[i] = getAsciiLines(inputLine, i + 1)
		for j := 0; j < len(inputLine); j++{
			fmt.Print(asciiLines[i][j])
		}
		fmt.Println()
	}

}

func splitInput(inputText string) string{
	s := strings.Split(inputText, "\\n")
	if len(s) > 1 {
		PrintAsciiArt(strings.Join(s[0:len(s)-1], "\\n"))
	}
	return s[len(s)-1]
}

func getAsciiLines(inputLine string, n int) []string{
	
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(file), "\n")
	var outputLine []string

	for _, char := range inputLine {
		l := (int(char) - 32) * 9 + n
		outputLine = append(outputLine, lines[l]) 
	}

	return outputLine 
}