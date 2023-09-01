package ascii

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Create output flag and returns output file name
func OutputFlag() (string, bool) {
	var err bool = false

	// Create a flag named "output"
	outputFlag := flag.String("output", "output.txt", "assign output file")
	flag.Parse()

	// Confirm flag format at input
	if strings.Contains(os.Args[1], "-output") && !strings.HasPrefix(os.Args[1], "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		err = true
	}
	return *outputFlag, err
}

// Print art to <file name> from output flag
func PrintOutput(art string, fileName string) {
	err := os.WriteFile(fileName, []byte(art), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
