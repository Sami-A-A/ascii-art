package ascii

import "fmt"

// import "fmt"

func FirstLine(char rune) (firstLine int, err string) {
	// var firstLine int
	if char >= 32 && char <= 126 {
		diff := int(char) - 32
		firstLine = diff*9 + 2
	} else {
		err = "ERROR: Character out of range"
		fmt.Println(err)
	}
	return firstLine, err
}

// move error handling to main
