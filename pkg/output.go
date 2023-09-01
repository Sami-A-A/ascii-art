package ascii

import (
	"log"
	"os"
)

// this isn't exactly being used anywhere
// figure out how to import this to main
// Function to export art to specified file
func Output(art string, outputFlag *string) {
	err := os.WriteFile(*outputFlag, []byte(art), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
