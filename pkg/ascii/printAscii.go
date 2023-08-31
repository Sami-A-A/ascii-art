package ascii

import "strings"

func printAsciiArt(inputString string) {

	inputLine := splitInput(inputString)

	asciiLines := make([][]string, 9)
	for i := 0; i < 9; i++ {
		asciiLines[i] = make([]string, len(inputLine))
	}

	for i := 0; i < len(inputLine); i++{
		for j := 0; j < 9; j++ {
			asciiLines[i][j] = getAsciiLine(inputLine, j)
		}
	}

}

func splitInput(inputText string) string{
	s := strings.Split(inputText, "\\n")
	printAsciiArt(strings.Join(s[1:], "\\n"))
	return s[0]
}

func getAsciiLine(char string, line int) string{
	
	// read the file
	// formula to get the line
	
	return
}